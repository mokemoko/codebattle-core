/*
 * CodeBattle API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"net/http"
	"time"

	"github.com/mokemoko/codebattle-core/server/models"

	"github.com/gin-gonic/gin"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func getUserIdFromJWT(c *gin.Context) string {
	// TODO: refactor
	claims, _ := authMiddleware.GetClaimsFromJWT(c)
	if claims == nil {
		return ""
	}
	if val, ok := claims[userIdKey]; ok {
		return val.(string)
	}
	return ""
}

// GetContestById -
func GetContestById(c *gin.Context) {
	userId := getUserIdFromJWT(c)
	contestId := c.Param("contestId")

	contest, err := models.Contests(
		models.ContestWhere.ID.EQ(contestId),
		Load("OwnerUser"),
	).OneG(c.Request.Context())

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	rankingEntries, err := contest.Entries(
		models.EntryWhere.ContestID.EQ(contestId),
		models.EntryWhere.Status.EQ(models.EntryStatusRegistered.Code),
		OrderBy("score desc"),
		Load("User"),
		Limit(20),
	).AllG(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	matches, err := contest.Matches(
		Load("Entry"),
		OrderBy("created_at desc"),
		OrderBy("id"),
		OrderBy("entry_id"),
		Limit(40),
	).AllG(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	res := Contest{
		Id: contest.ID,
		Owner: User{
			Id:   contest.R.OwnerUser.ID,
			Name: contest.R.OwnerUser.Name,
		},
		Name:          contest.Name,
		RecentMatches: []Match{},
		Ranking:       []Entry{},
	}
	if contest.Description.Valid {
		res.Description = contest.Description.String
	}
	users := map[string]User{}
	for _, entry := range rankingEntries {
		users[entry.R.User.ID] = User{
			Id:   entry.R.User.ID,
			Name: entry.R.User.Name,
			Icon: entry.R.User.Icon,
		}
		e := Entry{
			Id:    entry.ID,
			Name:  entry.Name,
			User:  users[entry.R.User.ID],
			Score: int32(entry.Score),
		}
		res.Ranking = append(res.Ranking, e)
	}
	tmpMatch := Match{}
	for _, match := range matches {
		if tmpMatch.Id == "" {
			createdAt, _ := time.Parse("2006-01-02T15:04:05Z", match.CreatedAt)
			tmpMatch.Id = match.ID
			tmpMatch.Type = models.NewMatchType(match.Type).Name
			tmpMatch.Status = models.NewMatchStatus(match.Status).Name
			tmpMatch.CreatedAt = createdAt
		} else if tmpMatch.Id != match.ID {
			res.RecentMatches = append(res.RecentMatches, tmpMatch)
			createdAt, _ := time.Parse("2006-01-02T15:04:05Z", match.CreatedAt)
			tmpMatch = Match{
				Id:        match.ID,
				Type:      models.NewMatchType(match.Type).Name,
				Status:    models.NewMatchStatus(match.Status).Name,
				CreatedAt: createdAt,
			}
		}
		tmpMatch.Entries = append(tmpMatch.Entries, MatchEntry{
			Id:          match.EntryID,
			Name:        match.R.Entry.Name,
			User:        users[match.R.Entry.UserID],
			Rank:        int32(match.Rank),
			BeforeScore: int32(match.BeforeScore),
			AfterScore:  int32(match.AfterScore),
		})
	}
	if tmpMatch.Id != "" {
		res.RecentMatches = append(res.RecentMatches, tmpMatch)
	}

	if userId != "" {
		// ログイン済みの場合は自身のエントリー情報を返す
		res.OwnEntries = []Entry{}

		ownEntries, err := contest.Entries(
			models.EntryWhere.ContestID.EQ(contestId),
			models.EntryWhere.UserID.EQ(userId),
			OrderBy("created_at"),
		).AllG(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		for _, entry := range ownEntries {
			e := Entry{
				Id:         entry.ID,
				Name:       entry.Name,
				Score:      int32(entry.Score),
				Error:      entry.Error.String,
				Status:     models.NewEntryStatus(entry.Status).Name,
				Repository: entry.Repository,
			}
			res.OwnEntries = append(res.OwnEntries, e)
		}
	}

	c.JSON(http.StatusOK, res)
}

// GetContests -
func GetContests(c *gin.Context) {
	contests, err := models.Contests().AllG(c.Request.Context())
	if err != nil {
		return
	}
	res := []GetContests200ResponseInner{}
	for _, contest := range contests {
		res = append(res, GetContests200ResponseInner{
			Id:   contest.ID,
			Name: contest.Name,
		})
	}
	c.JSON(http.StatusOK, res)
}

// PutMatch -
func PutMatch(c *gin.Context) {
	userId := getUserIdFromJWT(c)
	contestId := c.Param("contestId")

	var json PutMatchRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	matchId := uuid.NewString()
	tx, err := boil.GetDB().(*sql.DB).Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, entryId := range json.EntryIds {
		entry, err := models.Entries(
			models.EntryWhere.ID.EQ(entryId),
		).OneG(c.Request.Context())
		if err != nil {
			// ErrNoRows含めて500として処理
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		match := models.Match{
			ID:        matchId,
			EntryID:   entryId,
			ContestID: contestId,
			UserID:    null.StringFrom(userId),
			// 手動登録されたものはUnratedとする
			Type:        models.MatchTypeUnrated.Code,
			BeforeScore: entry.Score,
			AfterScore:  entry.Score,
		}
		if err = match.Insert(c.Request.Context(), tx, boil.Infer()); err != nil {
			_ = tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	if err = tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
