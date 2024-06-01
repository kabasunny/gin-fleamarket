package infra

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// GORMは、Go言語で開発されたオープンソースのORM (Object-Relational Mapping) ライブラリで、データベースとGo言語の構造体をマッピングすることで、データベース操作を簡単かつ効率的に行うことができる
)

func SetupDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"),
	)
	// セキュリティ上公開すべきでない情報を環境変数として設定し、os.Getenvを使って安全にアクセスすることができる

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		panic("Failed to connect database")
	}
	return db
}