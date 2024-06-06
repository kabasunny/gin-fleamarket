package main

import (
	"bytes"
	"encoding/json"
	"gin-fleamarket/dto"
	"gin-fleamarket/infra"
	"gin-fleamarket/models"
	"gin-fleamarket/services"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// TestMainは、テストの前にセットアップを行い、テストが完了した後にクリーンアップを行うための特別な関数
func TestMain(m *testing.M) {
	// testing.M型の値は、テストプロセス全体を表す
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalln("Error loading .env.test file")
	}

	code := m.Run() // テスト実行。mはテストランナーの制御を行うオブジェクト

	os.Exit(code) // 終了。今回後処理は特にない
}

// テスト用データの作成
func setupTestData(db *gorm.DB) {
	items := []models.Item{
		{Name: "テストアイテム1", Price: 1000, Description: "", SoldOut: false, UserID: 1},
		{Name: "テストアイテム2", Price: 2000, Description: "テスト2", SoldOut: true, UserID: 1},
		{Name: "テストアイテム3", Price: 3000, Description: "テスト3", SoldOut: false, UserID: 2},
	}

	users := []models.User{
		{Email: "test1@example.com", Password: "test1pass"},
		{Email: "test2@example.com", Password: "test2pass"},
	}

	for _, user := range users {
		db.Create(&user)
	}
	for _, item := range items {
		db.Create(&item)
	}
}

// 各テストで呼び出すデータベース初期化関数
func setup() *gin.Engine {
	db := infra.SetupDB()
	db.AutoMigrate(&models.Item{}, &models.User{})

	setupTestData(db)
	router := setupRouter(db)

	return router
}

func TestFindAll(t *testing.T) {
	// testing.Tはテストの制御や出力の管理を行う

	// テストのセットアップ
	router := setup()

	w := httptest.NewRecorder()                     // HTTPレスポンスを記録するために使用
	req, _ := http.NewRequest("GET", "/items", nil) // エンドポイントに対するリクエストを作成

	// APIリクエストの実行
	router.ServeHTTP(w, req) // リクエストをルーターに送信し、レスポンスをレコーダーに記録

	// APIの実行結果を取得
	var res map[string][]models.Item              // レスポンスデータを格納するためのマップ
	json.Unmarshal([]byte(w.Body.String()), &res) // レスポンスボディからJSONデータを解析し、マップに格納

	// アサーション
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 3, len(res["data"]))
	// 第2引数は、期待する値。第3引数は実際の値。
}

// 認証が必要のためトークンが必要
// 作成情報をリクエストボディにセットする
func TestCreate(t *testing.T) {
	// テストのセットアップ
	router := setup()

	// トークンの作成
	token, err := services.CreateToken(1, "test1@example.com")
	assert.Equal(t, nil, err)

	// リクエストボディの情報
	createItemInuput := dto.CreateItemInuput{
		Name:        "テストアイテム4",
		Price:       4000,
		Description: "Createテスト",
	}
	reqBody, _ := json.Marshal(createItemInuput)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(reqBody))
	req.Header.Set("Authorization", "Bearer "+*token)

	// APIリクエストの実行
	router.ServeHTTP(w, req)

	// APIの実行結果を取得
	var res map[string]models.Item
	json.Unmarshal([]byte(w.Body.String()), &res)

	// アサーション
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, uint(4), res["data"].ID)
}
