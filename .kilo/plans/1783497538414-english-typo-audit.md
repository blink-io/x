# Plan: English Typo Audit of /Users/heisonyee/Projects/x

## Goal
Scan the Go monorepo at `/Users/heisonyee/Projects/x` for English misspellings (typos) and produce a report with the affected files and exact line numbers. The previous audit (subagent) searched for both Chinese 错别字 and English typos; this plan focuses narrowly on **English typos in human-readable text** (comments, error messages, log strings, doc files) and also re-checks the previous run to surface anything the broad regex may have missed.

## Scope

### In scope
- All `*.go` files under `/Users/heisonyee/Projects/x` — production and test code.
- All `*.md` files (top-level `README.md`, per-module `README.md`, `docs/**/*.md`).
- Text content: comments (`//`, `/* */`), error messages (`fmt.Errorf`, `errors.New`, `log.*`, `panic(...)`), struct/function doc comments, and Markdown prose.
- Subagent `task` tool can be used to perform a deep, multi-angle pass.

### Out of scope
- Auto-generated files: `*.pb.go`, `*.gen.go`, `zz_generated_*.go`, files under `vendor/`, `node_modules/`, `third_party/`.
- `go.sum`, `LICENSE`, `*.lock` files.
- Go identifiers (variable/function/type names) unless the misspelling is unambiguous (e.g., `Recieve` instead of `Receive`).
- Binary files.

## Approach

1. **Targeted word-list sweep** with `rg -n` against a curated list of common English misspellings:
   - `recieve`, `recieved`, `reciever`
   - `occured`, `occuring`, `occurance`
   - `succesful`, `succesfully`, `sucess`, `sucessful`
   - `paramter`, `parmeter`, `paramaters`
   - `lenght`, `widht`, `heigth`
   - `seperator`, `seperate`, `seperately`
   - `existant`, `existance`
   - `writting`, `writte`, `writen`
   - `enviroment`, `enviromnent`
   - `definately`, `definatly`
   - `untill`, `wether` (when used as a conjunction), `acheive`, `acheived`
   - `beggining`, `comming`, `embeded`, `embarass`
   - `teh`, `adn`, `nad`, `htat`
   - `priviledge`, `priviledge`
   - `accross`, `accomodate`, `accomodation`
   - `catagory`, `catagories`
   - `maintainance`, `maintenence`
   - `inital`, `initally`, `initialy`
   - `noticable`, `noticably`
   - `persistant`, `persistance`
   - `preceeding`, `succeding` (when not the legal term)
   - `tomorow`, `tommorow`
   - `truely`, `unecessary`, `usefull`
   - `wierd`, `wich` (when "which" is meant)
   - `alot`, `everytime` (as adverb)

2. **Heuristic pattern sweep** with `rg` for suspect constructions:
   - `\b(the|to|of|a|is)\b\s+\b(the|to|of|a|is)\b` repeated prepositions/articles (catches doubled words like "the the").
   - `[a-z]+[A-Z][a-z]+` inside string literals or comments where it looks like a smashed word.
   - `[^aeiou]{5,}` runs (long consonant clusters often indicate typos).

3. **Subagent deep-pass** with two parallel `task` calls:
   - One agent scans comments + error messages + log messages for English misspellings.
   - One agent scans `*.md` files for prose typos.

4. **Manual verification**: for every match, open the file and confirm the line is actually a typo in natural language (not a correct technical term, library name, or identifier).

## Deliverable

A report (printed in chat) with the following format:

```
## English typo report — /Users/heisonyee/Projects/x

### Critical
- path/to/file.go:LINE — `typo` → `correction`  (context: "surrounding text")

### Minor
- ...
```

If the codebase is clean for English typos, state so explicitly with the count of files inspected.

## Steps

1. Run the targeted `rg` word-list sweep; record matches with `-n`.
2. Run the heuristic pattern sweep; record matches with `-n`.
3. Dispatch two parallel `task` subagents (comments/code strings and Markdown).
4. Deduplicate results, verify each by reading the affected line.
5. Group by severity (critical = changed meaning; minor = spelling only).
6. Print the report in chat.

## Risks / Caveats
- False positives from identifiers like `Callback`, `Commit`, `Middleware`, `Occurrence` (correct spellings that may match regex roots). Mitigation: human-verify every match.
- Library / package names (e.g. `gRPC`, `Prometheus`) can contain unusual casing. Excluded from the word-list sweep.
- Very large `docs/` trees: rely on `rg` for speed.

## Validation
- Each reported line is read back with the `read` tool to confirm the typo and capture the surrounding context for the report.
- No source files are modified — this is a read-only audit.

## Open questions
- None. The user asked for a report; no fix is in scope unless requested.
