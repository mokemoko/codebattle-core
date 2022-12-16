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
	"net/http"

	"github.com/mokemoko/codebattle-core/server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// PostEntry -
func PostEntry(c *gin.Context) {
	userId, _ := c.Get(userIdKey)

	var json PostEntryRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	entry := models.Entry{
		ID:         uuid.NewString(),
		UserID:     userId.(string),
		ContestID:  json.ContestId,
		Name:       json.Name,
		Repository: json.Repository,
	}
	if err := entry.InsertG(c.Request.Context(), boil.Infer()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, entry)
}

// PutEntry -
func PutEntry(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
