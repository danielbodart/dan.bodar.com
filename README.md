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

All Go scripts use `bootstrap.sh` for automatic dependency installation (mise, Hugo, Go).

```bash
# Set environment variables
export WP_DB_HOST=localhost
export WP_DB_PORT=3306
export WP_DB_USER=your_user
export WP_DB_PASSWORD=your_password
export WP_DB_NAME=your_database

# Execute scripts directly - bootstrap.sh handles setup automatically
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

The `run` script (converted to Go) provides common tasks:

```bash
./run.go dev      # Start Hugo dev server
./run.go build    # Build site
./run.go version  # Show version info
./run.go clean    # Clean artifacts
./run.go ci       # CI build
./run.go tag      # Create git tag
```

Or use Hugo directly (bootstrap.sh ensures dependencies are installed):

```bash
./bootstrap.sh
hugo server -D
```

## Bootstrap System

All scripts use `bootstrap.sh` which automatically installs:
- mise (tool version manager)
- Go (via mise and `.tool-versions`)
- Hugo (via mise and `.tool-versions`)

Go scripts use the shebang `//$(cd "$(dirname "$0")"; pwd)/bootstrap.sh go run "$0" "$@"; exit` which:
1. Uses `//` prefix (Go comment) that shell collapses to `/`
2. Constructs absolute path to bootstrap.sh dynamically
3. bootstrap.sh installs mise/Go/Hugo, then executes `go run "$0" "$@"`
4. The `;exit` prevents shell from executing Go code as commands

## Dependencies

Managed automatically via mise:
- Go (specified in `.tool-versions`)
- Hugo (specified in `.tool-versions`)
- MySQL client library: `github.com/go-sql-driver/mysql` (via `go.mod`)
