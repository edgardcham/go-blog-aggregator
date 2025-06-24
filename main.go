package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/edgardcham/go-blog-aggregator/internal/config"
	"github.com/edgardcham/go-blog-aggregator/internal/database"
	_ "github.com/lib/pq" // postgres driver
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}

	// Open DB connection
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// test DB connection
	if err := db.Ping(); err != nil {
		panic(fmt.Errorf("failed to ping DB: %w", err))
	}

	// create the database queries instnace
	dbQueries := database.New(db)
	s := state{
		cfg: &cfg,
		db:  dbQueries,
	}

	cmds := commands{
		commands: make(map[string]func(*state, command) error),
	}

	// register commands
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed, ))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	// get args
	args := os.Args
	// I am running go run . login and this is not triggering the if statement
	if len(args) < 2 {
		fmt.Println("Usage: gator <command> [args...]")
		os.Exit(1)
	}

	cmdName := args[1]
	cmdArgs := args[2:]

	if err := cmds.run(&s, command{name: cmdName, args: cmdArgs}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
