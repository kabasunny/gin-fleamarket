package infra

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// GORMは、Go言語で開発されたオープンソースのORM (Object-Relational Mapping) ライブラリで、データベースとGo言語の構造体をマッピングすることで、データベース操作を簡単かつ効率的に行うことができる
)

func SetupDB() *gorm.DB {
	env := os.Getenv("ENV")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"),
	)
	// os.Getenv関数を通じて環境変数に直接アクセスされ、その値はDSNの構築に使用されるが、関数の外部には保存されず、関数の実行が完了するとスコープ外になる

	var (
		db  *gorm.DB
		err error
	)

	if env == "prod" {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		log.Println("Setup postgresql database")
	} else {
		db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		log.Println("Setup sqlite database")
	}

	if err != nil {
		panic("Failed to connect database")
	}

	return db
}
