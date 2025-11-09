//$(cd "$(dirname "$0")"; pwd)/bootstrap.sh go run "$0" "$@"; exit
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	host := os.Getenv("WP_DB_HOST")
	port := os.Getenv("WP_DB_PORT")
	user := os.Getenv("WP_DB_USER")
	password := os.Getenv("WP_DB_PASSWORD")
	dbname := os.Getenv("WP_DB_NAME")
	
	if user == "" || password == "" || dbname == "" {
		log.Fatal("Missing database credentials. Please set WP_DB_* environment variables.")
	}
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT post_title, post_name, post_type, post_date
		FROM wp_posts
		WHERE post_status = 'publish' AND post_type IN ('post', 'page')
		ORDER BY post_date DESC
		LIMIT 5
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("\n=== URL Verification ===\n")

	for rows.Next() {
		var title, slug, postType string
		var postDate time.Time
		rows.Scan(&title, &slug, &postType, &postDate)
		
		fmt.Printf("Title: %s\n", title)
		
		wpURL := ""
		hugoURL := ""
		if postType == "page" {
			wpURL = fmt.Sprintf("https://dan.bodar.com/%s/", slug)
			hugoURL = fmt.Sprintf("https://dan.bodar.com/%s/", slug)
		} else {
			wpURL = fmt.Sprintf("https://dan.bodar.com/%d/%02d/%02d/%s/", 
				postDate.Year(), postDate.Month(), postDate.Day(), slug)
			hugoURL = fmt.Sprintf("https://dan.bodar.com/%d/%02d/%02d/%s/", 
				postDate.Year(), postDate.Month(), postDate.Day(), slug)
		}
		
		fmt.Printf("WordPress: %s\n", wpURL)
		fmt.Printf("Hugo:      %s\n", hugoURL)
		
		if wpURL == hugoURL {
			fmt.Println("✓ URLs MATCH!")
		} else {
			fmt.Println("✗ URLs DIFFER")
		}
		fmt.Println()
	}
}
