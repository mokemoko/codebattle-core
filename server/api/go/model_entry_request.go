/*
 * CodeBattle API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type EntryRequest struct {

	ContestId string `json:"contestId"`

	Name string `json:"name"`

	Repository string `json:"repository"`

	IsDisabled bool `json:"isDisabled,omitempty"`
}
