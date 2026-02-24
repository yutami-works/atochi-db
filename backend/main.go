// backend/main.go
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq" // PostgreSQLドライバを読み込む
)

// データベースのテーブル構造に合わせた型定義
type Location struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Histories []History `json:"histories"` // 紐づく履歴の配列
}

type History struct {
	ID           int     `json:"id"`
	LocationID   int     `json:"location_id"`
	Name         string  `json:"name"`
	FloorInfo    *string `json:"floor_info"` // NULLが入り得るカラムはポインタ型（*string）にする
	Note         *string `json:"note"`
	DisplayOrder int     `json:"display_order"`
}

var db *sql.DB // グローバル変数としてDBコネクションを保持

// DB接続の初期化処理
func initDB() {
	var err error
	// docker-compose.yml で設定したDBの接続情報
	connStr := "user=atochi_user password=atochi_password dbname=atochi_db sslmode=disable host=localhost port=5432"

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("DB接続設定エラー:", err)
	}

	// 実際にDBに接続できるかテスト
	if err = db.Ping(); err != nil {
		log.Fatal("DB疎通エラー（DockerのDBは立ち上がっていますか？）:", err)
	}
	fmt.Println("データベースへの接続成功！")
}

func getLocationsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// 1. locations テーブルからすべての場所を取得
	rows, err := db.Query("SELECT id, name, address, latitude, longitude FROM locations")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var locations []Location
	for rows.Next() {
		var loc Location
		// 取得した行データを構造体にマッピング
		if err := rows.Scan(&loc.ID, &loc.Name, &loc.Address, &loc.Latitude, &loc.Longitude); err != nil {
			log.Println("Location scan error:", err)
			continue
		}

		// 2. この場所に紐づく履歴を histories テーブルから取得 (表示順の降順)
		histRows, err := db.Query("SELECT id, location_id, name, floor_info, note, display_order FROM histories WHERE location_id = $1 ORDER BY display_order DESC", loc.ID)
		if err != nil {
			log.Println("History query error:", err)
			continue
		}
		defer histRows.Close()

		var histories []History
		for histRows.Next() {
			var hist History
			if err := histRows.Scan(&hist.ID, &hist.LocationID, &hist.Name, &hist.FloorInfo, &hist.Note, &hist.DisplayOrder); err != nil {
				log.Println("History scan error:", err)
				continue
			}
			histories = append(histories, hist)
		}
		loc.Histories = histories
		locations = append(locations, loc)
	}

	json.NewEncoder(w).Encode(locations)
}

func main() {
	initDB() // 起動時に一度だけDB接続

	http.HandleFunc("/api/locations", getLocationsHandler)

	fmt.Println("サーバー起動: http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("サーバーエラー:", err)
	}
}
