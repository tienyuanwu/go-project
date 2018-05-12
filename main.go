package main

import (
	//"./db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	/*
		err := db.InitDb()
		if err != nil {
			fmt.Println(err)
			return
		}
	*/

	router := gin.Default()
	router.Use(cors.Default())
	InitRouter(router)
	router.Run(":8080")
}
