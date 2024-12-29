package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/bootdotdev/gator/internal/config"
	"github.com/bootdotdev/gator/internal/database"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

type state struct {
	db    *database.Queries
	cfg   *config.Config
	rawDB *sql.DB // Add this line
}

func runMigrations(db *sql.DB) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set dialect: %w", err)
	}

	if err := goose.Up(db, "sql/schema"); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}
	return nil
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()

	// These lines run migrations
	if err := runMigrations(db); err != nil {
		log.Fatalf("error running migrations: %v", err)
	}

	dbQueries := database.New(db)

	programState := &state{
		db:    dbQueries,
		cfg:   &cfg,
		rawDB: db, // Add this line
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("register", handlerRegister)
	cmds.register("login", handlerLogin)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerListFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerListFeedFollows))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))
	cmds.register("fetch", handleFetch)
	fmt.Println("Registered commands:", cmds.registeredCommands)

	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [args...]")
		return
	}

	// Debugging print statements
	fmt.Println("Complete os.Args:", os.Args)

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	// More detailed trace
	fmt.Println("Extracted command:", cmdName)
	fmt.Println("Extracted arguments:", cmdArgs)

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
