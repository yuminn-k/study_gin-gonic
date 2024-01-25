package controllers

import (
	"gin-fleamarket/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// IItemController IItemController 인터페이스는 ItemController 구조체가 구현해야 하는 메서드를 정의한다.
type IItemController interface {
	FindAll(ctx *gin.Context)
}

// ItemController ItemController 구조체는 IItemController 인터페이스를 구현한다.
type ItemController struct {
	service services.IItemService
}

// NewItemController NewItemController() 메서드는 ItemController 구조체를 생성한다.
func NewItemController(service services.IItemService) IItemController {
	return &ItemController{service: service}
}

// FindAll FindAll() 메서드는 모든 Item을 반환한다.
func (c *ItemController) FindAll(ctx *gin.Context) {
	items, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error occurred"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": items})
}
