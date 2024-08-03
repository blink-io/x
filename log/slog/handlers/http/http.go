package http

import "github.com/samber/slog-http"

type Config = sloghttp.Config

var New = sloghttp.New

var NewWithConfig = sloghttp.NewWithConfig

var NewWithFilters = sloghttp.NewWithFilters

var Recovery = sloghttp.Recovery
