package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/crud/db"
)

func main() {
	dbLoad := db.Init()
	h := db.LoadDB(dbLoad)

	router := gin.Default()

	router.LoadHTMLGlob("public/*.html")
	router.Static("/create", "./public")
	router.Static("/css", "./css")
	router.Static("/js", "./js")

	router.GET("/edit/:id", h.Update)
	router.GET("/delete/:id", h.Delete)
	router.POST("/create", h.Create)
	router.GET("/list", h.List)
	router.PUT("/edit/:id", h.Update)

	router.Run(":8080")
}
