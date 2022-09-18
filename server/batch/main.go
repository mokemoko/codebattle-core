package main

import (
	"context"
	"database/sql"
	"flag"
	"github.com/google/uuid"
	"log"
	"math/rand"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mokemoko/codebattle-core/server/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Args struct {
	ContestId string
	IsDebug   bool
}

type Result struct {
	Status int
}

func parseArgs() Args {
	args := Args{}
	flag.BoolVar(&args.IsDebug, "debug", false, "")
	flag.StringVar(&args.ContestId, "contestId", "", "")
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

func makeMatch(contestId string) ([]*models.Match, error) {
	entries, err := models.Entries(
		models.EntryWhere.ContestID.EQ(contestId),
		OrderBy("RANDOM()"),
		Limit(4),
	).AllG(context.Background())
	if err != nil {
		return nil, err
	}

	matchId := uuid.NewString()
	ts := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	var matchEntries []*models.Match

	tx, err := boil.GetDB().(*sql.DB).Begin()
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		matchEntry := models.Match{
			ID:          matchId,
			EntryID:     entry.ID,
			ContestID:   contestId,
			BeforeScore: entry.Score,
			CreatedAt:   ts,
			UpdatedAt:   ts,
		}
		err = matchEntry.Insert(context.Background(), tx, boil.Infer())
		if err != nil {
			_ = tx.Rollback()
			return nil, err
		}
		matchEntries = append(matchEntries, &matchEntry)
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return matchEntries, nil
}

func executeMatch(matchEntries []*models.Match) (Result, error) {
	// TODO: replace
	result := Result{
		Status: 1,
	}
	for _, matchEntry := range matchEntries {
		matchEntry.Rank = int64(rand.Intn(6) - 1)
	}
	return result, nil
}

func rateMatch(matchEntries []*models.Match) error {
	// TODO: implement
	for _, matchEntry := range matchEntries {
		matchEntry.AfterScore = matchEntry.BeforeScore + matchEntry.Rank*10
	}
	return nil
}

func saveMatch(matchEntries []*models.Match) error {
	tx, err := boil.GetDB().(*sql.DB).Begin()
	if err != nil {
		return err
	}

	for _, matchEntry := range matchEntries {
		_, err = matchEntry.Update(context.Background(), tx, boil.Infer())
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	return err
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

	matchEntries, err := makeMatch(args.ContestId)
	if err != nil {
		log.Fatal(err)
	}

	result, err := executeMatch(matchEntries)
	if err != nil {
		log.Fatal(err)
	}

	err = rateMatch(matchEntries)
	if err != nil {
		log.Fatal(err)
	}

	err = saveMatch(matchEntries)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", result)
}
