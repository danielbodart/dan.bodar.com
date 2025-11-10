# dan.bodar.com

Personal technical blog powered by Hugo, migrated from WordPress.

## Migration Scripts

This repository includes Go scripts used for the WordPress to Hugo migration:

- **`wp_to_hugo.go`** - Exports WordPress posts and pages to Hugo markdown format
  - Connects to WordPress MySQL database
  - Exports posts with front matter (title, date, categories, tags)
  - Preserves comments as JSON in front matter
  - Copies images/attachments and rewrites URLs to Hugo paths
  - Creates `content/posts/`, `content/pages/`, `static/images/`, `static/downloads/` directories
  - Requires `WP_DB_*`, `WP_SITE_URL`, `WP_BACKUP_DIR` environment variables

- **`query_wp.go`** - Queries WordPress database to compare URL structures
  - Verifies permalink structure compatibility
  - Shows sample posts and their WordPress vs Hugo URLs
  - Helps identify redirect requirements

- **`verify_urls.go`** - Verifies URL mapping between WordPress and Hugo
  - Checks date-based permalink structure
  - Confirms URL consistency for posts and pages

### Running Migration Scripts

All Go scripts use `bootstrap.sh` for automatic dependency installation (mise, Hugo, Go).

```bash
# Set environment variables
export WP_DB_HOST=localhost
export WP_DB_PORT=3306
export WP_DB_USER=your_user
export WP_DB_PASSWORD=your_password
export WP_DB_NAME=your_database
export WP_SITE_URL=https://your-site.com
export WP_BACKUP_DIR=/path/to/wordpress/backup

# Execute scripts directly - bootstrap.sh handles setup automatically
./wp_to_hugo.go        # Migrates posts, copies images, rewrites URLs
./query_wp.go          # Query database for URL structure
./verify_urls.go       # Verify URL mappings
```

## Hugo Setup

Standard Hugo configuration with:
- Theme: [hugo-blog-awesome](https://github.com/hugo-sid/hugo-blog-awesome)
- Permalinks: `/:year/:month/:day/:slug/` for posts
- Content in `content/posts/` and `content/pages/`
- Site: https://dan.bodar.com/

### Development

The `run` script (converted to Go) provides common tasks:

```bash
./run dev      # Start Hugo dev server
./run build    # Build site
```

## Bootstrap System

All Go scripts are executable and automatically install their dependencies (mise, Go, Hugo) when run.

To create a new executable Go script, add this line at the top:
```go
//$(cd "$(dirname "$0")"; pwd)/bootstrap.sh "$0" "$@"; exit
package main
```

## Deployment

The site deploys automatically to GitHub Pages via GitHub Actions when pushing to `master`.

### GitHub Pages Setup

1. Go to repository Settings → Pages
2. Source: "GitHub Actions"
3. The workflow builds the site using `./run build` and deploys to Pages

### DNS Configuration

Add these DNS records for `dan.bodar.com`:
- `A` record pointing to GitHub Pages IPs:
  - `185.199.108.153`
  - `185.199.109.153`
  - `185.199.110.153`
  - `185.199.111.153`
- `CNAME` record for `www` pointing to `danielbodart.github.io`

Or use a single `CNAME` record for the apex domain if your DNS provider supports it:
- `CNAME` record for `@` pointing to `danielbodart.github.io`

## Dependencies

Managed automatically via mise:
- Go (specified in `.tool-versions`)
- Hugo (specified in `.tool-versions`)
- MySQL client library: `github.com/go-sql-driver/mysql` (via `go.mod`)
