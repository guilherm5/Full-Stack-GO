package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/crud/models"
)

func (h Handler) Create(c *gin.Context) {
	var createHTML = []models.Tarefas{{
		NomeTarefa:      c.Request.FormValue("nome_tarefa"),
		DescricaoTarefa: c.Request.FormValue("descricao_tarefa"),
		QuandoFazer:     c.Request.FormValue("quando_fazer"),
	}}

	h.DB.Create(&createHTML)

	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	c.Writer.WriteString("<script>alert('Produto criado com sucesso'); window.location.href='/list';</script>")

}
