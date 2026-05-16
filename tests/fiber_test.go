package tests

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
	"testing"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/envvar"
	expvarmw "github.com/gofiber/fiber/v3/middleware/expvar"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/paginate"
	"github.com/gofiber/fiber/v3/middleware/pprof"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/gofiber/fiber/v3/middleware/responsetime"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/require"
)

func init() {
}

func setupRouters(r fiber.Router) {
	r.Get("/", func(c fiber.Ctx) {
		_, _ = c.WriteString("OK")
	})

	rg := r.Group("/api")

	// Use the default probe on the conventional endpoints
	rg.Get(healthcheck.LivenessEndpoint, healthcheck.New())
	rg.Get(healthcheck.ReadinessEndpoint, healthcheck.New(healthcheck.Config{
		Probe: func(c fiber.Ctx) bool {
			return true
		},
	}))
	rg.Get(healthcheck.StartupEndpoint, healthcheck.New())

	rg.Get("/", func(c fiber.Ctx) error {
		sess := session.FromContext(c)
		sess.Set("user_id", 123)
		sess.Set("authenticated", true)
		fmt.Println(sess)
		req := c.Request()
		log.Info("req: " + req.String())
		return nil
	})

	rg.Get("/express", func(req fiber.Req, res fiber.Res) error {
		return res.SendString("Hello from Express-style handlers!")
	})

	rg.Get("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "Hello from net/http!")
	}))

	rg.Get("/hi", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "Hi!")
	})

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "Hello from net/http!")
	}
	rg.Get("/hello/:name", adaptor.HTTPHandlerWithContext(http.HandlerFunc(helloHandler)))

	rg.Get("/users", func(c fiber.Ctx) error {
		pageInfo, ok := paginate.FromContext(c)
		if !ok {
			return fiber.ErrBadRequest
		}

		// Use pageInfo.Page, pageInfo.Limit, pageInfo.Start()
		// GET /users?page=2&limit=20 → Page: 2, Limit: 20, Start(): 20
		return c.JSON(pageInfo)
	})

	rg.Get("/json", func(c fiber.Ctx) error {
		sess := session.FromContext(c)
		v1 := sess.Get("user_id")
		v2 := sess.Get("authenticated")
		fmt.Println(v1, v2)

		action := c.Params("action")
		req := c.Request()
		log.Info("req: " + req.String())
		fmt.Println(action)
		c.JSON(fiber.Map{
			"hello": "world",
			"type":  "info",
		})
		return nil
	})
}

func setupApp(app *fiber.App) {

}

func TestFiber_1(t *testing.T) {
	cfg := &fiber.Config{
		ServerHeader: "my-fiber-server",
		AppName:      "Test Fiber Server",
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ServicesStartupContextProvider: func() context.Context {
			return context.Background()
		},
		ServicesShutdownContextProvider: func() context.Context {
			return context.Background()
		},
	}

	// Initialize service.
	cfg.Services = append(cfg.Services, &myService{})

	logz := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app := fiber.New(*cfg)
	app.Use(expvarmw.New())
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/New_York",
	}))
	app.Use("/expose/envvars", envvar.New())
	// Custom File Writer
	accessLog, err := os.OpenFile("./access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening access.log file: %v", err)
	}
	defer accessLog.Close()
	app.Use(logger.New(logger.Config{
		Stream: accessLog,
	}))
	app.Use(helmet.New())
	app.Use(pprof.New(pprof.Config{Prefix: "/endpoint-prefix"}))
	app.Use(recover.New())
	app.Use(requestid.New(requestid.Config{
		Generator: func() string {
			switch rand.Intn(3) {
			case 0:
				return ulid.Make().String()
			case 1:
				return ksuid.New().String()
			default:
				return uuid.New().String()
			}
		},
	}))
	app.Use(compress.New())
	app.Use(session.New())
	app.Use(responsetime.New())
	app.Use(paginate.New(paginate.Config{
		DefaultPage:  1,
		DefaultLimit: 20,
		SortKey:      "sort",
		DefaultSort:  "id",
		LimitKey:     "perPage",
		//AllowedSorts: []string{"id", "name", "created_at"},
	}))

	app.Hooks().OnListen(func(listenData fiber.ListenData) error {
		if fiber.IsChild() {
			return nil
		}
		scheme := "http"
		if listenData.TLS {
			scheme = "https"
		}
		logz.Info(scheme + "://" + listenData.Host + ":" + listenData.Port)
		return nil
	})
	app.Hooks().OnRoute(func(r fiber.Route) error {
		logz.Info("Route info",
			slog.String("name", r.Name),
			slog.String("method", r.Method),
			slog.String("path", r.Path),
		)
		return nil
	})

	setupRouters(app)

	err = app.Listen(":14004")
	require.NoError(t, err)
}

var _ fiber.Service = (*myService)(nil)

type myService struct{}

func (m *myService) Start(ctx context.Context) error {
	return nil
}

func (m *myService) String() string {
	return "my-svc"
}

func (m *myService) State(ctx context.Context) (string, error) {
	//state, err := m.ctr.State(ctx)
	//if err != nil {
	//	return "", fmt.Errorf("container state: %w", err)
	//}
	return "running", nil
}

func (m *myService) Terminate(ctx context.Context) error {
	return nil
}
