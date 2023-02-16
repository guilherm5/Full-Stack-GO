package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilherm5/crud/models"
)

func (h Handler) Update(c *gin.Context) {
	var tarefa models.Tarefas
	id := c.Param("id")
	if err := h.DB.Where("id=?", id).First(&tarefa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tarefa não encontrada"})
		return
	}

	if c.Request.Method == http.MethodPost {
		c.Bind(&tarefa)
		h.DB.Save(&tarefa)
		c.Redirect(http.StatusFound, "/list")
		return
	}

	c.HTML(http.StatusOK, "edit", gin.H{"tarefa": tarefa})

	log.Println(tarefa.ID)
}
