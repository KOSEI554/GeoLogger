package handler

import (
	"encoding/json"
	"net/http"

	"geologger/model"
	"geologger/service"
)

// 位置情報のリクエストを処理するハンドラー関数
// /api/locationsエンドポイントへのすべてのリクエストを処理
//
// 引数:
//   - w: HTTPレスポンスライター
//   - r: HTTPリクエスト
func HandleLocations(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var loc model.Location
		if err := json.NewDecoder(r.Body).Decode(&loc); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		service.SaveLocation(loc)
		w.WriteHeader(http.StatusCreated)

	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(service.FetchLocations())

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
