package web

import (
	"github.com/lcserny/go-authservice/src/generated"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lcserny/go-authservice/src/config"
	"github.com/lcserny/go-authservice/src/db"
	"github.com/lcserny/go-authservice/src/users"
)

func InitRouter(cfg *config.Config, repoProvider db.RepositoryProvider) chi.Router {
	var controllers []generated.Router
	controllers = append(controllers, userControllers(cfg, repoProvider)...)
	controllers = append(controllers, authControllers(cfg, repoProvider)...)
	return generated.NewRouter(controllers...)
}

func userControllers(cfg *config.Config, repoProvider db.RepositoryProvider) []generated.Router {
	userService := users.NewService(cfg, repoProvider.GetUserRepository())
	return []generated.Router{
		generated.NewGetSingleUserAPIController(userService),
		generated.NewGetUsersResourceAPIController(userService),
		generated.NewCreateUserResourceAPIController(userService),
		generated.NewUpdateSingleUserDataAPIController(userService),
	}
}

func authControllers(cfg *config.Config, repoProvider db.RepositoryProvider) []generated.Router {
	return []generated.Router{
		// TODO init the rest of controllers and service
		// TODO use a XXXAPIOption to add authMiddlewareGuard
	}
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Example: Check for a valid Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Call the next handler if authorized
		next.ServeHTTP(w, r)
	})
}
