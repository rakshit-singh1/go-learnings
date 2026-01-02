package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	// IMPORTANT
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	longURL := r.FormValue("url")
	if longURL == "" {
		http.Error(w, "URL required", http.StatusBadRequest)
		return
	}

	res, err := db.Exec("INSERT INTO urls(long_url) VALUES(?)", longURL)
	if err != nil {
		http.Error(w, "DB error", 500)
		return
	}

	id, _ := res.LastInsertId()
	shortCode := encodeBase62(id)

	_, err = db.Exec(
		"UPDATE urls SET short_code=? WHERE id=?",
		shortCode, id,
	)
	if err != nil {
		http.Error(w, "DB error", 500)
		return
	}

	shortURL := fmt.Sprintf("http://localhost:5050/%s", shortCode)
	w.Write([]byte(shortURL))
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/")
	if code == "" {
		w.Write([]byte("URL Shortener Running"))
		return
	}

	var longURL string
	err := db.QueryRow(
		"SELECT long_url FROM urls WHERE short_code=?",
		code,
	).Scan(&longURL)

	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

func main() {
	initDB()

	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/", redirectHandler)

	fmt.Println("Server running on :5050")
	log.Fatal(http.ListenAndServe(":5050", nil))
}
