package main

import (
	"log"
	"net/http"

	"geologger/middleware"
	"geologger/router"
)

// アプリケーションのエントリーポイント
// 1. ルーターをセットアップ
// 2. CORSミドルウェアを適用
// 3. 8080ポートでHTTPサーバーを起動
func main() {
	handler := middleware.WithCORS(router.SetupRoutes())

	log.Println("🚀 Backend running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
