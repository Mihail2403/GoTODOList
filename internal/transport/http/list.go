package http_handler

import (
	"go_todo_list/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	id, err := getUserID(c)
	if err != nil {
		return
	}
	var input entity.TodoList
	err = c.Bind(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	list_id, err := h.services.TodoList.Create(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": list_id,
	})
}

type getAllListsResponse struct {
	Data []entity.TodoList `json:"data"`
}

func (h *Handler) getAllList(c *gin.Context) {
	id, err := getUserID(c)
	if err != nil {
		return
	}
	lists, err := h.services.TodoList.GetAll(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK,
		getAllListsResponse{
			Data: lists,
		},
	)

}

func (h *Handler) getListById(c *gin.Context) {
	id, err := getUserID(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id must be a number")
		return
	}
	list, err := h.services.TodoList.GetById(id, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(c *gin.Context) {
	id, err := getUserID(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id must be a number")
		return
	}
	var input entity.UpdateTodoListInput
	err = c.Bind(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.TodoList.Update(id, listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "")
}

func (h *Handler) deleteList(c *gin.Context) {
	id, err := getUserID(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id must be a number")
		return
	}
	err = h.services.TodoList.Delete(id, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "")
}
