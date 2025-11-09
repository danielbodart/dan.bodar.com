# dan.bodar.com

Personal technical blog powered by Hugo, migrated from WordPress.

## Migration Scripts

This repository includes Go scripts used for the WordPress to Hugo migration:

- **`wp_to_hugo.go`** - Exports WordPress posts and pages to Hugo markdown format
  - Connects to WordPress MySQL database
  - Exports posts with front matter (title, date, categories, tags)
  - Preserves comments as JSON in front matter
  - Creates `content/posts/` and `content/pages/` directories
  - Requires `WP_DB_*` environment variables

- **`query_wp.go`** - Queries WordPress database to compare URL structures
  - Verifies permalink structure compatibility
  - Shows sample posts and their WordPress vs Hugo URLs
  - Helps identify redirect requirements

- **`verify_urls.go`** - Verifies URL mapping between WordPress and Hugo
  - Checks date-based permalink structure
  - Confirms URL consistency for posts and pages

### Running Migration Scripts

```bash
# Set environment variables
export WP_DB_HOST=localhost
export WP_DB_PORT=3306
export WP_DB_USER=your_user
export WP_DB_PASSWORD=your_password
export WP_DB_NAME=your_database

# Execute scripts directly (they use shebang)
./wp_to_hugo.go
./query_wp.go
./verify_urls.go
```

## Hugo Setup

Standard Hugo configuration with:
- Theme: [hugo-blog-awesome](https://github.com/hugo-sid/hugo-blog-awesome)
- Permalinks: `/:year/:month/:day/:slug/` for posts
- Content in `content/posts/` and `content/pages/`
- Site: https://dan.bodar.com/

### Development

```bash
# Install dependencies
go mod download

# Run Hugo dev server
hugo server -D

# Build site
hugo
```

## Dependencies

- Go 1.x (for migration scripts)
- Hugo (for static site generation)
- MySQL client library: `github.com/go-sql-driver/mysql`
