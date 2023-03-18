package main

import (
	"icl-broker/pkg/adapter/api"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	postController := api.NewPostController()
	e.GET("/api/users/:userId/posts", postController.UserPosts)
	e.GET("/api/posts/:id", postController.PostById)

	e.POST("/api/posts", postController.Create)
	e.POST("/api/posts/images", postController.AddImageToPost)
	e.POST("/api/posts/:id/images", postController.AddImageToPost)

	return e
}

// package main

// import (
// 	"net/http"

// 	"icl-broker/cmd/api/handlers"
// 	"icl-broker/cmd/api/helpers"
// 	"icl-broker/cmd/api/middlewares"

// 	"github.com/go-chi/chi/v5"
// 	"github.com/go-chi/chi/v5/middleware"
// 	"github.com/go-chi/cors"
// 	"github.com/go-chi/jwtauth"
// )

// // receiver of app, which is type pointer to config
// func (app *Config) routes() http.Handler {
// 	mux := chi.NewRouter()

// 	// specify who is allowed to connect
// 	mux.Use(cors.Handler(cors.Options{
// 		AllowedOrigins:   []string{"https://*", "http://*"},
// 		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
// 		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
// 		ExposedHeaders:   []string{"Link"},
// 		AllowCredentials: true,
// 		MaxAge:           300,
// 	}))

// 	mux.Use(middleware.Heartbeat("/ping"))

// 	// auth routes
// 	mux.Route("/api/auth", func(r chi.Router) {
// 		r.Post("/login", handlers.Login)
// 		r.Post("/register", handlers.Register)

// 	})

// 	// Protected routes
// 	mux.Group(func(r chi.Router) {
// 		// Seek, verify and validate JWT tokens
// 		r.Use(jwtauth.Verifier(helpers.TokenAuth))

// 		// Handle valid / invalid tokens. In this example, we use
// 		// the provided authenticator middleware, but you can write your
// 		// own very easily, look at the Authenticator method in jwtauth.go
// 		// and tweak it, its not scary.
// 		r.Use(jwtauth.Authenticator)
// 		r.Use(middlewares.IsUserEnabled)

// 		r.Get("/api/auth/user", handlers.GetCurrentUser)

// 	})
// 	mux.Post("/api/image/upload", handlers.UploadImage)

// 	return mux
// }
