package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// データの構造を定義
type Location struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Address string   `json:"address"`
	Tenants []Tenant `json:"tenants"`
}

type Tenant struct {
	Order int    `json:"order"`
	Name  string `json:"name"`
	Note  string `json:"note"`
}

// リクエストが来た時の処理
func getLocationsHandler(w http.ResponseWriter, r *http.Request) {
	// 別のポート（Next.js）からのアクセスを許可する設定
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 返すデータがJSON形式であることを伝える設定
	w.Header().Set("Content-Type", "application/json")

	// 返す仮のデータ（配列）
	mockData := []Location{
		{
			ID:      1,
			Name:    "千葉鑑定団 八千代店",
			Address: "千葉県八千代市勝田台南１丁目１８−１",
			Tenants: []Tenant{
				{Order: 1, Name: "ミドリ電化 八千代店", Note: "エディオンへ移管"},
				{Order: 2, Name: "エディオン 八千代店", Note: "ミドリ電化から店名変更"},
				{Order: 3, Name: "インターネットカフェ", Note: "SV過去断面で確認"},
				{Order: 4, Name: "千葉鑑定団 八千代店", Note: "現在の店舗"},
			},
		},
	}

	// データをJSONに変換してレスポンスとして返す
	json.NewEncoder(w).Encode(mockData)
}

func main() {
	// "/api/locations" というURLにアクセスが来たら getLocationsHandler を動かす
	http.HandleFunc("/api/locations", getLocationsHandler)

	fmt.Println("サーバー起動: http://localhost:8080")
	// 8080ポートでサーバーを起動
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("サーバーエラー:", err)
	}
}