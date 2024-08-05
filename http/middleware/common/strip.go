package common

import (
	chimw "github.com/go-chi/chi/v5/middleware"
)

var StripSlashes = chimw.StripSlashes

var RedirectSlashes = chimw.RedirectSlashes
