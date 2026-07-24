#!/usr/bin/env python3
"""Read go.mod require dependencies and fetch latest version release time and stars."""

import asyncio
import json
import re
import sys
from pathlib import Path

import aiohttp


GO_MOD = Path("/Users/heisonyee/Projects/x/go.mod")
OUTPUT_MD = Path("/Users/heisonyee/Projects/x/docs/deps-release-time.md")
CACHE_FILE = Path("/Users/heisonyee/Projects/x/.deps-cache.json")
CONCURRENCY = 12
TIMEOUT = 30
UNGH = "https://ungh.cc/repos"


def parse_direct_deps(text: str) -> list[tuple[str, str]]:
    """Parse direct dependencies from go.mod (skip indirect)."""
    deps: list[tuple[str, str]] = []
    in_require = False
    for raw in text.splitlines():
        line = raw.strip()
        if line.startswith("require ("):
            in_require = True
            continue
        if in_require and line == ")":
            in_require = False
            continue
        if not in_require:
            continue
        if not line or line.startswith("//"):
            continue
        if "// indirect" in line:
            continue
        parts = line.split()
        if len(parts) < 2:
            continue
        deps.append((parts[0], parts[1]))
    return deps


# ---------------------------------------------------------------------------
# Module -> GitHub owner/repo mapping for non-github.com module paths.
# Each entry: (matcher, callable(module) -> "owner/repo").
# ---------------------------------------------------------------------------

def _exact(owner_repo: str):
    def fn(module: str) -> str:
        return owner_repo
    return fn


def _prefix(prefix: str, owner_repo: str):
    def fn(module: str) -> str:
        return owner_repo
    return fn


def _golang_x(module: str) -> str:
    sub = re.sub(r"/v\d+$", "", module).split("/")[-1]
    return f"golang/{sub}"


def _github_owner_repo(module: str) -> str:
    parts = module.split("/")
    if len(parts) >= 3:
        return f"{parts[1]}/{parts[2]}"
    return ""


def _k8s_io(module: str) -> str:
    sub = module.split("/")[1]
    return f"kubernetes/{sub}"


def _gopkg(module: str) -> str:
    first = module[len("gopkg.in/"):].split("/")[0]
    owner = first.split(".")[0]
    return f"go-{owner}/{owner}"


GITHUB_MAP: list[tuple[str, callable]] = [
    ("golang.org/x/", _golang_x),
    ("google.golang.org/genproto", _exact("googleapis/go-genproto")),
    ("google.golang.org/grpc", _exact("grpc/grpc-go")),
    ("google.golang.org/protobuf", _exact("protocolbuffers/protobuf-go")),
    ("gopkg.in/", _gopkg),
    ("cel.dev/", _exact("google/cel-spec")),
    ("go.opentelemetry.io/", _exact("open-telemetry/opentelemetry-go")),
    ("go.temporal.io/", _exact("temporalio/sdk-go")),
    ("go.mongodb.org/", _exact("mongodb/mongo-go-driver")),
    ("go.etcd.io/etcd", _exact("etcd-io/etcd")),
    ("go.uber.org/", lambda m: "uber-go/" + m[len("go.uber.org/"):].split("/")[0]),
    ("resty.dev/", _exact("go-resty/resty")),
    ("k8s.io/", _k8s_io),
    ("filippo.io/edwards25519", _exact("FiloSottile/edwards25519")),
    ("filippo.io/csrf", _exact("FiloSottile/csrf")),
    ("filippo.io/age", _exact("FiloSottile/age")),
    ("resenje.org/", _exact("janos/singleflight")),
    ("go4.org/", _exact("camh/go4")),
    ("git.sr.ht/~jamesponddot/", _exact("")),
    ("github.com/", _github_owner_repo),
    ("gitlab.com/", lambda m: m.split("/")[1] + "/" + m.split("/")[2]),
    ("bitbucket.org/", lambda m: m.split("/")[1] + "/" + m.split("/")[2]),
]


def to_owner_repo(module: str) -> str:
    """Best-effort module path -> GitHub 'owner/repo' string."""
    for prefix, fn in GITHUB_MAP:
        if module.startswith(prefix):
            r = fn(module)
            if r and "/" in r and not r.endswith("/"):
                return r
    return ""


# ---------------------------------------------------------------------------
# Network fetchers
# ---------------------------------------------------------------------------

async def fetch_release(
    session: aiohttp.ClientSession,
    sem: asyncio.Semaphore,
    module: str,
    cache: dict,
) -> dict:
    if module in cache and "version" in cache[module]:
        return cache[module]
    encoded = "".join(f"!{c.lower()}" if c.isupper() else c for c in module)
    url = f"https://proxy.golang.org/{encoded}/@latest"
    last_error = "lookup failed"
    async with sem:
        for attempt in range(4):
            try:
                async with session.get(
                    url,
                    timeout=aiohttp.ClientTimeout(total=TIMEOUT),
                    headers={"Accept": "application/json"},
                ) as resp:
                    if resp.status != 200:
                        last_error = f"HTTP {resp.status}"
                        if resp.status == 404:
                            break
                    else:
                        data = await resp.json(content_type=None)
                        result = cache.setdefault(module, {
                            "module": module,
                            "stars": None,
                            "stars_error": None,
                        })
                        result["version"] = data.get("Version", "")
                        result["time"] = data.get("Time", "")
                        result["origin"] = (data.get("Origin") or {}).get("URL", "")
                        return result
            except Exception as e:  # noqa: BLE001
                last_error = f"{type(e).__name__}: {e}"
            await asyncio.sleep(0.5 * (2 ** attempt))
    entry = cache.setdefault(module, {"module": module, "stars": None})
    entry.setdefault("version", "")
    entry.setdefault("time", "")
    entry.setdefault("origin", "")
    entry["error"] = last_error
    return entry


async def fetch_stars(
    session: aiohttp.ClientSession,
    sem: asyncio.Semaphore,
    module: str,
    owner_repo: str,
    cache: dict,
) -> None:
    """Populate stars field for module in cache."""
    if not owner_repo:
        return
    entry = cache.setdefault(module, {"module": module})
    if entry.get("stars") is not None or entry.get("stars_error"):
        return
    url = f"{UNGH}/{owner_repo}"
    last_error = "lookup failed"
    async with sem:
        for attempt in range(3):
            try:
                async with session.get(
                    url,
                    timeout=aiohttp.ClientTimeout(total=TIMEOUT),
                    headers={"Accept": "application/json"},
                ) as resp:
                    if resp.status != 200:
                        last_error = f"HTTP {resp.status}"
                        if resp.status == 404:
                            break
                    else:
                        data = await resp.json(content_type=None)
                        stars = (data.get("repo") or {}).get("stars")
                        if isinstance(stars, int):
                            entry["stars"] = stars
                            entry["stars_owner_repo"] = owner_repo
                            return
                        last_error = "no stars field"
            except Exception as e:  # noqa: BLE001
                last_error = f"{type(e).__name__}: {e}"
            await asyncio.sleep(0.5 * (2 ** attempt))
    entry["stars"] = None
    entry["stars_error"] = last_error
    entry["stars_owner_repo"] = owner_repo


# ---------------------------------------------------------------------------
# Output
# ---------------------------------------------------------------------------

def fmt_int(n: int | None) -> str:
    if n is None:
        return "—"
    return f"{n:,}"


async def main() -> int:
    text = GO_MOD.read_text()
    deps = parse_direct_deps(text)
    print(f"Found {len(deps)} direct dependencies", file=sys.stderr)

    cache: dict = {}
    if CACHE_FILE.exists():
        try:
            cache = json.loads(CACHE_FILE.read_text())
        except Exception:  # noqa: BLE001
            cache = {}

    sem = asyncio.Semaphore(CONCURRENCY)
    connector = aiohttp.TCPConnector(limit=CONCURRENCY)
    async with aiohttp.ClientSession(connector=connector) as session:
        await asyncio.gather(
            *(fetch_release(session, sem, m, cache) for m, _ in deps)
        )
        owner_repos = {m: to_owner_repo(m) for m, _ in deps}
        await asyncio.gather(
            *(
                fetch_stars(session, sem, m, owner_repos[m], cache)
                for m, _ in deps
                if owner_repos[m]
            )
        )

    CACHE_FILE.write_text(json.dumps(cache, indent=2, sort_keys=True))

    lines: list[str] = []
    lines.append("# Direct Dependencies — Latest Version & Stars")
    lines.append("")
    lines.append(
        f"Source: [`go.mod`](./go.mod) • Generated by "
        f"`{Path(__file__).name}` • {len(deps)} direct dependencies"
    )
    lines.append("")
    lines.append(
        "Latest version data from `proxy.golang.org`; star counts from "
        "[`ungh.cc`](https://ungh.cc) (GitHub repo proxy)."
    )
    lines.append("")
    lines.append("| # | Module | Current | Latest | Latest Released | Stars | Repo |")
    lines.append("|---|--------|---------|--------|-----------------|-------|------|")

    rows: list[dict] = []
    for module, current in deps:
        info = cache.get(module, {})
        latest = info.get("version") or "—"
        when = info.get("time", "")
        when_date = when[:10] if when else "—"
        origin = info.get("origin", "")
        owner_repo = to_owner_repo(module)
        if origin:
            repo_cell = f"[link]({origin})"
        elif owner_repo:
            repo_cell = f"[link](https://github.com/{owner_repo})"
        else:
            repo_cell = "—"
        stars = info.get("stars")
        stars_cell = fmt_int(stars)
        err = info.get("error")
        if err or (latest == "—" and when_date == "—"):
            latest_cell = "— ⚠️"
            when_cell = f"_{err or 'lookup failed'}_"
            repo_cell = "—"
            stars_cell = "—"
        else:
            latest_cell = f"`{latest}`"
            when_cell = when_date
        rows.append({
            "module": module,
            "current": current,
            "latest_cell": latest_cell,
            "when_cell": when_cell,
            "stars_cell": stars_cell,
            "repo_cell": repo_cell,
            "sort_key": info.get("time") or "",
        })

    rows.sort(key=lambda r: (r["sort_key"] == "", r["sort_key"]))

    for i, r in enumerate(rows, 1):
        lines.append(
            f"| {i} | `{r['module']}` | `{r['current']}` | {r['latest_cell']} | "
            f"{r['when_cell']} | {r['stars_cell']} | {r['repo_cell']} |"
        )

    OUTPUT_MD.parent.mkdir(parents=True, exist_ok=True)
    OUTPUT_MD.write_text("\n".join(lines) + "\n")
    print(f"Wrote {OUTPUT_MD}", file=sys.stderr)
    return 0


if __name__ == "__main__":
    sys.exit(asyncio.run(main()))