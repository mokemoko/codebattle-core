package main

import (
	"batch/commands"
	"database/sql"
	"flag"
	"log"
	"math/rand"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Args struct {
	ContestId string
	IsDebug   bool
	Command   string
}

type Result struct {
	Status int
}

func parseArgs() Args {
	args := Args{}
	flag.BoolVar(&args.IsDebug, "debug", false, "")
	flag.StringVar(&args.ContestId, "contestId", "", "")
	flag.StringVar(&args.Command, "command", "execute", "choose in [entry, matchmake, execute]")
	flag.Parse()
	return args
}

func setupDatabase(args Args) error {
	db, err := sql.Open("sqlite3", "../sql/db.sqlite3")
	if err != nil {
		return err
	}
	if args.IsDebug {
		boil.DebugMode = true
	}
	boil.SetDB(db)
	return nil
}

func init() {
	log.SetFlags(log.Llongfile)
	rand.Seed(time.Now().UnixNano())
}

func main() {
	args := parseArgs()

	err := setupDatabase(args)
	if err != nil {
		log.Fatal(err)
	}

	switch args.Command {
	case "entry":
		commands.RunEntry()
	case "matchmake":
		commands.RunMatchMake(args.ContestId)
	case "execute":
		// TODO: dequeue match
		commands.RunExecute()
	default:
		log.Fatalf("Invalid command: %s", args.Command)
	}
}
