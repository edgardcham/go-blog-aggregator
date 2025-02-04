package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("No username provided")
	} else if len(cmd.arguments) > 1 {
		return fmt.Errorf("Too many arguments provided")
	}
	username := cmd.arguments[0]
	if err := s.config.SetUser(username); err != nil {
		return err
	}
	fmt.Println("User set to", username)

	return nil
}
