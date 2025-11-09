//$(cd "$(dirname "$0")"; pwd)/bootstrap.sh go run "$0" "$@"; exit
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Configuration
const (
	hugoContentDir = "/home/dan/Projects/dan.bodar.com/content"
	hugoStaticDir  = "/home/dan/Projects/dan.bodar.com/static"
)

// Database configuration from environment variables
func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

var (
	dbHost     = getEnv("WP_DB_HOST", "localhost")
	dbPort     = getEnv("WP_DB_PORT", "3306")
	dbUser     = getEnv("WP_DB_USER", "")
	dbPassword = getEnv("WP_DB_PASSWORD", "")
	dbName     = getEnv("WP_DB_NAME", "")
)

// Comment represents a WordPress comment
type Comment struct {
	Author  string `json:"author"`
	Email   string `json:"email"`
	URL     string `json:"url"`
	Date    string `json:"date"`
	Content string `json:"content"`
	Parent  int64  `json:"parent"`
}

// PostData represents a WordPress post
type PostData struct {
	ID       int64
	Title    string
	Slug     string
	Date     time.Time
	Content  string
	Excerpt  string
	PostType string
}

// FrontMatter represents Hugo front matter
type FrontMatter struct {
	Title      string            `json:"title"`
	Date       string            `json:"date"`
	Slug       string            `json:"slug"`
	Type       string            `json:"type"`
	Categories []string          `json:"categories,omitempty"`
	Tags       []string          `json:"tags,omitempty"`
	Summary    string            `json:"summary,omitempty"`
	Comments   []Comment         `json:"comments,omitempty"`
	Extra      map[string]string `json:"-"`
}

func cleanHTML(text string) string {
	if text == "" {
		return ""
	}
	return html.UnescapeString(text)
}

func getPostTaxonomies(db *sql.DB, postID int64) ([]string, []string, error) {
	query := `
	SELECT t.name, tt.taxonomy
	FROM wp_term_relationships tr
	JOIN wp_term_taxonomy tt ON tr.term_taxonomy_id = tt.term_taxonomy_id
	JOIN wp_terms t ON tt.term_id = t.term_id
	WHERE tr.object_id = ? AND tt.taxonomy IN ('category', 'post_tag')
	`

	rows, err := db.Query(query, postID)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var categories, tags []string
	for rows.Next() {
		var name, taxonomy string
		if err := rows.Scan(&name, &taxonomy); err != nil {
			return nil, nil, err
		}

		if taxonomy == "category" {
			categories = append(categories, name)
		} else if taxonomy == "post_tag" {
			tags = append(tags, name)
		}
	}

	return categories, tags, nil
}

func getPostComments(db *sql.DB, postID int64) ([]Comment, error) {
	query := `
	SELECT comment_author, comment_author_email, comment_author_url,
	       comment_date, comment_content, comment_parent
	FROM wp_comments
	WHERE comment_post_ID = ? AND comment_approved = '1'
	ORDER BY comment_date ASC
	`

	rows, err := db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var c Comment
		var commentDate time.Time
		if err := rows.Scan(&c.Author, &c.Email, &c.URL, &commentDate, &c.Content, &c.Parent); err != nil {
			return nil, err
		}
		c.Date = commentDate.Format(time.RFC3339)
		c.Content = cleanHTML(c.Content)
		comments = append(comments, c)
	}

	return comments, nil
}

func createHugoPost(post PostData, categories, tags []string, comments []Comment) string {
	fm := FrontMatter{
		Title:      post.Title,
		Date:       post.Date.Format(time.RFC3339),
		Slug:       post.Slug,
		Categories: categories,
		Tags:       tags,
		Comments:   comments,
	}

	if post.Excerpt != "" {
		fm.Summary = cleanHTML(post.Excerpt)
	}

	var content strings.Builder
	content.WriteString("---\n")

	// Write front matter as YAML
	content.WriteString(fmt.Sprintf("title: %q\n", fm.Title))
	content.WriteString(fmt.Sprintf("date: %s\n", fm.Date))
	content.WriteString(fmt.Sprintf("slug: %q\n", fm.Slug))

	if len(categories) > 0 {
		content.WriteString("categories:\n")
		for _, cat := range categories {
			content.WriteString(fmt.Sprintf("  - %s\n", cat))
		}
	}

	if len(tags) > 0 {
		content.WriteString("tags:\n")
		for _, tag := range tags {
			content.WriteString(fmt.Sprintf("  - %s\n", tag))
		}
	}

	if fm.Summary != "" {
		if strings.Contains(fm.Summary, "\n") || strings.Contains(fm.Summary, "\"") {
			content.WriteString("summary: |\n")
			for _, line := range strings.Split(fm.Summary, "\n") {
				content.WriteString(fmt.Sprintf("  %s\n", line))
			}
		} else {
			content.WriteString(fmt.Sprintf("summary: %q\n", fm.Summary))
		}
	}

	if len(comments) > 0 {
		content.WriteString("comments:\n")
		for _, comment := range comments {
			commentJSON, _ := json.Marshal(comment)
			content.WriteString(fmt.Sprintf("  - %s\n", commentJSON))
		}
	}

	content.WriteString("---\n\n")
	content.WriteString(cleanHTML(post.Content))

	return content.String()
}

func exportPosts() error {
	// Build connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Connect to database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}
	defer db.Close()

	// Test connection
	if err := db.Ping(); err != nil {
		return fmt.Errorf("error pinging database: %w", err)
	}

	// Get all published posts and pages
	query := `
	SELECT ID, post_title, post_name, post_date, post_content,
	       post_excerpt, post_type
	FROM wp_posts
	WHERE post_status = 'publish' AND post_type IN ('post', 'page')
	ORDER BY post_date
	`

	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("error querying posts: %w", err)
	}
	defer rows.Close()

	var posts []PostData
	for rows.Next() {
		var post PostData
		if err := rows.Scan(&post.ID, &post.Title, &post.Slug, &post.Date,
			&post.Content, &post.Excerpt, &post.PostType); err != nil {
			return fmt.Errorf("error scanning post: %w", err)
		}
		if post.Slug == "" {
			post.Slug = fmt.Sprintf("%d", post.ID)
		}
		posts = append(posts, post)
	}

	fmt.Printf("Found %d posts to export\n", len(posts))

	// Create content directory structure
	postsDir := filepath.Join(hugoContentDir, "posts")
	pagesDir := filepath.Join(hugoContentDir, "pages")

	if err := os.MkdirAll(postsDir, 0755); err != nil {
		return fmt.Errorf("error creating posts directory: %w", err)
	}
	if err := os.MkdirAll(pagesDir, 0755); err != nil {
		return fmt.Errorf("error creating pages directory: %w", err)
	}

	exportedCount := 0

	for _, post := range posts {
		// Get taxonomies
		categories, tags, err := getPostTaxonomies(db, post.ID)
		if err != nil {
			log.Printf("Warning: error getting taxonomies for post %d: %v", post.ID, err)
		}

		// Get comments
		comments, err := getPostComments(db, post.ID)
		if err != nil {
			log.Printf("Warning: error getting comments for post %d: %v", post.ID, err)
		}

		// Create Hugo content
		hugoContent := createHugoPost(post, categories, tags, comments)

		// Determine output directory and filename
		var outputDir string
		if post.PostType == "page" {
			outputDir = pagesDir
		} else {
			outputDir = postsDir
		}

		filename := fmt.Sprintf("%s.md", post.Slug)
		outputPath := filepath.Join(outputDir, filename)

		// Write file
		if err := os.WriteFile(outputPath, []byte(hugoContent), 0644); err != nil {
			return fmt.Errorf("error writing file %s: %w", outputPath, err)
		}

		exportedCount++
		fmt.Printf("Exported: %s (%s) -> %s\n", post.Title, post.PostType, outputPath)
		if len(comments) > 0 {
			fmt.Printf("  - Included %d comments\n", len(comments))
		}
	}

	fmt.Printf("\nExport complete! Exported %d posts/pages\n", exportedCount)
	fmt.Printf("Posts directory: %s\n", postsDir)
	fmt.Printf("Pages directory: %s\n", pagesDir)

	return nil
}

func main() {
	if err := exportPosts(); err != nil {
		log.Fatalf("Error exporting posts: %v", err)
	}
}
