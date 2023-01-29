package main

import (
	"fmt"
	"gin/basic/db"
	"gin/basic/router"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	r := gin.Default()
	_, err := db.GetPostgresDBConnection()
	if err != nil {
		log.Fatal("Unable to establish DB connection", err)
	} else {
		log.Println("DB connection establsihed")
	}

	router.PostRouter(r)
	router.UserRouter(r)
	r.Run(fmt.Sprintf(":%s", port))

}
