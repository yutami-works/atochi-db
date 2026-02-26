// backend/main.go
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Location struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Histories []History `json:"histories"`
}

type History struct {
	ID           int     `json:"id"`
	LocationID   int     `json:"location_id"`
	Name         string  `json:"name"`
	FloorInfo    *string `json:"floor_info"`
	Note         *string `json:"note"`
	StartDate    *string `json:"start_date"` // 追加: 開始日
	EndDate      *string `json:"end_date"`   // 追加: 終了日
	ImageURL     *string `json:"image_url"`  // 追加: 画像URL
	DisplayOrder int     `json:"display_order"`
}

var db *sql.DB

func initDB() {
	var err error
	connStr := "user=atochi_user password=atochi_password dbname=atochi_db sslmode=disable host=localhost port=5432"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("DB接続設定エラー:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("DB疎通エラー:", err)
	}
	fmt.Println("データベースへの接続成功！")
}

func getLocationsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	rows, err := db.Query("SELECT id, name, address, latitude, longitude FROM locations")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var locations []Location
	for rows.Next() {
		var loc Location
		if err := rows.Scan(&loc.ID, &loc.Name, &loc.Address, &loc.Latitude, &loc.Longitude); err != nil {
			log.Println("Location scan error:", err)
			continue
		}

		// ▼ 取得するカラムに start_date, end_date, image_url を追加
		histRows, err := db.Query("SELECT id, location_id, name, floor_info, note, start_date, end_date, image_url, display_order FROM histories WHERE location_id = $1 ORDER BY display_order DESC", loc.ID)
		if err != nil {
			log.Println("History query error:", err)
			continue
		}
		defer histRows.Close()

		var histories []History
		for histRows.Next() {
			var hist History
			// ▼ Scanの割り当ても追加
			if err := histRows.Scan(&hist.ID, &hist.LocationID, &hist.Name, &hist.FloorInfo, &hist.Note, &hist.StartDate, &hist.EndDate, &hist.ImageURL, &hist.DisplayOrder); err != nil {
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
	initDB()
	http.HandleFunc("/api/locations", getLocationsHandler)
	fmt.Println("サーバー起動: http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("サーバーエラー:", err)
	}
}
