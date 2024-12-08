package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

var connStr = "user=wira_user password=ilham2002 dbname=wira_db sslmode=disable"

type Player struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	ClassID  int    `json:"classID"`
	Score    int    `json:"score"`
}

func getRankings(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	page := r.URL.Query().Get("page")

	if limit == "" {
		limit = "50" 
	if page == "" {
		page = "1" 
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		log.Fatal("Error converting limit to integer:", err)
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		log.Fatal("Error converting page to integer:", err)
	}

	offset := (pageInt - 1) * limitInt

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	defer db.Close()

	query := `SELECT a.username, s.reward_score 
			  FROM Account a 
			  INNER JOIN Scores s ON a.acc_id = s.char_id 
			  ORDER BY s.reward_score DESC 
			  LIMIT $1 OFFSET $2`
	rows, err := db.Query(query, limitInt, offset)
	if err != nil {
		log.Fatal("Query execution failed:", err)
	}
	defer rows.Close()

	var players []Player
	for rows.Next() {
		var player Player
		if err := rows.Scan(&player.Username, &player.Score); err != nil {
			log.Fatal("Failed to scan row:", err)
		}
		players = append(players, player)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(players)
}

func addPlayer(w http.ResponseWriter, r *http.Request) {
	var player Player
	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	defer db.Close()

	query := `INSERT INTO Account (username, email) VALUES ($1, $2) RETURNING acc_id`
	var accID int
	err = db.QueryRow(query, player.Username, player.Email).Scan(&accID)
	if err != nil {
		log.Fatal("Failed to insert into Account table:", err)
	}

	_, err = db.Exec(`INSERT INTO Character (acc_id, class_id) VALUES ($1, $2)`, accID, player.ClassID)
	if err != nil {
		log.Fatal("Failed to insert into Character table:", err)
	}

	_, err = db.Exec(`INSERT INTO Scores (char_id, reward_score) VALUES ($1, $2)`, accID, player.Score)
	if err != nil {
		log.Fatal("Failed to insert into Scores table:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(player)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/rankings", getRankings) 
	mux.HandleFunc("/api/add-player", addPlayer) 

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8081"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	handler := c.Handler(mux)

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
