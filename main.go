package main

import (
	"gin-fleamarket/controllers"
	"gin-fleamarket/infra"
	"gin-fleamarket/middlewares"

	// "gin-fleamarket/models"
	"gin-fleamarket/repositories"
	"gin-fleamarket/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	itemRepository := repositories.NewItemRepository(db) //メモリからDBへ切替
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	r := gin.Default() //ginのデフォルトルータを初期化
	// エンドポイントのグループ化
	r.Use(cors.Default()) //CORS


	itemRouter := r.Group("/items")
	itemRouterWithAuth := r.Group("/items", middlewares.AuthMiddleware(authService))
	authRouter := r.Group("/auth")

	//以下は、各リクエストにおける、エンドポイントとハンドラ関数の紐づけ設定を行っている
	//ハンドラ関数は、実際にリクエストがあった際に発動する
	itemRouter.GET("", itemController.FindAll)
	itemRouterWithAuth.GET("/:id", itemController.FindById)
	itemRouterWithAuth.POST("", itemController.Create)
	itemRouterWithAuth.PUT("/:id", itemController.Update)
	itemRouterWithAuth.DELETE("/:id", itemController.Delete)

	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)

	return r

}


func main() {
	infra.Initialize()
	db := infra.SetupDB()
	r := setupRouter(db)

	// items := []models.Item{
	// 	{ID: 1, Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
	// 	{ID: 2, Name: "商品2", Price: 2000, Description: "説明2", SoldOut: true},
	// 	{ID: 3, Name: "商品3", Price: 3000, Description: "説明3", SoldOut: false},
	// }
	// itemRepository := repositories.NewItemMemoryRepository(items)
	
	//以下は、サーバーを立ち上げ（以降、リクエストの待機、レスポンス送信行う）
	r.Run("localhost:8080")
}