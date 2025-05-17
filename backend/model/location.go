package model

// 位置情報を表現する構造体
// フロントエンドとバックエンド間でやり取りされるデータモデル
type Location struct {
	// 緯度
	Lat float64 `json:"lat"`
	// 経度
	Lng float64 `json:"lng"`
}
