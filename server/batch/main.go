package main

import (
	entry "batch/commands"
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os/exec"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mokemoko/codebattle-core/server/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"gonum.org/v1/gonum/stat/combin"
)

type Args struct {
	ContestId string
	IsDebug   bool
	Command   string
}

type MatchInfo struct {
	Contest      *models.Contest
	MatchEntries []*models.Match
	Entries      []*models.Entry
}

type ExecuteParams struct {
	Images []string    `json:"images"`
	Meta   interface{} `json:"meta,omitempty"`
}

type ExecuteResult struct {
	Status int `json:"status"`
	Rank   []int
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

func makeMatch(contestId string) (*MatchInfo, error) {
	matchInfo := MatchInfo{}

	contest, err := models.FindContestG(context.Background(), contestId)
	if err != nil {
		return nil, err
	}
	matchInfo.Contest = contest

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
			AfterScore:  entry.Score,
			CreatedAt:   ts,
			UpdatedAt:   ts,
		}
		err = matchEntry.Insert(context.Background(), tx, boil.Infer())
		if err != nil {
			_ = tx.Rollback()
			return nil, err
		}
		matchInfo.MatchEntries = append(matchInfo.MatchEntries, &matchEntry)
		matchInfo.Entries = append(matchInfo.Entries, entry)
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &matchInfo, nil
}

func executeMatch(matchInfo *MatchInfo) (ExecuteResult, error) {
	result := ExecuteResult{
		Status: -1,
	}
	// TODO: Contest.Meta を参照して設定
	params := ExecuteParams{}
	for _, entry := range matchInfo.Entries {
		// TODO: use hash
		params.Images = append(params.Images, entry.Repository)
	}

	b, err := json.Marshal(params)
	if err != nil {
		return result, err
	}

	// TODO: replace executor
	stdout, err := exec.Command("/usr/local/bin/docker", "run", "--rm", "-v", "/var/run/docker.sock:/var/run/docker.sock", "cbe-bomberman", string(b)).Output()
	if err != nil {
		return result, err
	}

	// TODO: put log to s3
	err = ioutil.WriteFile(fmt.Sprintf("logs/%s.log", matchInfo.MatchEntries[0].ID), stdout, 0644)
	if err != nil {
		return result, err
	}

	// parse match result from last line
	lines := strings.Split(string(stdout), "\n")
	err = json.Unmarshal([]byte(lines[len(lines)-2]), &result)
	if err != nil {
		return result, err
	}

	for idx, matchEntry := range matchInfo.MatchEntries {
		matchEntry.Rank = int64(result.Rank[idx])
	}
	return result, nil
}

func calcRateDiff(winScore int64, loseScore int64) int64 {
	// use elo rating
	prob := 1 / (math.Pow(10, float64(winScore-loseScore)/400) + 1)
	return int64(math.Round(prob * 32))
}

func rateMatch(matchEntries []*models.Match) {
	var validEntries []*models.Match
	for _, entry := range matchEntries {
		// 失格 / 無効判定のエントリーは除外
		if entry.Rank > 0 {
			validEntries = append(validEntries, entry)
		}
	}
	if len(validEntries) < 2 {
		return
	}
	for _, pair := range combin.Combinations(len(validEntries), 2) {
		entry1 := validEntries[pair[0]]
		entry2 := validEntries[pair[1]]
		if entry1.Rank < entry2.Rank {
			diff := calcRateDiff(entry1.BeforeScore, entry2.BeforeScore)
			entry1.AfterScore += diff
			entry2.AfterScore -= diff
		} else if entry2.Rank < entry1.Rank {
			diff := calcRateDiff(entry2.BeforeScore, entry1.BeforeScore)
			entry2.AfterScore += diff
			entry1.AfterScore -= diff
		} else {
			// draw
			diff := calcRateDiff(entry1.BeforeScore, entry2.BeforeScore) - 16
			entry1.AfterScore += diff
			entry2.AfterScore -= diff
		}
	}
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

	switch args.Command {
	case "entry":
		err := entry.Execute()
		if err != nil {
			log.Fatal(err)
		}
	case "matchmake":
		matchInfo, err := makeMatch(args.ContestId)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%+v", matchInfo)
	case "execute":
		// TODO: dequeue match
		matchInfo, err := makeMatch(args.ContestId)
		result, err := executeMatch(matchInfo)
		if err != nil {
			log.Fatal(err)
		}
		rateMatch(matchInfo.MatchEntries)
		err = saveMatch(matchInfo.MatchEntries)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%+v", result)
	default:
		log.Fatalf("Invalid command: %s", args.Command)
	}
}
