package controllers

import (
	"github.com/revel/revel"
	"github.com/sliceking/todo/app/models"
)

type Todo struct {
	App
}

func (c Todo) Index() revel.Result {
	c.Txn.Save(&models.Todo{Description: "Fidget spinner", Done: false})
	var todos []*models.Todo
	c.Txn.Find(&todos)
	return c.Render(todos)
}
