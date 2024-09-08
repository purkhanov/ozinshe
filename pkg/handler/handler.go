package handler

import (
	"net/http"
	"ozinshe/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Static("/movies", "./data/movies")
	router.Static("/images", "./data/images")
	router.Static("/screenshots", "./data/screenshots")

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"pong": "pong"})
	})

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	apiV1 := router.Group("/api/v1")
	{
		movies := apiV1.Group("/movies")
		{
			movies.GET("/", h.getMovies)
			movies.GET("/:id", h.getMovieById)
		}

		admin := apiV1.Group("/admin", h.adminIdentity)
		{
			adminUsers := admin.Group("/users")
			{
				adminUsers.GET("/", h.getAllUsers)
				adminUsers.GET("/:id", h.getUserById)
				adminUsers.PUT("/:id", h.updateUserById)
				adminUsers.DELETE("/:id", h.deleteUserById)
			}

			adminMovies := admin.Group("/movies")
			{
				adminMovies.GET("/", h.getAllMovies)
				adminMovies.GET("/:id", h.getMovieById)
				adminMovies.POST("/", h.addMovieInfo)
				adminMovies.POST("/:id", h.uploadMovie)
				adminMovies.PUT("/:id", h.updateMovie)
				adminMovies.DELETE("/:id", h.deleteMovie)

				adminMovieScreens := admin.Group("/screenshots")
				{
					adminMovieScreens.GET("/:id", h.getScreenshots)
					adminMovieScreens.POST("/:id", h.uploadScreenshot)
					adminMovieScreens.DELETE("/:id", h.deleteScreenshot)
				}
			}

			adminGenres := admin.Group("/genres")
			{
				adminGenres.GET("/", h.getAllGenres)
				adminGenres.POST("/", h.addGenre)
				adminGenres.PUT("/:id", h.updateGenre)
				adminGenres.DELETE("/:id", h.deleteGenre)
			}
		}

		user := apiV1.Group("/user", h.userIdentity)
		{
			user.GET("/", h.getUser)
			user.PUT("/", h.updateUser)
			user.DELETE("/", h.deleteUser)

			userMovies := user.Group("/favorites")
			{
				userMovies.GET("/", h.getFavoriteMovies)
				userMovies.POST("/:id", h.addFavoriteMovie)
				userMovies.DELETE("/:id", h.deleteFavoriteMovie)
			}

			userWatchedMovies := user.Group("/watched-movies")
			{
				userWatchedMovies.GET("/", h.getWtchedMovies)
				userWatchedMovies.POST("/:id", h.addWtchedMovies)
				userWatchedMovies.DELETE("/:id", h.deleteWtchedMovies)
			}
		}
	}

	return router
}
