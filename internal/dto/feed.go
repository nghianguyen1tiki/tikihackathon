package dto

import "github.com/nghiant3223/tikihackathon/internal/model"

type GetFeedResponse struct {
	PopularRecipes      []model.Recipe `json:"popular_recipes"`
	PersonalizedRecipes []model.Recipe `json:"personalized_recipes"`
}

type GetFeedQuery struct {
	Page      int   `form:"page" binding:"required"`
	Size      int   `form:"size" binding:"required"`
	Blacklist []int `form:"blacklist"`
	Whitelist []int `form:"whitelist"`
	Pantry    []int `form:"pantry"`
}
