# gator

gator is a command-line blog aggregator that fetches RSS feeds, stores posts in a PostgreSQL database, and lets you browse the posts from feeds you follow—all from your terminal.

## Prerequisites

- **Go**: Install [Go](https://golang.org/dl/) (version 1.23 or later is recommended).
- **PostgreSQL**: Install [PostgreSQL](https://www.postgresql.org/download/) and create a database for gator.

## Installation

### Using go install

To install gator as a production-ready CLI, run:

```bash
go install github.com/edgardcham/go-blog-aggregator@latest
```

This will compile the application and install the gator binary into your $GOPATH/bin. Make sure that directory is in your system’s PATH.

For development purposes, you can also run:


```bash
go run .
```
Configuration

gator relies on a configuration file to set up the database connection and other runtime settings.
	1.	Create a configuration file named .gatorconfig.json in your root directory.
Example .gatorconfig.json file:

```json
{
    "db_url": "postgres://username:password@localhost:5432/your_database?sslmode=disable",
    "current_user_name": ""
}
```

After installation, you can run gator with various commands. Below are a few of the commands you can use:
•	register: Register a new user.
•	login: Log in as an existing user.
•	addfeed: Add an RSS feed.
•	feeds: List all available feeds.
•	follow: Follow a feed.
•	unfollow: Unfollow a feed.
•	following: Show the feeds you are following.
•	agg: Continuously fetch posts from RSS feeds (aggregation).
Example (fetch feeds every 1 minute):

```bash
gator agg 1m
```

	•	browse: View posts from feeds you follow. This command accepts an optional limit parameter (defaults to 2).
Example (browse latest 5 posts):

```bash
gator browse 5
```

Running gator
•	Development Mode:
Run using the Go toolchain:

```bash
go run .
```

	•	Production Mode:
After installation via go install, simply run:

```bash
gator <command> [arguments...]
```
