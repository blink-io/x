module github.com/blink-io/x

go 1.25.3

//godebug default=go1.23

require (
	github.com/42wim/httpsig v1.2.3
	github.com/99designs/gqlgen v0.17.86
	github.com/Azure/go-amqp v1.5.1
	github.com/BurntSushi/toml v1.6.0
	github.com/Code-Hex/go-generics-cache v1.5.1
	github.com/DATA-DOG/go-sqlmock v1.5.2
	github.com/IBM/sarama v1.46.3
	github.com/Masterminds/squirrel v1.5.4
	github.com/Netflix/go-env v0.1.2
	github.com/ProtonMail/gopenpgp/v2 v2.9.0
	github.com/VictoriaMetrics/easyproto v1.1.3
	github.com/VictoriaMetrics/fastcache v1.13.2
	github.com/alexedwards/argon2id v1.0.0
	github.com/alexlast/bunzap v0.1.0
	github.com/allegro/bigcache/v3 v3.1.0
	github.com/alphadose/haxmap v1.4.1
	github.com/amazon-ion/ion-go v1.5.0
	github.com/ammario/tlru v0.4.0
	github.com/apache/pulsar-client-go v0.18.0
	github.com/apache/thrift v0.22.0
	github.com/apple/pkl-go v0.12.1
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2
	github.com/avast/retry-go/v4 v4.7.0
	github.com/bartventer/httpcache v0.12.0
	github.com/beevik/etree v1.6.0
	github.com/beevik/guid v1.0.0
	github.com/bits-and-blooms/bloom/v3 v3.7.1
	github.com/blink-io/hyperbun v0.0.0-20250826092651-ce41bbffc41b
	github.com/blink-io/hypersql v0.0.0-20250715100453-239dcf60539c
	github.com/blink-io/kratos-transport/transport/http3 v0.0.0-20250903025134-baecc2c6e6cd
	github.com/blink-io/opt v0.0.0-20250826090112-aa8ab860d240
	github.com/blink-io/sq v0.0.0-20250704034929-fc04e805db4a
	github.com/blink-io/sqx v0.0.0-20250630144706-16ce62bfdaed
	github.com/bmatcuk/doublestar/v4 v4.9.2
	github.com/brianvoe/gofakeit/v7 v7.14.0
	github.com/bwmarrin/snowflake v0.3.0
	github.com/caarlos0/env/v11 v11.3.1
	github.com/carlmjohnson/requests v0.25.1
	github.com/cbroglie/mustache v1.4.0
	github.com/cenkalti/backoff/v4 v4.3.0
	github.com/cespare/xxhash/v2 v2.3.0
	github.com/cloudwego/fastpb v0.0.6
	github.com/coder/websocket v1.8.14
	github.com/cohesivestack/valgo v0.7.1
	github.com/coreos/go-oidc/v3 v3.17.0
	github.com/dchest/siphash v1.2.3
	github.com/dghubble/sling v1.4.2
	github.com/dgraph-io/ristretto/v2 v2.3.0
	github.com/disgoorg/snowflake/v2 v2.0.3
	github.com/dlclark/regexp2 v1.11.5
	github.com/dustin/go-humanize v1.0.1
	github.com/elastic/elastic-transport-go/v8 v8.8.0
	github.com/elastic/go-elasticsearch/v8 v8.19.1
	github.com/elliotchance/pie/v2 v2.9.1
	github.com/ergochat/readline v0.1.3
	github.com/failsafe-go/failsafe-go v0.9.5
	github.com/fatih/color v1.18.0
	github.com/fxamacker/cbor/v2 v2.9.0
	github.com/georgysavva/scany/v2 v2.1.4
	github.com/getsentry/sentry-go v0.40.0
	github.com/go-chi/chi/v5 v5.2.3
	github.com/go-co-op/gocron-redis-lock/v2 v2.2.1
	github.com/go-co-op/gocron/v2 v2.19.0
	github.com/go-crypt/crypt v0.4.7
	github.com/go-echarts/statsview v0.4.2
	github.com/go-faker/faker/v4 v4.7.0
	github.com/go-faster/city v1.0.1
	github.com/go-json-experiment/json v0.0.0-20251027170946-4849db3c2f7e
	github.com/go-kit/log v0.2.1
	github.com/go-kratos/kratos/v2 v2.9.2
	github.com/go-logr/logr v1.4.3
	github.com/go-logr/stdr v1.2.2
	github.com/go-openapi/inflect v0.21.5
	github.com/go-resty/resty/v2 v2.17.1
	github.com/go-task/slim-sprig/v3 v3.0.0
	github.com/go-test/deep v1.1.1
	github.com/go-viper/mapstructure/v2 v2.5.0
	github.com/gobwas/ws v1.4.0
	github.com/goccy/go-json v0.10.5
	github.com/gocraft/dbr/v2 v2.7.7
	github.com/gofiber/fiber/v3 v3.0.0-rc.3
	github.com/gofrs/uuid/v5 v5.4.0
	github.com/gogo/protobuf v1.3.2
	github.com/gomodule/redigo v1.9.3
	github.com/google/flatbuffers v25.12.19+incompatible
	github.com/google/uuid v1.6.0
	github.com/gorilla/csrf v1.7.3
	github.com/gorilla/websocket v1.5.4-0.20250319132907-e064f32e3674
	github.com/graphql-go/graphql v0.8.1
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.3.3
	github.com/guregu/null/v5 v5.0.0
	github.com/guregu/null/v6 v6.0.0
	github.com/h2non/filetype v1.1.3
	github.com/hashicorp/consul/api v1.33.0
	github.com/hashicorp/go-retryablehttp v0.7.8
	github.com/hashicorp/golang-lru/v2 v2.0.7
	github.com/hashicorp/mdns v1.0.6
	github.com/http-wasm/http-wasm-host-go v0.7.0
	github.com/huandu/xstrings v1.5.0
	github.com/huml-lang/go-huml v0.3.0
	github.com/imroc/req/v3 v3.57.0
	github.com/jackc/pgx/v5 v5.8.0
	github.com/jackc/puddle/v2 v2.2.2
	github.com/jaevor/go-nanoid v1.4.0
	github.com/jaswdr/faker/v2 v2.9.1
	github.com/jellydator/ttlcache/v3 v3.4.0
	github.com/jfcg/sorty/v2 v2.1.2
	github.com/jmoiron/sqlx v1.4.0
	github.com/joho/godotenv v1.5.1
	github.com/justinas/nosurf v1.2.0
	github.com/karlseguin/ccache/v3 v3.0.8
	github.com/kelindar/event v1.5.2
	github.com/klauspost/compress v1.18.2
	github.com/klauspost/cpuid/v2 v2.3.0
	github.com/liamg/memoryfs v1.6.0
	github.com/lib/pq v1.10.9
	github.com/libp2p/go-reuseport v0.4.0
	github.com/linkedin/goavro/v2 v2.14.1
	github.com/lithammer/shortuuid/v4 v4.2.0
	github.com/lmittmann/tint v1.1.2
	github.com/madflojo/testcerts v1.5.0
	github.com/mailgun/raymond/v2 v2.0.48
	github.com/matoous/go-nanoid v1.5.1
	github.com/matthewhartstonge/argon2 v1.4.4
	github.com/mattn/go-runewidth v0.0.19
	github.com/meilisearch/meilisearch-go v0.35.1
	github.com/mholt/archives v0.1.5
	github.com/miekg/dns v1.1.70
	github.com/minio/highwayhash v1.0.3
	github.com/minio/minlz v1.0.1
	github.com/montanaflynn/stats v0.7.1
	github.com/natefinch/lumberjack/v3 v3.0.0-alpha
	github.com/nats-io/nats.go v1.48.0
	github.com/ncruces/go-strftime v1.0.0
	github.com/negrel/conc v0.5.0
	github.com/nicksnyder/go-i18n/v2 v2.6.1
	github.com/nyaruka/phonenumbers v1.6.8
	github.com/oklog/ulid/v2 v2.1.1
	github.com/onatm/clockwerk v1.1.0
	github.com/onsi/ginkgo/v2 v2.27.5
	github.com/onsi/gomega v1.39.0
	github.com/opensearch-project/opensearch-go/v4 v4.6.0
	github.com/panjf2000/ants/v2 v2.11.4
	github.com/pborman/uuid v1.2.1
	github.com/pelletier/go-toml/v2 v2.2.4
	github.com/phuslu/log v1.0.121
	github.com/phuslu/log-contrib/grpc v0.0.0-20250309100953-cb5b9b62dd61
	github.com/projectdiscovery/machineid v0.0.0-20250715113114-c77eb3567582
	github.com/qiniu/qmgo v1.1.10
	github.com/quic-go/quic-go v0.59.0
	github.com/qustavo/sqlhooks/v2 v2.1.0
	github.com/rabbitmq/amqp091-go v1.10.0
	github.com/redis/go-redis/extra/rediscmd/v9 v9.17.2
	github.com/redis/go-redis/v9 v9.17.2
	github.com/redis/rueidis v1.0.70
	github.com/redis/rueidis/rueidishook v1.0.70
	github.com/reugn/async v0.9.0
	github.com/reugn/go-quartz v0.15.2
	github.com/riverqueue/river v0.30.0
	github.com/riverqueue/river/riverdriver/riverdatabasesql v0.30.0
	github.com/riverqueue/river/riverdriver/riverpgxv5 v0.30.0
	github.com/rs/xid v1.6.0
	github.com/runZeroInc/mustache/v2 v2.1.0
	github.com/samber/do/v2 v2.0.0
	github.com/samber/go-singleflightx v0.3.2
	github.com/samber/hot v0.11.0
	github.com/samber/oops v1.20.0
	github.com/samber/slog-common v0.19.0
	github.com/samber/slog-http v1.10.0
	github.com/samber/slog-logrus/v2 v2.5.2
	github.com/samber/slog-multi v1.7.0
	github.com/samber/slog-sentry/v2 v2.10.1
	github.com/samber/slog-zap/v2 v2.6.2
	github.com/samber/slog-zerolog/v2 v2.9.0
	github.com/sanity-io/litter v1.5.8
	github.com/santhosh-tekuri/jsonschema/v6 v6.0.2
	github.com/segmentio/encoding v0.5.3
	github.com/segmentio/kafka-go v0.4.49
	github.com/segmentio/ksuid v1.0.4
	github.com/sethvargo/go-limiter v1.1.0
	github.com/sethvargo/go-password v0.3.1
	github.com/smartystreets/goconvey v1.8.1
	github.com/sourcegraph/conc v0.3.0
	github.com/spf13/cast v1.10.0
	github.com/spyzhov/ajson v0.9.6
	github.com/sqids/sqids-go v0.4.1
	github.com/stephenafamo/scan v0.7.0
	github.com/stretchr/testify v1.11.1
	github.com/teris-io/shortid v0.0.0-20220617161101-71ec9f2aa569
	github.com/tidwall/gjson v1.18.0
	github.com/twmb/franz-go v1.20.6
	github.com/twmb/franz-go/pkg/kmsg v1.12.0
	github.com/twmb/murmur3 v1.1.8
	github.com/unrolled/render v1.7.0
	github.com/uptrace/bun v1.2.16
	github.com/uptrace/bun/dialect/pgdialect v1.2.16
	github.com/uptrace/bun/extra/bundebug v1.2.16
	github.com/uptrace/bun/extra/bunslog v1.2.16
	github.com/uptrace/opentelemetry-go-extra/otelzap v0.3.2
	github.com/uptrace/uptrace-go v1.39.0
	github.com/valkey-io/valkey-go v1.0.70
	github.com/valkey-io/valkey-go/valkeyhook v1.0.70
	github.com/vmihailenco/go-tinylfu v0.2.2
	github.com/vmihailenco/msgpack/v5 v5.4.1
	github.com/xtaci/kcp-go/v5 v5.6.61
	github.com/zeebo/xxh3 v1.0.2
	github.com/zitadel/passwap v0.12.0
	gitlab.com/greyxor/slogor v1.6.4
	go.etcd.io/etcd/client/v3 v3.6.7
	go.mongodb.org/mongo-driver/v2 v2.4.1
	go.opentelemetry.io/otel v1.39.0
	go.opentelemetry.io/otel/trace v1.39.0
	go.temporal.io/sdk v1.39.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.1
	go.uber.org/zap/exp v0.3.0
	golang.org/x/crypto v0.47.0
	golang.org/x/exp v0.0.0-20260112195511-716be5621a96
	golang.org/x/net v0.49.0
	golang.org/x/oauth2 v0.34.0
	golang.org/x/sys v0.40.0
	golang.org/x/text v0.33.0
	google.golang.org/genproto/googleapis/api v0.0.0-20260112192933-99fd39fd28a9
	google.golang.org/grpc v1.78.0
	google.golang.org/protobuf v1.36.11
	gopkg.in/yaml.v3 v3.0.1
	k8s.io/client-go v0.35.0
	resenje.org/singleflight v0.4.3
	resty.dev/v3 v3.0.0-beta.6
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/AthenZ/athenz v1.12.29 // indirect
	github.com/DataDog/zstd v1.5.7 // indirect
	github.com/DmitriyVTitov/size v1.5.0 // indirect
	github.com/Masterminds/semver/v3 v3.4.0 // indirect
	github.com/ProtonMail/go-crypto v1.3.0 // indirect
	github.com/ProtonMail/go-mime v0.0.0-20230322103455-7d82a3887f2f // indirect
	github.com/RoaringBitmap/roaring/v2 v2.8.0 // indirect
	github.com/STARRY-S/zip v0.2.3 // indirect
	github.com/andybalholm/brotli v1.2.0 // indirect
	github.com/ardielle/ardielle-go v1.5.2 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bits-and-blooms/bitset v1.24.4 // indirect
	github.com/bodgit/plumbing v1.3.0 // indirect
	github.com/bodgit/sevenzip v1.6.1 // indirect
	github.com/bodgit/windows v1.0.1 // indirect
	github.com/bokwoon95/sq v0.5.1 // indirect
	github.com/bytedance/gopkg v0.1.3 // indirect
	github.com/cenkalti/backoff/v5 v5.0.3 // indirect
	github.com/clipperhouse/stringish v0.1.1 // indirect
	github.com/clipperhouse/uax29/v2 v2.3.0 // indirect
	github.com/cloudflare/circl v1.6.1 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd/v22 v22.6.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dsnet/compress v0.0.2-0.20230904184137-39efe44ab707 // indirect
	github.com/eapache/go-resiliency v1.7.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20230731223053-c322873962e3 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.11 // indirect
	github.com/go-crypt/x v0.4.9 // indirect
	github.com/go-echarts/go-echarts/v2 v2.6.7 // indirect
	github.com/go-jose/go-jose/v4 v4.1.3 // indirect
	github.com/go-kratos/aegis v0.2.0 // indirect
	github.com/go-logfmt/logfmt v0.6.1 // indirect
	github.com/go-playground/form/v4 v4.3.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.28.0 // indirect
	github.com/go-redsync/redsync/v4 v4.14.0 // indirect
	github.com/go-sql-driver/mysql v1.9.3 // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gofiber/schema v1.6.0 // indirect
	github.com/gofiber/utils/v2 v2.0.0-rc.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.2 // indirect
	github.com/golang-jwt/jwt/v5 v5.3.0 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v1.0.0 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/pprof v0.0.0-20251114195745-4902fdda35c8 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/gorilla/securecookie v1.1.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.3 // indirect
	github.com/hamba/avro/v2 v2.30.0 // indirect
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
	github.com/icholy/digest v1.1.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.4 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jfcg/sixb/v2 v2.3.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jonboulle/clockwork v0.5.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/klauspost/pgzip v1.2.6 // indirect
	github.com/klauspost/reedsolomon v1.12.0 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-sqlite3 v1.14.32 // indirect
	github.com/microsoft/go-mssqldb v1.9.4 // indirect
	github.com/mikelolasagasti/xz v1.0.1 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.3-0.20250322232337-35a7c28c31ee // indirect
	github.com/mschoch/smat v0.2.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/nats-io/nkeys v0.4.11 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/nexus-rpc/sdk-go v0.5.1 // indirect
	github.com/nwaples/rardecode/v2 v2.2.1 // indirect
	github.com/philhofer/fwd v1.2.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.22 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_golang v1.23.2 // indirect
	github.com/prometheus/client_model v0.6.2 // indirect
	github.com/prometheus/common v0.67.2 // indirect
	github.com/prometheus/procfs v0.19.2 // indirect
	github.com/puzpuzpuz/xsync/v3 v3.5.1 // indirect
	github.com/quic-go/qpack v0.6.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20250401214520-65e299d6c5c9 // indirect
	github.com/refraction-networking/utls v1.8.1 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/riverqueue/river/riverdriver v0.30.0 // indirect
	github.com/riverqueue/river/rivershared v0.30.0 // indirect
	github.com/riverqueue/river/rivertype v0.30.0 // indirect
	github.com/robfig/cron v1.2.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/rs/cors v1.11.1 // indirect
	github.com/rs/zerolog v1.34.0 // indirect
	github.com/samber/go-type-to-string v1.8.0 // indirect
	github.com/samber/lo v1.52.0 // indirect
	github.com/segmentio/asm v1.2.1 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/smarty/assertions v1.16.0 // indirect
	github.com/sorairolake/lzip-go v0.3.8 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/spf13/afero v1.15.0 // indirect
	github.com/stretchr/objx v0.5.3 // indirect
	github.com/tetratelabs/wazero v1.10.1 // indirect
	github.com/theparanoids/crypki v1.20.9 // indirect
	github.com/tidwall/match v1.2.0 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tidwall/sjson v1.2.5 // indirect
	github.com/tinylib/msgp v1.5.0 // indirect
	github.com/tjfoc/gmsm v1.4.1 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/ulikunitz/xz v0.5.15 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelutil v0.3.2 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.68.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/vektah/gqlparser/v2 v2.5.31 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/xo/dburl v0.23.8 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
	go.etcd.io/etcd/api/v3 v3.6.7 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.6.7 // indirect
	go.mongodb.org/mongo-driver v1.17.6 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.63.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/runtime v0.64.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp v0.15.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v1.38.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v1.39.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.39.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.39.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.39.0 // indirect
	go.opentelemetry.io/otel/log v0.15.0 // indirect
	go.opentelemetry.io/otel/metric v1.39.0 // indirect
	go.opentelemetry.io/otel/sdk v1.39.0 // indirect
	go.opentelemetry.io/otel/sdk/log v0.15.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.39.0 // indirect
	go.opentelemetry.io/proto/otlp v1.9.0 // indirect
	go.temporal.io/api v1.59.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/goleak v1.3.0 // indirect
	go.yaml.in/yaml/v2 v2.4.3 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	go4.org v0.0.0-20230225012048-214862532bf5 // indirect
	golang.org/x/mod v0.32.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/time v0.14.0 // indirect
	golang.org/x/tools v0.41.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251222181119-0a764e51fe1b // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/apimachinery v0.35.0 // indirect
	k8s.io/klog/v2 v2.130.1 // indirect
	k8s.io/kube-openapi v0.0.0-20250910181357-589584f1c912 // indirect
	k8s.io/utils v0.0.0-20251002143259-bc988d571ff4 // indirect
	modernc.org/libc v1.67.0 // indirect
	modernc.org/mathutil v1.7.1 // indirect
	modernc.org/memory v1.11.0 // indirect
	modernc.org/sqlite v1.40.1 // indirect
	sigs.k8s.io/json v0.0.0-20250730193827-2d320260d730 // indirect
	sigs.k8s.io/randfill v1.0.0 // indirect
	sigs.k8s.io/structured-merge-diff/v6 v6.3.0 // indirect
)

//replace github.com/blink-io/sq => /home/heisonyee/Data/projects/go/sq
//replace github.com/blink-io/sq => /Users/heisonyee/Projects/Go/sq
