module github.com/blink-io/x

go 1.24

toolchain go1.24.1

//godebug default=go1.23

require (
	github.com/42wim/httpsig v1.2.2
	github.com/99designs/gqlgen v0.17.70
	github.com/Azure/go-amqp v1.4.0
	github.com/BurntSushi/toml v1.5.0
	github.com/Code-Hex/go-generics-cache v1.5.1
	github.com/DATA-DOG/go-sqlmock v1.5.2
	github.com/IBM/sarama v1.45.1
	github.com/Netflix/go-env v0.1.2
	github.com/ProtonMail/gopenpgp/v2 v2.8.3
	github.com/VictoriaMetrics/easyproto v0.1.4
	github.com/VictoriaMetrics/fastcache v1.12.2
	github.com/alexedwards/argon2id v1.0.0
	github.com/allegro/bigcache/v3 v3.1.0
	github.com/alphadose/haxmap v1.4.1
	github.com/amazon-ion/ion-go v1.5.0
	github.com/ammario/tlru v0.4.0
	github.com/apache/pulsar-client-go v0.14.0
	github.com/apache/thrift v0.21.0
	github.com/apple/pkl-go v0.9.0
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2
	github.com/avast/retry-go/v4 v4.6.1
	github.com/beevik/guid v1.0.0
	github.com/bits-and-blooms/bloom/v3 v3.7.0
	github.com/blink-io/hyperbun v0.0.0-20250313040812-504f16b9229c
	github.com/blink-io/hypersql v0.0.0-20250215034254-fc8ce67faaa1
	github.com/blink-io/kratos-transport/transport/http3 v0.0.0-20250311064252-02bc3079463e
	github.com/blink-io/opt v0.0.0-20241010071220-8e1697ac4737
	github.com/blink-io/sqx v0.0.0-20241015070421-b93bddbcf4ec
	github.com/bmatcuk/doublestar/v4 v4.8.1
	github.com/bokwoon95/sq v0.5.1
	github.com/brianvoe/gofakeit/v7 v7.2.1
	github.com/bwmarrin/snowflake v0.3.0
	github.com/caarlos0/env/v11 v11.3.1
	github.com/carlmjohnson/requests v0.24.3
	github.com/cbroglie/mustache v1.4.0
	github.com/cenkalti/backoff/v4 v4.3.0
	github.com/cespare/xxhash/v2 v2.3.0
	github.com/cohesivestack/valgo v0.4.1
	github.com/coreos/go-oidc/v3 v3.13.0
	github.com/dchest/siphash v1.2.3
	github.com/dghubble/sling v1.4.2
	github.com/disgoorg/snowflake/v2 v2.0.3
	github.com/dlclark/regexp2 v1.11.5
	github.com/elastic/elastic-transport-go/v8 v8.6.1
	github.com/elastic/go-elasticsearch/v8 v8.17.1
	github.com/elliotchance/pie/v2 v2.9.1
	github.com/ergochat/readline v0.1.3
	github.com/failsafe-go/failsafe-go v0.6.9
	github.com/fxamacker/cbor/v2 v2.8.0
	github.com/georgysavva/scany/v2 v2.1.4
	github.com/getsentry/sentry-go v0.31.1
	github.com/go-chi/chi/v5 v5.2.1
	github.com/go-co-op/gocron/v2 v2.16.1
	github.com/go-crypt/crypt v0.3.5
	github.com/go-echarts/statsview v0.4.2
	github.com/go-faker/faker/v4 v4.6.0
	github.com/go-faster/city v1.0.1
	github.com/go-json-experiment/json v0.0.0-20250223041408-d3c622f1b874
	github.com/go-kit/log v0.2.1
	github.com/go-kratos/kratos/v2 v2.8.4
	github.com/go-logr/logr v1.4.2
	github.com/go-logr/stdr v1.2.2
	github.com/go-resty/resty/v2 v2.16.5
	github.com/go-sql-driver/mysql v1.9.1
	github.com/go-task/slim-sprig/v3 v3.0.0
	github.com/go-test/deep v1.1.1
	github.com/goccy/go-json v0.10.5
	github.com/gocraft/dbr/v2 v2.7.7
	github.com/gofiber/fiber/v3 v3.0.0-beta.4
	github.com/gofrs/uuid/v5 v5.3.2
	github.com/gogo/protobuf v1.3.2
	github.com/gomodule/redigo v1.9.2
	github.com/google/uuid v1.6.0
	github.com/graphql-go/graphql v0.8.1
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.3.1
	github.com/guregu/null/v5 v5.0.0
	github.com/h2non/filetype v1.1.3
	github.com/hashicorp/consul/api v1.32.0
	github.com/hashicorp/go-retryablehttp v0.7.7
	github.com/hashicorp/golang-lru/v2 v2.0.7
	github.com/hashicorp/mdns v1.0.6
	github.com/http-wasm/http-wasm-host-go v0.7.0
	github.com/jackc/pgx/v5 v5.7.4
	github.com/jackc/puddle/v2 v2.2.2
	github.com/jaevor/go-nanoid v1.4.0
	github.com/jaswdr/faker/v2 v2.3.3
	github.com/jellydator/ttlcache/v3 v3.3.0
	github.com/jfcg/sorty/v2 v2.1.1
	github.com/jmoiron/sqlx v1.4.0
	github.com/joho/godotenv v1.5.1
	github.com/karlseguin/ccache/v3 v3.0.6
	github.com/klauspost/cpuid/v2 v2.2.10
	github.com/leporo/sqlf v1.4.0
	github.com/liamg/memoryfs v1.6.0
	github.com/libp2p/go-reuseport v0.4.0
	github.com/linkedin/goavro/v2 v2.13.1
	github.com/lithammer/shortuuid/v4 v4.2.0
	github.com/lmittmann/tint v1.0.7
	github.com/madflojo/testcerts v1.4.0
	github.com/mailgun/raymond/v2 v2.0.48
	github.com/matthewhartstonge/argon2 v1.2.0
	github.com/mattn/go-runewidth v0.0.16
	github.com/mattn/go-sqlite3 v1.14.25
	github.com/meilisearch/meilisearch-go v0.31.0
	github.com/miekg/dns v1.1.64
	github.com/minio/highwayhash v1.0.3
	github.com/minio/minlz v1.0.0
	github.com/montanaflynn/stats v0.7.1
	github.com/natefinch/lumberjack/v3 v3.0.0-alpha
	github.com/nats-io/nats.go v1.40.1
	github.com/ncruces/go-strftime v0.1.9
	github.com/negrel/conc v0.5.0
	github.com/nicksnyder/go-i18n/v2 v2.5.1
	github.com/nyaruka/phonenumbers v1.6.0
	github.com/oklog/ulid/v2 v2.1.0
	github.com/onsi/ginkgo/v2 v2.23.3
	github.com/onsi/gomega v1.36.3
	github.com/opensearch-project/opensearch-go/v4 v4.4.0
	github.com/outcaste-io/ristretto v0.2.3
	github.com/panjf2000/ants/v2 v2.11.2
	github.com/pborman/uuid v1.2.1
	github.com/pelletier/go-toml/v2 v2.2.3
	github.com/phuslu/log v1.0.115
	github.com/phuslu/log-contrib/grpc v0.0.0-20250309100953-cb5b9b62dd61
	github.com/projectdiscovery/machineid v0.0.0-20240226150047-2e2c51e35983
	github.com/qiniu/qmgo v1.1.9
	github.com/quic-go/quic-go v0.50.1
	github.com/qustavo/sqlhooks/v2 v2.1.0
	github.com/rabbitmq/amqp091-go v1.10.0
	github.com/redis/go-redis/extra/rediscmd/v9 v9.7.3
	github.com/redis/go-redis/v9 v9.7.3
	github.com/redis/rueidis v1.0.56
	github.com/redis/rueidis/rueidishook v1.0.56
	github.com/reugn/go-quartz v0.14.0
	github.com/riverqueue/river v0.19.0
	github.com/riverqueue/river/riverdriver/riverdatabasesql v0.19.0
	github.com/riverqueue/river/riverdriver/riverpgxv5 v0.19.0
	github.com/rs/xid v1.6.0
	github.com/runZeroInc/mustache/v2 v2.1.0
	github.com/samber/go-singleflightx v0.3.2
	github.com/samber/hot v0.6.1
	github.com/samber/oops v1.17.0
	github.com/samber/slog-common v0.18.1
	github.com/samber/slog-http v1.6.0
	github.com/samber/slog-logrus/v2 v2.5.2
	github.com/samber/slog-multi v1.4.0
	github.com/samber/slog-sentry/v2 v2.9.3
	github.com/samber/slog-zap/v2 v2.6.2
	github.com/samber/slog-zerolog/v2 v2.7.3
	github.com/sanity-io/litter v1.5.8
	github.com/santhosh-tekuri/jsonschema/v6 v6.0.1
	github.com/segmentio/encoding v0.4.1
	github.com/segmentio/kafka-go v0.4.47
	github.com/segmentio/ksuid v1.0.4
	github.com/sethvargo/go-limiter v1.0.0
	github.com/sethvargo/go-password v0.3.1
	github.com/smartystreets/goconvey v1.8.1
	github.com/sourcegraph/conc v0.3.0
	github.com/spf13/cast v1.7.1
	github.com/spyzhov/ajson v0.9.6
	github.com/stretchr/testify v1.10.0
	github.com/teris-io/shortid v0.0.0-20220617161101-71ec9f2aa569
	github.com/tidwall/gjson v1.18.0
	github.com/twmb/franz-go v1.18.1
	github.com/twmb/franz-go/pkg/kmsg v1.9.0
	github.com/twmb/murmur3 v1.1.8
	github.com/unrolled/render v1.7.0
	github.com/uptrace/bun v1.2.11
	github.com/uptrace/bun/dialect/pgdialect v1.2.11
	github.com/uptrace/opentelemetry-go-extra/otelzap v0.3.2
	github.com/uptrace/uptrace-go v1.35.0
	github.com/valkey-io/valkey-go v1.0.56
	github.com/valkey-io/valkey-go/valkeyhook v1.0.56
	github.com/vmihailenco/go-tinylfu v0.2.2
	github.com/vmihailenco/msgpack/v5 v5.4.1
	github.com/zeebo/xxh3 v1.0.2
	github.com/zitadel/passwap v0.7.0
	gitlab.com/greyxor/slogor v1.6.1
	go.etcd.io/etcd/client/v3 v3.5.21
	go.mongodb.org/mongo-driver/v2 v2.1.0
	go.opentelemetry.io/otel v1.35.0
	go.opentelemetry.io/otel/trace v1.35.0
	go.temporal.io/sdk v1.33.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
	go.uber.org/zap/exp v0.3.0
	golang.org/x/crypto v0.36.0
	golang.org/x/exp v0.0.0-20250305212735-054e65f0b394
	golang.org/x/sys v0.31.0
	golang.org/x/text v0.23.0
	google.golang.org/genproto/googleapis/api v0.0.0-20250324211829-b45e905df463
	google.golang.org/grpc v1.71.1
	google.golang.org/protobuf v1.36.6
	gopkg.in/yaml.v3 v3.0.1
	k8s.io/client-go v0.32.3
	modernc.org/sqlite v1.37.0
	resenje.org/singleflight v0.4.3
)

replace github.com/bokwoon95/sq v0.5.1 => github.com/blink-io/sq v0.0.0-20240912023304-c0785df6eed1

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4 // indirect
	github.com/99designs/keyring v1.2.2 // indirect
	github.com/AthenZ/athenz v1.12.13 // indirect
	github.com/DataDog/zstd v1.5.7 // indirect
	github.com/DmitriyVTitov/size v1.5.0 // indirect
	github.com/Masterminds/squirrel v1.5.4 // indirect
	github.com/ProtonMail/go-crypto v1.1.6 // indirect
	github.com/ProtonMail/go-mime v0.0.0-20230322103455-7d82a3887f2f // indirect
	github.com/aarondl/json v0.0.0-20221020222930-8b0db17ef1bf // indirect
	github.com/andybalholm/brotli v1.1.1 // indirect
	github.com/ardielle/ardielle-go v1.5.2 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bits-and-blooms/bitset v1.22.0 // indirect
	github.com/cloudflare/circl v1.6.0 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/danieljoos/wincred v1.2.2 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/dvsekhvalnov/jose2go v1.8.0 // indirect
	github.com/eapache/go-resiliency v1.7.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20230731223053-c322873962e3 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect
	github.com/go-crypt/x v0.3.6 // indirect
	github.com/go-echarts/go-echarts/v2 v2.5.2 // indirect
	github.com/go-jose/go-jose/v4 v4.0.5 // indirect
	github.com/go-kratos/aegis v0.2.0 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-playground/form/v4 v4.2.1 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.26.0 // indirect
	github.com/godbus/dbus v0.0.0-20190726142602-4481cbc300e2 // indirect
	github.com/gofiber/schema v1.3.0 // indirect
	github.com/gofiber/utils/v2 v2.0.0-beta.8 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.2 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.2 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v1.0.0 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/pprof v0.0.0-20250317173921-a4b03ec1a45e // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.26.3 // indirect
	github.com/gsterjov/go-libsecret v0.0.0-20161001094733-a6f4afe4910c // indirect
	github.com/hamba/avro/v2 v2.28.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.6.3 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-metrics v0.5.4 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/hashicorp/serf v0.10.2 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.4 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jfcg/sixb v1.4.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jonboulle/clockwork v0.5.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/mailru/easyjson v0.9.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/microsoft/go-mssqldb v1.8.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mtibben/percent v0.2.1 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/nats-io/nkeys v0.4.10 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/nexus-rpc/sdk-go v0.3.0 // indirect
	github.com/philhofer/fwd v1.1.3-0.20240916144458-20a13a1f6b7c // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pierrec/lz4/v4 v4.1.22 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_golang v1.21.1 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.63.0 // indirect
	github.com/prometheus/procfs v0.16.0 // indirect
	github.com/puzpuzpuz/xsync/v3 v3.5.1 // indirect
	github.com/quic-go/qpack v0.5.1 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/riverqueue/river/riverdriver v0.19.0 // indirect
	github.com/riverqueue/river/rivershared v0.19.0 // indirect
	github.com/riverqueue/river/rivertype v0.19.0 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/robfig/cron v1.2.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/rs/cors v1.11.1 // indirect
	github.com/rs/zerolog v1.34.0 // indirect
	github.com/samber/lo v1.49.1 // indirect
	github.com/segmentio/asm v1.2.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/smarty/assertions v1.16.0 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/stephenafamo/scan v0.6.1 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/tetratelabs/wazero v1.9.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tidwall/sjson v1.2.5 // indirect
	github.com/tinylib/msgp v1.2.5 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/uptrace/bun/dialect/mysqldialect v1.2.11 // indirect
	github.com/uptrace/bun/dialect/sqlitedialect v1.2.11 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelutil v0.3.2 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.59.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/vektah/gqlparser/v2 v2.5.23 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/xo/dburl v0.23.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
	go.etcd.io/etcd/api/v3 v3.5.21 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.21 // indirect
	go.mongodb.org/mongo-driver v1.17.3 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/runtime v0.60.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp v0.11.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v1.35.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.35.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.35.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.35.0 // indirect
	go.opentelemetry.io/otel/log v0.11.0 // indirect
	go.opentelemetry.io/otel/metric v1.35.0 // indirect
	go.opentelemetry.io/otel/sdk v1.35.0 // indirect
	go.opentelemetry.io/otel/sdk/log v0.11.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.35.0 // indirect
	go.opentelemetry.io/proto/otlp v1.5.0 // indirect
	go.temporal.io/api v1.46.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/goleak v1.3.0 // indirect
	go.uber.org/mock v0.5.0 // indirect
	golang.org/x/mod v0.24.0 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/oauth2 v0.28.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
	golang.org/x/term v0.30.0 // indirect
	golang.org/x/time v0.11.0 // indirect
	golang.org/x/tools v0.31.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250324211829-b45e905df463 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/apimachinery v0.32.3 // indirect
	k8s.io/klog/v2 v2.130.1 // indirect
	k8s.io/utils v0.0.0-20250321185631-1f6e0b77f77e // indirect
	modernc.org/libc v1.62.1 // indirect
	modernc.org/mathutil v1.7.1 // indirect
	modernc.org/memory v1.9.1 // indirect
	sigs.k8s.io/json v0.0.0-20241014173422-cfa47c3a1cc8 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.6.0 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)
