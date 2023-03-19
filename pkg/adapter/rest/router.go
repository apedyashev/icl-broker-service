package rest

import (
	"icl-broker/pkg/adapter/rest/controller"
	appMiddleware "icl-broker/pkg/adapter/rest/middleware"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo) *echo.Echo {
	jwtSecret := os.Getenv("JWT_SECRET")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// auth routes
	authController := controller.NewAuthController()
	g := e.Group("/api/auth")
	g.POST("/login", authController.Login)
	g.POST("/register", authController.Register)

	postController := controller.NewPostController()
	e.GET("/api/users/:userId/posts", postController.UserPosts)
	e.GET("/api/posts/:id", postController.PostById)

	e.POST("/api/posts", postController.Create)
	e.POST("/api/posts/images", postController.AddImageToPost)
	e.POST("/api/posts/:id/images", postController.AddImageToPost)

	// protected routes
	userController := controller.NewUserController()
	prg := e.Group("/api/auth")
	{
		prg.Use(echojwt.JWT([]byte(jwtSecret)))
		prg.Use(appMiddleware.AddClaims())

		prg.GET("/user", userController.CurrentUser)
	}

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
