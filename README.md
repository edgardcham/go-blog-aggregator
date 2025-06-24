# Gator Blog Aggregator

Gator is a powerful CLI-based RSS feed aggregator that allows users to:

* Add RSS feeds from across the internet to be collected
* Store the collected posts in a PostgreSQL database
* Follow and unfollow RSS feeds that other users have added
* View summaries of the aggregated posts in the terminal, with a link to the full post
* Continuously aggregate RSS feeds with automated fetching

## Installation

### Prerequisites

* Go 1.19 or higher
* PostgreSQL database
* `goose` for database migrations
* `sqlc` for SQL code generation

### Install from Source

#### Option 1: Using go install (Recommended)

```bash
# Install gator globally
go install github.com/edgardcham/go-blog-aggregator@latest

# Now you can use gator from anywhere
gator --help
```

#### Option 2: Build from Source

```bash
# Clone the repository
git clone https://github.com/edgardcham/go-blog-aggregator.git
cd go-blog-aggregator

# Build the binary
go build -o gator .

# Option A: Use locally
./gator --help

# Option B: Move to PATH for global access
sudo mv gator /usr/local/bin/
gator --help
```

### Database Setup

1. Create a PostgreSQL database:
```bash
createdb gator_db
```

2. Set up your database connection in the config file:
```bash
# The app will look for a config file in ~/.gatorconfig.json
# Set your DATABASE_URL appropriately
export DATABASE_URL="postgres://username:password@localhost:5432/gator_db?sslmode=disable"
```

3. Run database migrations:
```bash
# Install goose if you haven't already
go install github.com/pressly/goose/v3/cmd/goose@latest

# Run migrations
goose postgres $DATABASE_URL up
```

## Quick Start

```bash
# Install gator
go install github.com/edgardcham/go-blog-aggregator@latest

# Set up database connection
export DATABASE_URL="postgres://localhost/gator_db?sslmode=disable"

# Create your first user
gator register myusername

# Add an RSS feed
gator addfeed "Example Feed" https://example.com/rss

# Start aggregating feeds
gator agg 30s
```

## Goals

* Integrate Go with a PostgreSQL DB
* Use SQL to migrate a database using `sqlc` and `goose`
* Write a long-running service that continuously fetches new posts from RSS feeds and stores them in the database

## Commands Reference

### User Management

#### `register <username>`
Creates a new user account in the system.
- **Arguments**: `username` - The desired username (must be unique)
- **Example**: `gator register alice`
- **Notes**: Automatically logs you in as the new user

#### `login <username>`
Switches to an existing user account.
- **Arguments**: `username` - The username to log in as
- **Example**: `gator login bob`
- **Notes**: User must exist in the system

#### `users`
Lists all registered users, highlighting the currently logged-in user.
- **Arguments**: None
- **Example**: `gator users`
- **Output**: Shows all usernames with "(current)" marker for active user

#### `reset`
Clears the entire database including all users, feeds, and posts.
- **Arguments**: None
- **Example**: `gator reset`
- **Warning**: This action is irreversible!

### Feed Management

#### `addfeed <name> <url>` (requires login)
Adds a new RSS feed to the system and automatically follows it.
- **Arguments**: 
  - `name` - A friendly name for the feed
  - `url` - The RSS feed URL
- **Example**: `gator addfeed "Tech News" https://example.com/rss`
- **Output**: Feed ID, name, URL, and user ID
- **Notes**: Feed URLs must be unique in the system

#### `feeds`
Lists all feeds in the system with their creators.
- **Arguments**: None
- **Example**: `gator feeds`
- **Output**: Feed name, URL, and creator username

#### `follow <url>` (requires login)
Follows an existing feed by URL.
- **Arguments**: `url` - The RSS feed URL to follow
- **Example**: `gator follow https://example.com/rss`
- **Notes**: Feed must already exist in the system

#### `following` (requires login)
Lists all feeds you're currently following.
- **Arguments**: None
- **Example**: `gator following`
- **Output**: Feed name and URL for each followed feed

#### `unfollow <url>` (requires login)
Unfollows a feed by URL.
- **Arguments**: `url` - The RSS feed URL to unfollow
- **Example**: `gator unfollow https://example.com/rss`

### RSS Aggregation

#### `agg <duration>`
Starts the RSS aggregation service that continuously fetches feeds.
- **Arguments**: `duration` - Time interval between fetches (e.g., "30s", "5m", "1h")
- **Example**: `gator agg 1m`
- **Notes**: 
  - Runs continuously until stopped
  - Fetches one feed at a time in round-robin fashion
  - Stores new posts in the database

#### `browse [limit]` (requires login)
Displays posts from feeds you follow.
- **Arguments**: `limit` (optional) - Number of posts to display (default: 2)
- **Example**: `gator browse 10`
- **Output**: Post details from your followed feeds

## Authentication

Commands marked with "(requires login)" use middleware to ensure a user is logged in. The system maintains the current user state in a local configuration file.

## Usage Examples

### Getting Started

```bash
# After installation, register a new user
gator register alice

# Add your first RSS feed
gator addfeed "Go Blog" https://go.dev/blog/feed.atom

# Start aggregating feeds every 60 seconds
gator agg 60s
```

### Multi-User Workflow

```bash
# Alice adds some feeds
gator login alice
gator addfeed "Tech News" https://news.ycombinator.com/rss
gator addfeed "Go Blog" https://go.dev/blog/feed.atom

# Bob registers and follows Alice's feeds
gator register bob
gator feeds  # See all available feeds
gator follow https://news.ycombinator.com/rss

# Bob views his followed feeds
gator following
gator browse 5  # Show 5 latest posts
```

## Technical Implementation

### Database Schema
* **users**: Stores user information with unique usernames
* **feeds**: Stores RSS feed URLs with unique constraint  
* **feed_follows**: Many-to-many relationship between users and feeds
* **posts**: Stores aggregated posts from RSS feeds

### Key Technologies
* **PostgreSQL**: Database with foreign key constraints and cascading deletes
* **sqlc**: Type-safe SQL code generation
* **goose**: Database migration management
* **context**: Used for request cancellation and timeouts
* **UUID**: Unique identifiers for all entities
* **XML parsing**: For RSS feed processing with HTML entity decoding

### Error Handling
* Duplicate user/feed detection with PostgreSQL error codes
* Proper error propagation with wrapped errors
* User-friendly error messages
* Graceful handling of network timeouts

## Database Setup

### Prerequisites
- PostgreSQL installed and running
- Database created for the application
- Connection string configured

### Configuration
The application uses a local configuration file to store:
- Current logged-in user
- Database connection URL

## Goose

To install goose:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### Running a migration

Goose works on a `.sql` file.

It works by specifying `-- +goose Up` for the up migration and `-- +goose Down` for the down migration.

To run it:

```bash
goose postgres <connection_string> up
```

## SQLC

SQLC is a Go program that generates Go code from SQL queries. It's not exactly an ORM, but rather a tool that makes working with raw SQL easy and type-safe.

To configure SQLC, create a `sqlc.yaml` file, for example:

```bash
version: "2"
sql:
  - schema: "sql/schema"
    queries: "sql/queries"
    engine: "postgresql"
    gen:
      go:
        out: "internal/database"
```

Here, we're telling SQLC to look in the `sql/schema` directory for our schema structure (same files that Goose uses **BUT** SQLC disregards "down" migrations), and in the `sql/queries` direcotry for queries. We're also telling it to generate Go code in the `internal/database` directory.

Here's an example query:

```sql
-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;
```

Here the `$1`, `$2`, `$3` and `$4` are parameter that we'll be able to pass into the query in our Go code. The `:one` at teh end tells SQLC that we expect to get back a single row.

To generate the code, we use:

```bash
sqlc generate
```
