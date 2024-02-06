package controllers

import (
	"gin-fleamarket/dto"
	"gin-fleamarket/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// IItemController IItemController 인터페이스는 ItemController 구조체가 구현해야 하는 메서드를 정의한다.
type IItemController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
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

// FindById FindById() 메서드는 itemId에 해당하는 Item을 반환한다.
func (c *ItemController) FindById(ctx *gin.Context) {
	itemId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
		return
	}

	item, err := c.service.FindById(uint(itemId))
	if err != nil {
		if err.Error() == "Item not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error occurred"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": item})
}

// Create Create() 메서드는 Item을 생성한다.
func (c *ItemController) Create(ctx *gin.Context) {
	var input dto.CreateItemInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newItem, err := c.service.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": newItem})
}

// Update Update() 메서드는 Item을 수정한다.
func (c *ItemController) Update(ctx *gin.Context) {
	itemId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
		return
	}

	var input dto.UpdateItemInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedItem, err := c.service.Update(uint(itemId), input)
	if err != nil {
		if err.Error() == "Item not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error occurred"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": updatedItem})
}

// Delete Delete() 메서드는 Item을 삭제한다.
func (c *ItemController) Delete(ctx *gin.Context) {
	itemId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
		return
	}

	err = c.service.Delete(uint(itemId))
	if err != nil {
		if err.Error() == "Item not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error occurred"})
		return
	}

	ctx.Status(http.StatusOK)
}
