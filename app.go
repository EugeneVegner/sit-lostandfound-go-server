package app

import (
	"gopkg.in/gin-gonic/gin.v1"
	signIn "src/server/api/handlers/sign_in"
	signUp "src/server/api/handlers/sign_up"
	"net/http"
)

func init() {

	router := gin.Default()
	router.POST("/api/v1/sign_upq", signUp.POST)

	api := router.Group("/api",)
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/sign_in", signIn.POST)
			v1.POST("/sign_up", signUp.POST)
		}


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
	http.Handle("/", router)
}
