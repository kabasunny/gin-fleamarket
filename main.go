package main

import (
	"gin-fleamarket/controllers"
	"gin-fleamarket/infra"

	// "gin-fleamarket/models"
	"gin-fleamarket/repositories"
	"gin-fleamarket/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	// items := []models.Item{
	// 	{ID: 1, Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
	// 	{ID: 2, Name: "商品2", Price: 2000, Description: "説明2", SoldOut: true},
	// 	{ID: 3, Name: "商品3", Price: 3000, Description: "説明3", SoldOut: false},
	// }
	// itemRepository := repositories.NewItemMemoryRepository(items)
	itemRepository := repositories.NewItemRepository(db) //メモリからDBへ切替
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	r := gin.Default() //ginのデフォルトルータを初期化
	itemRouter := r.Group("/items")
	authRouter := r.Group("/auth")

	//以下は、各リクエストにおける、エンドポイントとハンドラ関数の紐づけ設定を行っている
	//ハンドラ関数は、実際にリクエストがあった際に発動する
	itemRouter.GET("", itemController.FindAll)
	itemRouter.GET("/:id", itemController.FindById)
	itemRouter.POST("", itemController.Create)
	itemRouter.PUT("/:id", itemController.Update)
	itemRouter.DELETE("/:id", itemController.Delete)

	authRouter.POST("/signup", authController.Signup)

	//以下は、サーバーを立ち上げ（以降、リクエストの待機、レスポンス送信行う）
	r.Run("localhost:8080")
}