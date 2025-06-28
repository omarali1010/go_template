package route

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/omaraliali1010/go_template/api/controller"
	"github.com/omaraliali1010/go_template/bootstrap"
	"github.com/omaraliali1010/go_template/domain"
	"github.com/omaraliali1010/go_template/internal/jwtservice"
	"github.com/omaraliali1010/go_template/repository"
	"github.com/omaraliali1010/go_template/usecase"
)

// Compile-time interface assertions
// To make sure that the JWTSerivice implements the methods in the interface
var (
	_ domain.AccessTokenCreator  = (*jwtservice.JWTService)(nil)
	_ domain.RefreshTokenCreator = (*jwtservice.JWTService)(nil)
)

func NewSignupRouter(app *bootstrap.Application, timeout time.Duration, router chi.Router) {
	userRepo := repository.NewUserRepository(app.DB)

	signupUC := usecase.NewSignupUsecase(userRepo, app.JWTService, app.JWTService, timeout)

	sc := &controller.SignupController{
		SignupUsecase: signupUC,
	}

	router.Post("/public/signup", sc.Signup)
}
