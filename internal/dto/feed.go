package dto

import "github.com/nghiant3223/tikihackathon/internal/model"

type GetFeedResponse struct {
	PopularRecipes      []model.Recipe `json:"popular_recipes"`
	PersonalizedRecipes []model.Recipe `json:"personalized_recipes"`
}

type GetFeedQuery struct {
	Page      string   `form:"page" binding:"required"`
	Size      string   `form:"size" binding:"required"`
	Blacklist []string `form:"blacklist[]"`
	Whitelist []string `form:"whitelist[]"`
	Pantry    []string `form:"pantry[]"`
}

type SearchQuery struct {
	Page      string   `form:"page" binding:"required"`
	Size      string   `form:"size" binding:"required"`
	Query     string   `form:"q" binding:"required"`
	Blacklist []string `form:"blacklist[]"`
	Whitelist []string `form:"whitelist[]"`
	Pantry    []string `form:"pantry[]"`
}
