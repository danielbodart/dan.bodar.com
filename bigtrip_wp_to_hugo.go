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
	hugoContentDir = "content/bigtrip"
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

	// Protect existing <pre> and <code> blocks from newline conversion
	// by temporarily replacing newlines inside them with a placeholder
	preRe := regexp.MustCompile(`(?s)<pre[^>]*>.*?</pre>`)
	codeRe := regexp.MustCompile(`(?s)<code[^>]*>.*?</code>`)
	protectedBlocks := make(map[string]string)
	placeholder := "___NEWLINE_PLACEHOLDER___"

	for _, re := range []*regexp.Regexp{preRe, codeRe} {
		content = re.ReplaceAllStringFunc(content, func(match string) string {
			protected := strings.ReplaceAll(match, "\n", placeholder)
			key := fmt.Sprintf("___PROTECTED_BLOCK_%d___", len(protectedBlocks))
			protectedBlocks[key] = protected
			return key
		})
	}

	// Convert line breaks to <br> so the HTML-to-Markdown converter preserves them
	content = strings.ReplaceAll(content, "\n", "<br>")

	// Restore protected blocks and their newlines
	for key, protected := range protectedBlocks {
		restored := strings.ReplaceAll(protected, placeholder, "\n")
		content = strings.ReplaceAll(content, key, restored)
	}

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

	// Normalize Big Trip image URLs in HTML BEFORE markdown conversion
	// Convert <img src="/bigtrip/pictures/file.jpg" /> to <img src="http://bodar.com/bigtrip/pictures/file.jpg" />
	// so they match our URL replacement regex later
	content = regexp.MustCompile(`<img\s+src="/bigtrip/`).ReplaceAllString(content, `<img src="http://bodar.com/bigtrip/`)

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

	// Convert external gallery links to local paths
	// Pattern: http://www.bodar.com/bigtrip/gallery/ -> /bigtrip/gallery/
	markdown = regexp.MustCompile(`https?://[^/]*bodar\.com/bigtrip/gallery`).ReplaceAllString(markdown, "/bigtrip/gallery")

	// Unescape underscores in URLs (markdown converter escapes them but they should be literal in URLs)
	// Match URLs and unescape underscores within them
	urlRe := regexp.MustCompile(`https?://[^\s\)]+`)
	markdown = urlRe.ReplaceAllStringFunc(markdown, func(url string) string {
		return strings.ReplaceAll(url, `\_`, `_`)
	})

	// Remove lines that only contain whitespace or non-breaking spaces (U+00A0)
	markdown = regexp.MustCompile(`(?m)^[\s\x{00A0}]+$`).ReplaceAllString(markdown, "")

	// Remove the "Â" character which appears from mishandled non-breaking spaces
	// This typically appears as standalone "Â" or "Â " in the text
	markdown = strings.ReplaceAll(markdown, "Â", "")

	// Remove excessive blank lines (more than 2 consecutive newlines)
	markdown = regexp.MustCompile(`\n{3,}`).ReplaceAllString(markdown, "\n\n")

	// Fix image markdown with leading slashes that should be relative (e.g. ![](/image.jpg) -> ![](image.jpg))
	// This needs to happen after markdown conversion but before our URL processing
	markdown = regexp.MustCompile(`!\[\]\(/([^/][^)]*)\)`).ReplaceAllString(markdown, "![]($1)")

	// Process image/attachment URLs and copy files
	if wpSiteURL != "" && wpBackupDir != "" {
		// Match image URLs - handle both /bigtrip/pictures/ and /wp-content/uploads/ patterns
		// For Big Trip blog: http://bodar.com/bigtrip/pictures/19.jpg
		// For standard WordPress: https://example.com/wp-content/uploads/2024/01/image.jpg

		postDir := filepath.Join(hugoContentDir, postSlug)

		// Pattern 1 & 2: /bigtrip/pictures/ (with or without domain)
		bigtripPicturesRe := regexp.MustCompile(`(?:https?://[^/]+)?/bigtrip/pictures/([^"'\s<>)]+)`)
		markdown = bigtripPicturesRe.ReplaceAllStringFunc(markdown, func(match string) string {
			// Extract filename after /pictures/
			parts := strings.SplitN(match, "/bigtrip/pictures/", 2)
			if len(parts) != 2 {
				return match
			}

			filename := parts[1]
			srcFile := filepath.Join(wpBackupDir, "pictures", filename)
			dstFile := filepath.Join(postDir, filename)

			// Copy the file
			if err := copyFile(srcFile, dstFile); err != nil {
				log.Printf("Warning: could not copy %s: %v", srcFile, err)
				return match
			}

			// Return relative path (just the filename since it's in the same directory)
			return filename
		})

		// Pattern 3 & 4: /bigtrip/gallery/dan/slides/ or /bigtrip/gallery/neil/slides/
		// Handle both single and double slashes before filename (e.g., /slides/file.jpg or /slides//file.jpg)
		bigtripGalleryRe := regexp.MustCompile(`https?://[^/]+/bigtrip/gallery/(dan|neil)/slides/+([^"'\s<>)]+)`)
		markdown = bigtripGalleryRe.ReplaceAllStringFunc(markdown, func(match string) string {
			// Extract person (dan/neil) and filename, handling multiple slashes
			re := regexp.MustCompile(`/bigtrip/gallery/(dan|neil)/slides/+([^"'\s<>)]+)`)
			matches := re.FindStringSubmatch(match)
			if len(matches) != 3 {
				return match
			}

			person := matches[1]
			filename := matches[2]
			// Clean filename of any leading slashes
			filename = strings.TrimLeft(filename, "/")

			srcFile := filepath.Join(wpBackupDir, "gallery", person, "slides", filename)
			dstFile := filepath.Join(postDir, filename)

			// Copy the file
			if err := copyFile(srcFile, dstFile); err != nil {
				log.Printf("Warning: could not copy %s: %v", srcFile, err)
				return match
			}

			// Return relative path (just the filename since it's in the same directory)
			return filename
		})

		// Also handle standard WordPress wp-content/uploads pattern
		siteURLPattern := strings.Replace(regexp.QuoteMeta(wpSiteURL), "https://", "https?://", 1)
		wpUploadsRe := regexp.MustCompile(siteURLPattern + `/wp-content/uploads/([^"'\s<>)]+)`)
		markdown = wpUploadsRe.ReplaceAllStringFunc(markdown, func(match string) string {
			// Extract the path after /uploads/
			parts := strings.SplitN(match, "/wp-content/uploads/", 2)
			if len(parts) != 2 {
				return match
			}

			uploadPath := parts[1]
			srcFile := filepath.Join(wpBackupDir, "html/wp-content/uploads", uploadPath)

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
	// Try modern WordPress schema first
	query := `
	SELECT tr.object_id, t.name, tt.taxonomy
	FROM wp_term_relationships tr
	JOIN wp_term_taxonomy tt ON tr.term_taxonomy_id = tt.term_taxonomy_id
	JOIN wp_terms t ON tt.term_id = t.term_id
	WHERE tt.taxonomy IN ('category', 'post_tag')
	`

	rows, err := db.Query(query)
	if err != nil {
		// If modern tables don't exist, try old WordPress 2.x schema
		oldQuery := `
		SELECT p2c.post_id, c.cat_name
		FROM wp_post2cat p2c
		JOIN wp_categories c ON p2c.category_id = c.cat_ID
		`
		rows, err = db.Query(oldQuery)
		if err != nil {
			return nil, nil, err
		}
		defer rows.Close()

		categories := make(map[int64][]string)
		tags := make(map[int64][]string)

		// Old WordPress only had categories, no tags
		for rows.Next() {
			var postID int64
			var name string
			if err := rows.Scan(&postID, &name); err != nil {
				return nil, nil, err
			}
			categories[postID] = append(categories[postID], name)
		}

		return categories, tags, nil
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
	// Note: Older WordPress versions may have empty post_type, treat as 'post'
	// Exclude "Finally setup my technical blog" as it's redundant with the main blog
	query := `
	SELECT ID, post_title, post_name, post_date, post_content,
	       post_excerpt, COALESCE(NULLIF(post_type, ''), 'post') as post_type
	FROM wp_posts
	WHERE post_status = 'publish' AND (post_type IN ('post', 'page') OR post_type = '')
	  AND post_title != 'Finally setup my technical blog'
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

	// Create section _index.md
	// Note: Title must be "Bigtrip" (not "Big Trip") to match the theme's menu logic
	// which compares lowercased section title with URL (bigtrip)
	indexContent := `---
title: "Bigtrip"
---
`
	if err := os.WriteFile(filepath.Join(hugoContentDir, "_index.md"), []byte(indexContent), 0644); err != nil {
		return fmt.Errorf("error creating _index.md: %w", err)
	}

	exportedCount := 0

	for _, post := range posts {
		// Lookup taxonomies and comments from maps
		categories := allCategories[post.ID]
		tags := allTags[post.ID]
		comments := allComments[post.ID]

		// All posts and pages go in content/bigtrip/slug/index.md (page bundle)
		outputDir := filepath.Join(hugoContentDir, post.Slug)
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return fmt.Errorf("error creating post directory %s: %w", outputDir, err)
		}
		outputPath := filepath.Join(outputDir, "index.md")

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
