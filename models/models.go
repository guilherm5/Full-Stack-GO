package models

type Tarefas struct {
	ID              int    `json:"id" gorm:"primaryKey"`
	NomeTarefa      string `json:"nome_tarefa"`
	DescricaoTarefa string `json:"descricao_tarefa"`
	QuandoFazer     string `json:"quando_fazer"`
}
