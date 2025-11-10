//$(cd "$(dirname "$0")"; pwd)/bootstrap.sh "$0" "$@"; exit
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	md "github.com/JohannesKaufmann/html-to-markdown/v2"
)

// Configuration
const (
	hugoContentDir = "content"
)

// Database configuration from environment variables
func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

var (
	dbHost      = getEnv("WP_DB_HOST", "localhost")
	dbPort      = getEnv("WP_DB_PORT", "3306")
	dbUser      = getEnv("WP_DB_USER", "")
	dbPassword  = getEnv("WP_DB_PASSWORD", "")
	dbName      = getEnv("WP_DB_NAME", "")
	wpBackupDir = getEnv("WP_BACKUP_DIR", "")
	wpSiteURL   = getEnv("WP_SITE_URL", "")
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

func copyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	os.MkdirAll(filepath.Dir(dst), 0755)
	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

func processContent(content string, postSlug string, postType string) string {
	// Remove all WordPress block comments (both opening and closing)
	wpCommentRe := regexp.MustCompile(`<!-- /?wp:[^>]+ -->`)
	content = wpCommentRe.ReplaceAllString(content, "")

	// Remove empty divs and paragraphs that contain only whitespace or &nbsp;
	content = regexp.MustCompile(`<div[^>]*>\s*(&nbsp;)?\s*</div>`).ReplaceAllString(content, "")
	content = regexp.MustCompile(`<p[^>]*>\s*(&nbsp;)?\s*</p>`).ReplaceAllString(content, "")

	// WordPress sometimes already has <br> or <br /> tags - normalize them to newlines first
	content = regexp.MustCompile(`<br\s*/?>`).ReplaceAllString(content, "\n")

	// Convert line breaks to <br> so the HTML-to-Markdown converter preserves them
	// But we need to do this BEFORE we create code blocks, so they don't get affected
	content = strings.ReplaceAll(content, "\n", "<br>")

	// Convert WordPress code shortcodes to HTML pre/code blocks
	// Match [language]...[/language] patterns for common languages
	// This needs to happen AFTER <br> conversion so code blocks are preserved properly
	languages := []string{"java", "javascript", "csharp", "python", "go", "bash", "sql", "xml", "html", "css"}
	for _, lang := range languages {
		pattern := fmt.Sprintf(`\[%s\]([\s\S]*?)\[/%s\]`, lang, lang)
		codeRe := regexp.MustCompile(pattern)
		content = codeRe.ReplaceAllStringFunc(content, func(match string) string {
			submatches := codeRe.FindStringSubmatch(match)
			if len(submatches) == 2 {
				code := submatches[1]
				// Remove the <br> tags that were added inside the code block
				code = strings.ReplaceAll(code, "<br>", "\n")
				// Wrap in pre/code tags with language class for markdown converter
				return fmt.Sprintf(`<pre><code class="language-%s">%s</code></pre>`, lang, code)
			}
			return match
		})
	}

	// Convert HTML to Markdown
	markdown, err := md.ConvertString(content)
	if err != nil {
		log.Printf("Warning: could not convert HTML to markdown: %v", err)
		markdown = content
	}

	// Fix malformed links where URL is in the link text and placeholder is in the href
	// Pattern: [https://example.com](_wp_link_placeholder) -> https://example.com
	linkFixRe := regexp.MustCompile(`\[(https?://[^\]]+)\]\([^)]*\)`)
	markdown = linkFixRe.ReplaceAllString(markdown, "$1")

	// Remove lines that only contain whitespace or non-breaking spaces (U+00A0)
	markdown = regexp.MustCompile(`(?m)^[\s\x{00A0}]+$`).ReplaceAllString(markdown, "")

	// Remove excessive blank lines (more than 2 consecutive newlines)
	markdown = regexp.MustCompile(`\n{3,}`).ReplaceAllString(markdown, "\n\n")

	// Process image/attachment URLs and copy files
	if wpSiteURL != "" && wpBackupDir != "" {
		// Match both http:// and https:// versions of the site URL
		siteURLPattern := strings.Replace(regexp.QuoteMeta(wpSiteURL), "https://", "https?://", 1)
		re := regexp.MustCompile(siteURLPattern + `/wp-content/uploads/([^"'\s<>)]+)`)
		markdown = re.ReplaceAllStringFunc(markdown, func(match string) string {
			// Extract the path after /uploads/
			parts := strings.SplitN(match, "/wp-content/uploads/", 2)
			if len(parts) != 2 {
				return match
			}

			uploadPath := parts[1]
			srcFile := filepath.Join(wpBackupDir, "html/wp-content/uploads", uploadPath)

			// Determine destination directory based on post type
			var postDir string
			if postType == "page" {
				postDir = filepath.Join(hugoContentDir, postSlug)
			} else {
				postDir = filepath.Join(hugoContentDir, "posts", postSlug)
			}

			// Copy file to post directory with original filename
			filename := filepath.Base(uploadPath)
			dstFile := filepath.Join(postDir, filename)

			// Copy the file
			if err := copyFile(srcFile, dstFile); err != nil {
				log.Printf("Warning: could not copy %s: %v", srcFile, err)
				return match
			}

			// Return relative path (just the filename since it's in the same directory)
			return filename
		})
	}

	return markdown
}

func getAllTaxonomies(db *sql.DB) (map[int64][]string, map[int64][]string, error) {
	query := `
	SELECT tr.object_id, t.name, tt.taxonomy
	FROM wp_term_relationships tr
	JOIN wp_term_taxonomy tt ON tr.term_taxonomy_id = tt.term_taxonomy_id
	JOIN wp_terms t ON tt.term_id = t.term_id
	WHERE tt.taxonomy IN ('category', 'post_tag')
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	categories := make(map[int64][]string)
	tags := make(map[int64][]string)

	for rows.Next() {
		var postID int64
		var name, taxonomy string
		if err := rows.Scan(&postID, &name, &taxonomy); err != nil {
			return nil, nil, err
		}

		if taxonomy == "category" {
			categories[postID] = append(categories[postID], name)
		} else if taxonomy == "post_tag" {
			tags[postID] = append(tags[postID], name)
		}
	}

	return categories, tags, nil
}

func getAllComments(db *sql.DB) (map[int64][]Comment, error) {
	query := `
	SELECT comment_post_ID, comment_author, comment_author_email, comment_author_url,
	       comment_date, comment_content, comment_parent
	FROM wp_comments
	WHERE comment_approved = '1'
	ORDER BY comment_post_ID, comment_date ASC
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := make(map[int64][]Comment)

	for rows.Next() {
		var postID int64
		var c Comment
		var commentDate time.Time
		if err := rows.Scan(&postID, &c.Author, &c.Email, &c.URL, &commentDate, &c.Content, &c.Parent); err != nil {
			return nil, err
		}
		c.Date = commentDate.Format(time.RFC3339)
		c.Content = cleanHTML(c.Content)
		comments[postID] = append(comments[postID], c)
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

	// Process content: convert to markdown, remove WP comments, copy images, rewrite URLs
	processedContent := processContent(post.Content, post.Slug, post.PostType)
	content.WriteString(processedContent)

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

	// Load all taxonomies and comments in bulk
	allCategories, allTags, err := getAllTaxonomies(db)
	if err != nil {
		return fmt.Errorf("error getting taxonomies: %w", err)
	}

	allComments, err := getAllComments(db)
	if err != nil {
		return fmt.Errorf("error getting comments: %w", err)
	}

	// Create content directory
	if err := os.MkdirAll(hugoContentDir, 0755); err != nil {
		return fmt.Errorf("error creating content directory: %w", err)
	}

	// Create posts section with _index.md
	postsDir := filepath.Join(hugoContentDir, "posts")
	if err := os.MkdirAll(postsDir, 0755); err != nil {
		return fmt.Errorf("error creating posts directory: %w", err)
	}
	postsIndexContent := `---
title: "Posts"
---
`
	if err := os.WriteFile(filepath.Join(postsDir, "_index.md"), []byte(postsIndexContent), 0644); err != nil {
		return fmt.Errorf("error creating posts _index.md: %w", err)
	}

	exportedCount := 0

	for _, post := range posts {
		// Lookup taxonomies and comments from maps
		categories := allCategories[post.ID]
		tags := allTags[post.ID]
		comments := allComments[post.ID]

		// Determine output path based on post type
		var outputPath string
		if post.PostType == "page" {
			// Pages go in their own directory with index.md
			outputDir := filepath.Join(hugoContentDir, post.Slug)
			if err := os.MkdirAll(outputDir, 0755); err != nil {
				return fmt.Errorf("error creating page directory %s: %w", outputDir, err)
			}
			outputPath = filepath.Join(outputDir, "index.md")
		} else {
			// Posts go in content/posts/slug/index.md (page bundle)
			outputDir := filepath.Join(hugoContentDir, "posts", post.Slug)
			if err := os.MkdirAll(outputDir, 0755); err != nil {
				return fmt.Errorf("error creating post directory %s: %w", outputDir, err)
			}
			outputPath = filepath.Join(outputDir, "index.md")
		}

		// Create Hugo content (images will be copied during content processing)
		hugoContent := createHugoPost(post, categories, tags, comments)

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
	fmt.Printf("Content directory: %s\n", hugoContentDir)

	return nil
}

func main() {
	if err := exportPosts(); err != nil {
		log.Fatalf("Error exporting posts: %v", err)
	}
}
