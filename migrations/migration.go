package main

import (
	"gin-fleamarket/infra"
	"gin-fleamarket/models"
)

//マイグレーションは、独立して行うため、mainで定義している
func main() {
	infra.Initialize() //.env ファイルから環境変数を読み込み、アプリケーションにロードするための初期化処理を行う。

	db := infra.SetupDB() //データベース接続を設定し、*gorm.DB オブジェクトを返す。このオブジェクトは、データベース操作を行うためのインターフェースを提供。

	// AutoMigrate:構造体を引数として渡し、構造体に定義されているフィールドに基づいて、データベースにテーブルを作成、更新
	if err := db.AutoMigrate(&models.Item{}, &models.User{}); err != nil {
		panic("Failed to migrate database")
	}
}