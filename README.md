# Gator Blog Aggregator

Gator is a CLI tool that allows users to:

* Add RSS feeds from across the internet to be collected
* Store the collected posts in a PostgreSQL database
* Follow and unfollow RSS feeds that other users have added
* View summaries of the aggregated posts in the terminal, with a link to the full post

## Goals

* Integrate Go with a PostgreSQL DB
* Use SQL to migrate a database using `sqlc` and `goose`
* Write a long runnning-service that continuously fetches new posts from RSS feeds and stores them in the database

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
