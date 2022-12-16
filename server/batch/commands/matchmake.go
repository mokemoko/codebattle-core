package commands

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/mokemoko/codebattle-core/server/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"log"
)

type MatchInfo struct {
	Contest      *models.Contest
	MatchEntries []*models.Match
	Entries      []*models.Entry
}

func makeMatch(contest *models.Contest) (*MatchInfo, error) {
	matchInfo := MatchInfo{
		Contest: contest,
	}

	entries, err := models.Entries(
		models.EntryWhere.ContestID.EQ(contest.ID),
		models.EntryWhere.Status.EQ(1),
		OrderBy("RANDOM()"),
		Limit(4), // 一旦4人対戦固定
	).AllG(context.Background())
	if err != nil {
		return nil, err
	}

	matchId := uuid.NewString()

	tx, err := boil.GetDB().(*sql.DB).Begin()
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		matchEntry := models.Match{
			ID:          matchId,
			EntryID:     entry.ID,
			ContestID:   contest.ID,
			BeforeScore: entry.Score,
			AfterScore:  entry.Score,
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

func RunMatchMake(contestId string) {
	var contestList []*models.Contest
	if contestId == "" {
		// TODO: 開催期間考慮
		contests, err := models.Contests().AllG(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		contestList = append(contestList, contests...)
	} else {
		contest, err := models.FindContestG(context.Background(), contestId)
		if err != nil {
			log.Fatal(err)
		}
		contestList = append(contestList, contest)
	}
	for _, contest := range contestList {
		_, err := makeMatch(contest)
		if err != nil {
			log.Print(err)
		}
	}
}
