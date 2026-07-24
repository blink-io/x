package humah3

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/go-kratos/kratos/v3"
	khttp "github.com/go-kratos/kratos/v3/transport/http"
)

func TestServer(t *testing.T) {
	// Options for the CLI.
	type Options struct {
		Port int `help:"Port to listen on" short:"p" default:"9988"`
	}

	// GreetingOutput represents the greeting operation response.
	type GreetingOutput struct {
		Body struct {
			Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
		}
	}

	addRoutes := func(api huma.API) {
		// Register GET /greeting/{name}
		huma.Register(api, huma.Operation{
			OperationID: "get-greeting",
			Method:      http.MethodGet,
			Path:        "/greeting/{name}",
			Summary:     "获取问候语",
			Description: "Get a greeting for a person by name.",
			Tags:        []string{"问候类"},
		}, func(ctx context.Context, input *struct {
			Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
		}) (*GreetingOutput, error) {
			resp := &GreetingOutput{}
			resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
			return resp, nil
		})
	}

	// Create a CLI app which takes a port option.
	cli := humacli.New(func(hooks humacli.Hooks, opts *Options) {
		// Create a new router & API
		srv := khttp.NewServer(khttp.Address(":" + strconv.Itoa(opts.Port)))
		kapp := kratos.New(kratos.Server(srv))

		api := New(
			srv,
			huma.DefaultConfig("My Kratos API", "1.0.0"),
		)

		addRoutes(api)

		hooks.OnStart(func() {
			fmt.Printf("Starting server on port %d...\n", opts.Port)
			if err := kapp.Run(); err != nil {
				fmt.Printf("server start error: %v\n", err)
			}
		})

		// Tell the CLI how to stop your server.
		hooks.OnStop(func() {
			if err := kapp.Stop(); err != nil {
				fmt.Printf("server stop error: %v\n", err)
			}
		})
	})

	// Run the CLI. When passed no commands, it starts the server.
	cli.Run()
}
