package main

import (
	"gin-fleamarket/controllers"
	"gin-fleamarket/infra"
	//"gin-fleamarket/models"
	"gin-fleamarket/repositories"
	"gin-fleamarket/services"
	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()    // .env 파일을 로드한다.
	db := infra.SetupDB() // DB 연결을 설정한다.

	//items := []models.Item{
	//	{ID: 1, Name: "商品１", Price: 1000, Description: "説明１", SoldOut: false},
	//	{ID: 2, Name: "商品２", Price: 2000, Description: "説明２", SoldOut: true},
	//	{ID: 3, Name: "商品３", Price: 3000, Description: "説明３", SoldOut: false},
	//}

	//itemRepository := repositories.NewItemMemoryRepository(items) // ItemMemoryRepository 구조체를 생성한다.
	itemRepository := repositories.NewItemRepository(db)         // ItemGormRepository 구조체를 생성한다.
	itemService := services.NewItemService(itemRepository)       // ItemService 구조체를 생성한다.
	itemController := controllers.NewItemController(itemService) // ItemController 구조체를 생성한다.

	r := gin.Default() // gin.Default()는 기본적으로 Logger와 Recovery 미들웨어를 포함한 Gin 엔진을 생성한다. HTTP 요청을 받고 응답을 전송하는 역할을 한다.
	// "/items" 경로로 들어오는 요청을 처리하는 라우터 그룹을 생성한다. 그룹을 사용하면 중복되는 경로를 줄일 수 있다. 또한, 그룹에 미들웨어를 적용할 수 있다.
	itemRouter := r.Group("/items")
	itemRouter.GET("", itemController.FindAll)       // GET 요청을 받고, "/items" 경로로 요청이 들어오면, itemController.FindAll() 메서드를 실행한다.
	itemRouter.GET("/:id", itemController.FindById)  // GET 요청을 받고, "/items/:id" 경로로 요청이 들어오면, itemController.FindById() 메서드를 실행한다
	itemRouter.POST("", itemController.Create)       // POST 요청을 받고, "/items" 경로로 요청이 들어오면, itemController.Create() 메서드를 실행한다.
	itemRouter.PUT("/:id", itemController.Update)    // PUT 요청을 받고, "/items/:id" 경로로 요청이 들어오면, itemController.Update() 메서드를 실행한다.
	itemRouter.DELETE("/:id", itemController.Delete) // DELETE 요청을 받고, "/items/:id" 경로로 요청이 들어오면, itemController.Delete() 메서드를 실행한다.
	r.Run("localhost:8080")                          // 8080 포트로 서버를 실행한다.
}
