module github.com/blink-io/x

go 1.23

require (
	github.com/BurntSushi/toml v1.4.0
	github.com/IBM/sarama v1.43.3
	github.com/VictoriaMetrics/fastcache v1.12.2
	github.com/allegro/bigcache/v3 v3.1.0
	github.com/ammario/tlru v0.4.0
	github.com/apache/thrift v0.20.0
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2
	github.com/avast/retry-go/v4 v4.6.0
	github.com/beevik/guid v1.0.0
	github.com/bits-and-blooms/bloom/v3 v3.7.0
	github.com/brianvoe/gofakeit/v7 v7.0.4
	github.com/bwmarrin/snowflake v0.3.0
	github.com/caarlos0/env/v10 v10.0.0
	github.com/carlmjohnson/requests v0.24.2
	github.com/cespare/xxhash/v2 v2.3.0
	github.com/cohesivestack/valgo v0.4.1
	github.com/dchest/siphash v1.2.3
	github.com/dghubble/sling v1.4.2
	github.com/disgoorg/snowflake/v2 v2.0.3
	github.com/failsafe-go/failsafe-go v0.6.8
	github.com/fxamacker/cbor/v2 v2.7.0
	github.com/georgysavva/scany/v2 v2.1.3
	github.com/getsentry/sentry-go v0.28.1
	github.com/go-chi/chi/v5 v5.1.0
	github.com/go-co-op/gocron/v2 v2.11.0
	github.com/go-faster/city v1.0.1
	github.com/go-kit/log v0.2.1
	github.com/go-kratos/kratos/v2 v2.8.0
	github.com/go-logr/logr v1.4.2
	github.com/go-logr/stdr v1.2.2
	github.com/go-resty/resty/v2 v2.14.0
	github.com/go-task/slim-sprig/v3 v3.0.0
	github.com/goccy/go-json v0.10.3
	github.com/gofiber/fiber/v3 v3.0.0-beta.3
	github.com/gofrs/uuid/v5 v5.3.0
	github.com/gogo/protobuf v1.3.2
	github.com/google/uuid v1.6.0
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.1.0
	github.com/guregu/null/v5 v5.0.0
	github.com/h2non/filetype v1.1.3
	github.com/hashicorp/consul/api v1.29.4
	github.com/hashicorp/go-retryablehttp v0.7.7
	github.com/hashicorp/golang-lru/v2 v2.0.7
	github.com/hashicorp/mdns v1.0.5
	github.com/jackc/puddle/v2 v2.2.1
	github.com/jaevor/go-nanoid v1.4.0
	github.com/jaswdr/faker/v2 v2.3.0
	github.com/jellydator/ttlcache/v3 v3.3.0
	github.com/joho/godotenv v1.5.1
	github.com/karlseguin/ccache/v3 v3.0.5
	github.com/linkedin/goavro/v2 v2.13.0
	github.com/lithammer/shortuuid/v4 v4.0.0
	github.com/lmittmann/tint v1.0.5
	github.com/mattn/go-runewidth v0.0.16
	github.com/miekg/dns v1.1.62
	github.com/montanaflynn/stats v0.7.1
	github.com/natefinch/lumberjack/v3 v3.0.0-alpha
	github.com/nats-io/nats.go v1.37.0
	github.com/ncruces/go-strftime v0.1.9
	github.com/nicksnyder/go-i18n/v2 v2.4.0
	github.com/oklog/ulid/v2 v2.1.0
	github.com/onsi/ginkgo/v2 v2.20.1
	github.com/onsi/gomega v1.34.1
	github.com/outcaste-io/ristretto v0.2.3
	github.com/pelletier/go-toml/v2 v2.2.3
	github.com/phuslu/log v1.0.110
	github.com/phuslu/log-contrib/grpc v0.0.0-20240622164905-82dca04dc910
	github.com/qiniu/qmgo v1.1.8
	github.com/quic-go/quic-go v0.46.0
	github.com/redis/go-redis/extra/rediscmd/v9 v9.5.3
	github.com/redis/go-redis/v9 v9.6.1
	github.com/redis/rueidis v1.0.44
	github.com/redis/rueidis/rueidishook v1.0.44
	github.com/reugn/go-quartz v0.12.0
	github.com/riverqueue/river v0.11.4
	github.com/rs/xid v1.6.0
	github.com/samber/slog-common v0.17.1
	github.com/samber/slog-fiber v1.16.2
	github.com/samber/slog-http v1.4.2
	github.com/samber/slog-logrus/v2 v2.5.0
	github.com/samber/slog-multi v1.2.1
	github.com/samber/slog-sentry/v2 v2.8.0
	github.com/samber/slog-zap/v2 v2.6.0
	github.com/samber/slog-zerolog/v2 v2.7.0
	github.com/santhosh-tekuri/jsonschema/v6 v6.0.1
	github.com/segmentio/encoding v0.4.0
	github.com/segmentio/ksuid v1.0.4
	github.com/sethvargo/go-password v0.3.1
	github.com/smartystreets/goconvey v1.8.1
	github.com/sourcegraph/conc v0.3.0
	github.com/spf13/cast v1.7.0
	github.com/stretchr/testify v1.9.0
	github.com/teris-io/shortid v0.0.0-20220617161101-71ec9f2aa569
	github.com/twmb/murmur3 v1.1.8
	github.com/tx7do/kratos-transport/transport/http3 v1.2.13
	github.com/uptrace/opentelemetry-go-extra/otelzap v0.3.1
	github.com/vingarcia/ksql v1.12.0
	github.com/vingarcia/ksql/adapters/modernc-ksqlite v1.12.0
	github.com/vmihailenco/go-tinylfu v0.2.2
	github.com/vmihailenco/msgpack/v5 v5.4.1
	github.com/zeebo/xxh3 v1.0.2
	github.com/zitadel/passwap v0.6.0
	gitlab.com/greyxor/slogor v1.2.11
	go.etcd.io/etcd/client/v3 v3.5.15
	go.mongodb.org/mongo-driver v1.16.1
	go.temporal.io/sdk v1.28.1
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
	go.uber.org/zap/exp v0.2.0
	golang.org/x/sys v0.24.0
	golang.org/x/text v0.17.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240826202546-f6391c0de4c7
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.2
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/bits-and-blooms/bitset v1.14.2 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/eapache/go-resiliency v1.7.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20230731223053-c322873962e3 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a // indirect
	github.com/fatih/color v1.17.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.5 // indirect
	github.com/go-kratos/aegis v0.2.0 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-playground/form/v4 v4.2.1 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.22.0 // indirect
	github.com/gofiber/fiber/v2 v2.52.5 // indirect
	github.com/gofiber/utils/v2 v2.0.0-beta.6 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/pprof v0.0.0-20240727154555-813a5fbdbec8 // indirect
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.22.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.6.3 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.6.0 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.4 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jonboulle/clockwork v0.4.0 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/klauspost/cpuid/v2 v2.2.8 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/nats-io/nkeys v0.4.7 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/nexus-rpc/sdk-go v0.0.10 // indirect
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/quic-go/qpack v0.4.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/riverqueue/river/riverdriver v0.11.4 // indirect
	github.com/riverqueue/river/rivershared v0.11.4 // indirect
	github.com/riverqueue/river/rivertype v0.11.4 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/robfig/cron v1.2.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/rs/zerolog v1.33.0 // indirect
	github.com/samber/lo v1.47.0 // indirect
	github.com/segmentio/asm v1.2.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/smarty/assertions v1.16.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelutil v0.3.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.55.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
	go.etcd.io/etcd/api/v3 v3.5.15 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.15 // indirect
	go.opentelemetry.io/otel v1.29.0 // indirect
	go.opentelemetry.io/otel/log v0.5.0 // indirect
	go.opentelemetry.io/otel/metric v1.29.0 // indirect
	go.opentelemetry.io/otel/trace v1.29.0 // indirect
	go.temporal.io/api v1.38.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/goleak v1.3.0 // indirect
	go.uber.org/mock v0.4.0 // indirect
	golang.org/x/crypto v0.26.0 // indirect
	golang.org/x/exp v0.0.0-20240823005443-9b4947da3948 // indirect
	golang.org/x/mod v0.20.0 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/time v0.6.0 // indirect
	golang.org/x/tools v0.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240826202546-f6391c0de4c7 // indirect
	modernc.org/gc/v3 v3.0.0-20240801135723-a856999a2e4a // indirect
	modernc.org/libc v1.59.9 // indirect
	modernc.org/mathutil v1.6.0 // indirect
	modernc.org/memory v1.8.0 // indirect
	modernc.org/sqlite v1.32.0 // indirect
	modernc.org/strutil v1.2.0 // indirect
	modernc.org/token v1.1.0 // indirect
)
