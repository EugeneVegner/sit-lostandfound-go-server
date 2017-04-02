package app

import (
	"gopkg.in/gin-gonic/gin.v1"
	//"net/http"
	signIn "src/server/api/handlers/sign_in"
	signUp "src/server/api/handlers/sign_up"
	//"src/server/api"
	"net/http"
	"src/server/api/handlers/auth"
	"src/server/api/validators"
)

func init() {

	router := gin.Default()
	//router := gin.New()

	//router.Use(gin.Logger())
	//router.Use(gin.Recovery())
	router.Use(validator.Client())
	//router.Use(Logger())

	v1 := router.Group("/api/v1")
	{
		authV1 := v1.Group("/auth")
		{
			authV1.POST("/fb", auth.FB)
			authV1.POST("/sign_in", signIn.POST)
			authV1.POST("/sign_up", signUp.POST)
			authV1.POST("/recovery", signUp.POST)
			authV1.POST("/reset", signUp.POST)
		}

		usersV1 := v1.Group("/users")
		usersV1.Use(validator.Token())
		{
			usersV1.GET("/:id", signUp.POST)
			usersV1.GET("/", signUp.POST)
			usersV1.PUT("/", signUp.POST)
			usersV1.POST("/", signUp.POST)
			usersV1.DELETE("/", signUp.POST)
		}
		adsV1 := v1.Group("/ads")
		adsV1.Use(validator.Token())
		{
			adsV1.GET("/:id", signUp.POST)
			adsV1.GET("/", signUp.POST)
			adsV1.PUT("/", signUp.POST)
			adsV1.POST("/", signUp.POST)
			adsV1.DELETE("/", signUp.POST)
		}
	}
	dev := router.Group("/api/dev")
	{
		dev.POST("/sign_in", signIn.POST)
	}

	//
	//v1 := router.Group("/v1") {
	//	v3 := v1.Group("")
	//
	//
	//
	//	v1.POST("/login", loginEndpoint)
	//	v1.POST("/submit", submitEndpoint)
	//	v1.POST("/read", readEndpoint)
	//}

	//router.GET("/someGet", getting)
	//router.POST("/somePost", posting)
	//router.PUT("/somePut", putting)
	//router.DELETE("/someDelete", deleting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	//router.Run(":8080")

	//router := src.ConfigureRoutes()
	//router.Run()
	http.Handle("/", router)
}
