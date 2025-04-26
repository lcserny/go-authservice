package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lcserny/go-authservice/src/auth"
	"github.com/lcserny/go-authservice/src/config"
	"github.com/lcserny/go-authservice/src/db"
	"github.com/lcserny/go-authservice/src/users"
)

func MountRoutes(cfg *config.Config, repoProdiver db.RepositoryProvider, r chi.Router) {
	r.Mount("/users", userRoutes(cfg, repoProdiver.GetUserRepository()))
	r.Mount("/auth", authRoutes(cfg, repoProdiver.GetAuthRepository()))
}

func userRoutes(cfg *config.Config, repo users.UserRepository) *chi.Mux {
	service := users.NewUserService(cfg, repo)
	ctrl := users.NewUsersController(cfg, service)
	r := chi.NewRouter()
	r.Get("/{userId}", ctrl.GetUserAPI)
	return r
}

func authRoutes(cfg *config.Config, repo auth.AuthRepository) *chi.Mux {
	// service := auth.NewAuthService(repo)
	ctrl := auth.NewAuthController(cfg)
	r := chi.NewRouter()
	r.With(authMiddleware).Get("/", ctrl.SignIn)
	return r
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
