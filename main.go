package main

import (
	"database/sql"
	"fmt"
	"github.com/edgardcham/go-blog-aggregator/internal/database"
	"os"

	"github.com/edgardcham/go-blog-aggregator/internal/config"
)

import _ "github.com/lib/pq"

type state struct {
	db     *database.Queries
	config *config.Config
}

type command struct {
	name      string
	arguments []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if f, ok := c.cmds[cmd.name]; ok {
		return f(s, cmd)
	}
	return fmt.Errorf("unknown command %s", cmd.name)
}

func main() {
	// Must have at least one argument (the command).
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments provided.")
		os.Exit(1)
	}

	// Load config and wrap it in state.
	gatorconfig := config.Read() // returns a config value

	// Connect to the database.
	dbUrl := gatorconfig.DB_URL
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		fmt.Println("Error connecting to database")
		os.Exit(1)
	}
	dbQueries := database.New(db)
	st := state{config: &gatorconfig, db: dbQueries}

	// Setup commands and register login.
	cmds := commands{cmds: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))

	// Build the command from CLI args.
	cmd := command{
		name:      os.Args[1],  // index 1 is the command
		arguments: os.Args[2:], // everything else is arguments
	}

	if err := cmds.run(&st, cmd); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
