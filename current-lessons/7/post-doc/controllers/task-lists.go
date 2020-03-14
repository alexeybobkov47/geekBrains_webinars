package controllers

import (
	"not-for-work/GeekBrainsWebinars/current-lessons/7/post-doc/models"

	"github.com/astaxie/beego"
)

type ListController struct {
	beego.Controller
	Explorer Explorer
}

// @Title GetAllLists
// @Description Вернет страницу со статистикой по задачам
// @Success 200 {html}
// @Failure 500 error
// @router / [get]
func (c *ListController) Get() {
	lists, err := c.Explorer.getAllLists()
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	c.Data["Lists"] = lists
	c.TplName = "lists.tpl"
}

type PostRequest struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
}

/*
	curl -vX POST -H "Content-Type: application/json"  -d'{"name":"NewTextMongo","desc":"newDesc"}' http://localhost:8090/lists
*/

// @Title CreateList
// @Description Создвание нового листа
// @Param  name      path   string true      "Имя листа"
// @Param  desc      path   string true      "Описание листа"
// @Success 200 SUCCESS
// @Failure 500 error
// @router / [post]
func (c *ListController) Post() {
	resp := new(PostRequest)
	if err := readAndUnmarshall(resp, c.Ctx.Request.Body); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	list := models.List{
		Name:        resp.Name,
		Description: resp.Description,
	}

	if err := c.Explorer.createList(list); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
	_, _ = c.Ctx.ResponseWriter.Write([]byte(`SUCCESS`))
}

/*
	curl -vX PUT -H "Content-Type: application/json"  -d'{"name":"NewTest_put","desc":"newDesc_put"}' http://localhost:8090/lists?list_name=Test
*/

// @Title UpdateList
// @Description Обновление листа
// @Param  listname      path   string true      "Текущее название листа"
// @Param  name      path   string true      "Новое название"
// @Param  desc      path   string true      "Новое описание"
// @Success 200 SUCCESS
// @Failure 500 error
// @router / [put]
func (c *ListController) Put() {
	listName := c.Ctx.Request.URL.Query().Get("list_name")

	if len(listName) == 0 {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte("empty list_name"))
		return
	}

	resp := new(PostRequest)
	if err := readAndUnmarshall(resp, c.Ctx.Request.Body); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	list := models.List{
		Name:        resp.Name,
		Description: resp.Description,
	}

	if err := c.Explorer.updateList(listName, list.Name, list.Description); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
	_, _ = c.Ctx.ResponseWriter.Write([]byte(`SUCCESS`))
}

/*
	curl -vX DELETE  http://localhost:8090/lists?list_name=NewTest_put
*/

// @Title DeleteList
// @Description Удаление листа
// @Param  listname      path   string true      "Название листа"
// @Success 200 body is empty
// @Failure 500 error
// @router / [delete]
func (c *ListController) Delete() {
	listName := c.Ctx.Request.URL.Query().Get("list_name")
	if len(listName) == 0 {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte("empty list_name"))
		return
	}

	err := c.Explorer.deleteList(listName)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
	_, _ = c.Ctx.ResponseWriter.Write([]byte(`SUCCESS`))
}
