package main

import (
	"fmt"
	"os"

	"github.com/edgardcham/go-blog-aggregator/internal/config"
)

type state struct {
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
	st := state{config: &gatorconfig}

	// Setup commands and register login.
	cmds := commands{cmds: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)

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
