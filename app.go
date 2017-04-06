package app

import (
	"gopkg.in/gin-gonic/gin.v1"
	//"net/http"
	signIn "src/server/api/handlers/sign_in"
	signUp "src/server/api/handlers/sign_up"

	ping "src/server/api/handlers/ping"
	//"src/server/api"
	"net/http"
	"src/server/api/handlers/auth"
	validate "src/server/api/validators"

	r "src/server/router"
	//"gopkg.in/go-playground/validator.v8"
)

func init() {

	router := gin.Default()
	//router := gin.New()

	//router.Use(gin.Logger())
	//router.Use(gin.Recovery())
	router.Use(validate.Client())
	v1 := router.Group("/api/v1")
	{
		pingV1 := v1.Group("/ping")
		{
			pingV1.POST("/test", r.Session(ping.Test))
		}

		authV1 := v1.Group("/auth")
		{
			authV1.POST("/fb", auth.FB)
			authV1.POST("/sign_in", signIn.POST)
			authV1.POST("/sign_up", signUp.POST)
			authV1.POST("/recovery", signUp.POST)
		}

		authTokenV1 := v1.Group("/auth")
		authTokenV1.Use(validate.Session())
		{
			authTokenV1.POST("/sign_out", signUp.POST)
			authTokenV1.POST("/reset", signUp.POST)
		}

		sysV1 := v1.Group("/sys")
		sysV1.Use(validate.Session())
		{
			sysV1.POST("/fb", auth.FB)
			sysV1.POST("/sign_in", signIn.POST)
		}


		usersV1 := v1.Group("/users")
		usersV1.Use(validate.Session())
		{
			usersV1.GET("/:id", signUp.POST)
			usersV1.GET("/", signUp.POST)
			usersV1.PUT("/", signUp.POST)
			usersV1.POST("/", signUp.POST)
			usersV1.DELETE("/", signUp.POST)
		}
		adsV1 := v1.Group("/ads")
		adsV1.Use(validate.Session())
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
