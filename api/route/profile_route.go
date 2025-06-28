package route

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/omaraliali1010/go_template/api/controller"
	"github.com/omaraliali1010/go_template/bootstrap"
	"github.com/omaraliali1010/go_template/repository"
	"github.com/omaraliali1010/go_template/usecase"
)

func NewProfileRouter(app *bootstrap.Application, timeout time.Duration, router chi.Router) {
	userRepo := repository.NewUserRepository(app.DB)

	profileUC := usecase.NewProfileUsecase(userRepo, timeout)

	sc := &controller.ProfileController{
		ProfileUsecase: profileUC,
	}

	router.Get("/protected/profile", sc.Fetch)

}
