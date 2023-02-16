package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilherm5/crud/models"
)

func (h Handler) List(c *gin.Context) {
	var list []models.Tarefas
	h.DB.Find(&list)

	list = append(list, models.Tarefas{})
	c.HTML(http.StatusOK, "list.html", gin.H{
		"list": list,
	})
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")

}
