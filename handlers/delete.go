package handlers

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/guilherm5/crud/models"
)

func (h Handler) Delete(c *gin.Context) {

	var deleted models.Tarefas

	if err := h.DB.Where("id = ?", c.Param("id")).First(&deleted).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		log.Println(err)

	}
	h.DB.Delete(&deleted)

	//c.Redirect(deleted.IDProd, "localhost:8080/list")
	q := url.Values{}

	location := url.URL{Path: "/list", RawQuery: q.Encode()}
	c.Redirect(http.StatusFound, location.RequestURI())

}
