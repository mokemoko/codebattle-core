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
	"github.com/gin-gonic/gin"
	"github.com/mokemoko/codebattle-core/server/models"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"net/http"
	"time"
)

// GetContestById -
func GetContestById(c *gin.Context) {
	contestId := c.Param("contestId")

	contest, err := models.Contests(
		Where("id = ?", contestId),
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

	entries, err := contest.Entries(Load("User")).AllG(c.Request.Context())
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

	matches, err := contest.Matches(Load("Entry")).AllG(c.Request.Context())
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
		Name: contest.Name,
	}
	if contest.Description.Valid {
		res.Description = contest.Description.String
	}
	users := map[string]User{}
	for _, entry := range entries {
		users[entry.R.User.ID] = User{
			Id:   entry.R.User.ID,
			Name: entry.R.User.Name,
		}
		res.Ranking = append(res.Ranking, PublicEntry{
			Id:    entry.ID,
			Name:  entry.Name,
			User:  users[entry.R.User.ID],
			Score: int32(entry.Score),
		})
	}
	tmpMatch := Match{}
	for _, match := range matches {
		if tmpMatch.Id == "" {
			createdAt, _ := time.Parse("2006-01-02T15:04:05Z", match.CreatedAt)
			tmpMatch.Id = match.ID
			tmpMatch.CreatedAt = createdAt
		} else if tmpMatch.Id != match.ID {
			res.RecentMatches = append(res.RecentMatches, tmpMatch)
			createdAt, _ := time.Parse("2006-01-02T15:04:05Z", match.CreatedAt)
			tmpMatch = Match{
				Id:        match.ID,
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
