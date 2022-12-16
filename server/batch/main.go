package main

import (
	"batch/commands"
	"flag"
	"github.com/mokemoko/codebattle-core/server/models"
	"log"
	"math/rand"
	"time"

	_ "github.com/mattn/go-sqlite3"
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

func init() {
	log.SetFlags(log.Llongfile)
	rand.Seed(time.Now().UnixNano())
}

func main() {
	args := parseArgs()

	if err := models.SetupDatabase(args.IsDebug); err != nil {
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
