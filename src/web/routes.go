package web

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/lcserny/go-authservice/src/generated"
	"github.com/lcserny/go-authservice/src/logging"
	slogchi "github.com/samber/slog-chi"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lcserny/go-authservice/src/config"
	"github.com/lcserny/go-authservice/src/db"
	"github.com/lcserny/go-authservice/src/users"
)

func NewRouter(cfg *config.Config, repoProvider db.RepositoryProvider) chi.Router {
	var controllers []generated.Router
	controllers = append(controllers, userControllers(cfg, repoProvider)...)
	controllers = append(controllers, authControllers(cfg, repoProvider)...)

	r := chi.NewRouter()
	r.Use(slogchi.New(logging.Logger()))
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	}))

	for _, api := range controllers {
		for _, route := range api.Routes() {
			var handler http.Handler = route.HandlerFunc
			r.Method(route.Method, route.Pattern, handler)
		}
	}

	return r
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
