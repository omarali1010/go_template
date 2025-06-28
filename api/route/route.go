package route

import (
	"time"

	"github.com/go-chi/chi/v5"
	middleware "github.com/omaraliali1010/go_template/api/middelware"
	"github.com/omaraliali1010/go_template/bootstrap"
)

func Setup(app *bootstrap.Application, timeout time.Duration, r *chi.Mux) {
	// Public APIs
	r.Group(func(r chi.Router) {
		NewSignupRouter(app, timeout, r)
		NewLoginRouter(app, timeout, r)
		NewRefreshTokenRouter(app, timeout, r)
	})

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.JwtAuthMiddleware(app.JWTService))
		NewProfileRouter(app, timeout, r)
	})
}
