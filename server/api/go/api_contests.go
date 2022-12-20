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

	// TODO: 自身のentryは別APIで取得するようにしてregisteredに絞る
	entries, err := contest.Entries(
		OrderBy("score desc"),
		Load("User"),
	).AllG(c.Request.Context())
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

	matches, err := contest.Matches(
		Load("Entry"),
		OrderBy("created_at desc"),
		Limit(20),
	).AllG(c.Request.Context())
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
	for _, entry := range entries {
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
		// 自身のentryは追加情報返す
		if entry.R.User.ID == userId {
			e.Error = entry.Error.String
			e.Status = models.NewEntryStatus(entry.Status).Name
			e.Repository = entry.Repository
		}
		res.Ranking = append(res.Ranking, e)
	}
	tmpMatch := Match{}
	for _, match := range matches {
		if tmpMatch.Id == "" {
			createdAt, _ := time.Parse("2006-01-02T15:04:05Z", match.CreatedAt)
			tmpMatch.Id = match.ID
			tmpMatch.Type = models.NewMatchType(match.Type).Name
			tmpMatch.CreatedAt = createdAt
		} else if tmpMatch.Id != match.ID {
			res.RecentMatches = append(res.RecentMatches, tmpMatch)
			createdAt, _ := time.Parse("2006-01-02T15:04:05Z", match.CreatedAt)
			tmpMatch = Match{
				Id:        match.ID,
				Type:      models.NewMatchType(match.Type).Name,
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
	c.JSON(http.StatusOK, gin.H{})
}
