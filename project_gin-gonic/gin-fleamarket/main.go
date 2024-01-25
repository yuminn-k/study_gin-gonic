package main

import (
	"gin-fleamarket/controllers"
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
	"gin-fleamarket/services"
	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "商品１", Price: 1000, Description: "説明１", SoldOut: false},
		{ID: 2, Name: "商品２", Price: 2000, Description: "説明２", SoldOut: true},
		{ID: 3, Name: "商品３", Price: 3000, Description: "説明３", SoldOut: false},
	}

	itemRepository := repositories.NewItemMemoryRepository(items) // ItemMemoryRepository 구조체를 생성한다.
	itemService := services.NewItemService(itemRepository)        // ItemService 구조체를 생성한다.
	itemController := controllers.NewItemController(itemService)  // ItemController 구조체를 생성한다.

	r := gin.Default()                      // gin.Default()는 기본적으로 Logger와 Recovery 미들웨어를 포함한 Gin 엔진을 생성한다. HTTP 요청을 받고 응답을 전송하는 역할을 한다.
	r.GET("/items", itemController.FindAll) // GET 요청을 받고, "/items" 경로로 요청이 들어오면, itemController.FindAll() 메서드를 실행한다.
	r.Run("localhost:8080")                 // 8080 포트로 서버를 실행한다.
}
