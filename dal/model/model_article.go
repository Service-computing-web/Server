/*
 * Swagger Blog
 *
 * Simple Blog
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type Article struct {
	Id int64 `json:"id,omitempty"`

	Title string `json:"title,omitempty"`

	Username string `json:"username,omitempty"`

	Tags []Tag `json:"tags,omitempty"`

	Date string `json:"date,omitempty"`

	Content string `json:"content,omitempty"`

	Comments []Comment `json:"comments,omitempty"`
}
