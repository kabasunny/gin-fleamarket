package infra

import (
	"log"

	"github.com/joho/godotenv" //.env ファイルから環境変数を読み込むために使用
)

func Initialize() {
	err := godotenv.Load() //プロジェクトルートディレクトリにある .env ファイルを探し、見つかった場合は環境変数としてアプリケーションにロード
	if err != nil {
		log.Fatal("Error loading .env file")// .env ファイルのロードに失敗した場合（例えば、ファイルが存在しない場合）、エラーメッセージをログに記録し、プログラムを終了
	}
}