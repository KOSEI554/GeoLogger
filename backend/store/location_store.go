package store

import (
	"sync"
	"geologger/model"
)

var (
	locations []model.Location
	mutex     sync.Mutex
)

// 新しい位置情報をメモリ内のスライスに追加する関数
//
// 引数:
//   - loc: 追加する位置情報
func AddLocation(loc model.Location) {
	mutex.Lock()
	defer mutex.Unlock()
	locations = append(locations, loc)
}

// メモリ内のスライスから全ての位置情報を取得する関数
//
// 戻り値:
//   - 保存されている全ての位置情報のコピー
func GetAllLocations() []model.Location {
	mutex.Lock()
	defer mutex.Unlock()
	return append([]model.Location{}, locations...) // シャローコピー
}
