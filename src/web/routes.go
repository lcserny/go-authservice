package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lcserny/go-authservice/src/auth"
	"github.com/lcserny/go-authservice/src/config"
	"github.com/lcserny/go-authservice/src/users"
)

func MountRoutes(cfg *config.Config, r chi.Router) {
	r.Mount("/users", userRoutes(cfg))
	r.Mount("/auth", authRoutes(cfg))
}

func userRoutes(cfg *config.Config) *chi.Mux {
	ctrl := users.NewUsersController(cfg)
	r := chi.NewRouter()
	r.Get("/{userId}", ctrl.GetUserAPI)
	return r
}

func authRoutes(cfg *config.Config) *chi.Mux {
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
