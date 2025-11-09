//usr/bin/env go run "$0" "$@"; exit
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
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

	// Get permalink structure from options
	var permalinkStructure string
	err = db.QueryRow("SELECT option_value FROM wp_options WHERE option_name = 'permalink_structure'").Scan(&permalinkStructure)
	if err != nil {
		log.Printf("Could not get permalink structure: %v\n", err)
		permalinkStructure = "/%postname%/"
	}
	
	fmt.Printf("\n=== WordPress Permalink Structure: %s ===\n\n", permalinkStructure)

	rows, err := db.Query(`
		SELECT post_title, post_name, post_type
		FROM wp_posts
		WHERE post_status = 'publish' AND post_type IN ('post', 'page')
		ORDER BY post_date DESC
		LIMIT 5
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var title, slug, postType string
		rows.Scan(&title, &slug, &postType)
		
		fmt.Printf("Title: %s\n", title)
		fmt.Printf("Type:  %s\n", postType)
		
		wpURL := fmt.Sprintf("https://dan.bodar.com/%s/", slug)
		hugoURL := ""
		if postType == "page" {
			hugoURL = fmt.Sprintf("https://dan.bodar.com/%s/", slug)
		} else {
			hugoURL = fmt.Sprintf("https://dan.bodar.com/posts/%s/", slug)
		}
		
		fmt.Printf("WP URL:   %s\n", wpURL)
		fmt.Printf("Hugo URL: %s\n", hugoURL)
		
		if wpURL == hugoURL {
			fmt.Println("✓ URLs MATCH")
		} else {
			fmt.Println("✗ URLs DIFFER - Will need redirects")
		}
		fmt.Println()
	}
}
