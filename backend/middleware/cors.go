package middleware

import "net/http"

// CORSヘッダーを追加するミドルウェア
// フロントエンドからのクロスオリジンリクエストを許可するために必要
// 
// 引数:
//   - next: 次に実行するハンドラー
//
// 戻り値:
//   - CORSヘッダーを追加するハンドラー
func WithCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
