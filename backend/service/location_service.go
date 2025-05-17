package service

import (
	"geologger/model"
	"geologger/store"
)

// 位置情報を保存するサービス関数
// ハンドラー層とデータストア層の間にあるビジネスロジック層
//
// 引数:
//   - loc: 保存する位置情報
func SaveLocation(loc model.Location) {
	store.AddLocation(loc)
}

// 全ての位置情報を取得するサービス関数
//
// 戻り値:
//   - 保存されている全ての位置情報
func FetchLocations() []model.Location {
	return store.GetAllLocations()
}
