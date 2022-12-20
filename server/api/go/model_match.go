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
	"time"
)

type Match struct {

	Id string `json:"id"`

	Type string `json:"type"`

	Status string `json:"status"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Entries []MatchEntry `json:"entries"`
}
