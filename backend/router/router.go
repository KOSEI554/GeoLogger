package router

import (
	"net/http"

	"geologger/handler"
)

// HTTPルーターをセットアップする関数
// URLパスとそれに対応するハンドラー関数をマッピング
//
// 戻り値:
//   - 設定されたHTTPハンドラー
func SetupRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/locations", handler.HandleLocations)
	return mux
}
