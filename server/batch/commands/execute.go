package commands

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/mokemoko/codebattle-core/server/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"gonum.org/v1/gonum/stat/combin"
	"io/ioutil"
	"log"
	"math"
	"os/exec"
	"strings"
)

type ExecuteParams struct {
	Images []string    `json:"images"`
	Meta   interface{} `json:"meta,omitempty"`
}

type ExecuteResult struct {
	Status int `json:"status"`
	Rank   []int
}

func getMatchList(matchId string) ([]*models.Match, error) {
	return models.Matches(
		models.MatchWhere.ID.EQ(matchId),
		Load(models.MatchRels.Entry),
	).AllG(context.Background())
}

func executeMatch(matchList []*models.Match) (ExecuteResult, error) {
	result := ExecuteResult{
		Status: -1,
	}
	// TODO: Contest.Meta を参照して設定
	params := ExecuteParams{}
	for _, match := range matchList {
		params.Images = append(params.Images, genRepoHash(match.R.Entry))
	}

	b, err := json.Marshal(params)
	if err != nil {
		return result, err
	}

	// TODO: replace executor
	commandArgs := []string{"run", "--rm", "-v", "/var/run/docker.sock:/var/run/docker.sock", "cbe-bomberman", string(b)}
	log.Print(commandArgs)
	stdout, err := exec.Command("/usr/local/bin/docker", commandArgs...).Output()
	if err != nil {
		return result, err
	}

	// TODO: put log to s3
	err = ioutil.WriteFile(fmt.Sprintf("logs/%s.log", matchList[0].ID), stdout, 0644)
	if err != nil {
		return result, err
	}

	// parse match result from last line
	lines := strings.Split(string(stdout), "\n")
	err = json.Unmarshal([]byte(lines[len(lines)-2]), &result)
	if err != nil {
		return result, err
	}

	for idx, matchEntry := range matchList {
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
		// TODO: 正規化
		matchEntry.Status = 1
		_, err = matchEntry.Update(context.Background(), tx, boil.Infer())
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	return err
}

func RunExecute() {
	matchList, err := models.Matches(
		models.MatchWhere.Status.EQ(0),
		Distinct("id"),
		Select("id"),
	).AllG(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, match := range matchList {
		matchList, err := getMatchList(match.ID)
		if err != nil {
			log.Fatal(err)
		}
		result, err := executeMatch(matchList)
		if err != nil {
			log.Fatal(err)
		}
		rateMatch(matchList)
		err = saveMatch(matchList)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%+v", result)
	}
}
