package main

import (
	"fmt"

	"github.com/RaymondCode/simple-demo/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	initRouter(r)
	err := repository.Init()
	if err != nil {
		fmt.Println("error in main db init")
		return
	} else {
		fmt.Println("db init successfully")
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
