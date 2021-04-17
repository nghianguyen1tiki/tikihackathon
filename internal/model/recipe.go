package model

type Recipe struct {
	ID          int         `json:"id"`
	Name        string      `json:"name" gorm:"unique"`
	Description string      `json:"description" gorm:"types:text"`
	Rating      float32     `json:"rating"`
	Servings    int         `json:"servings"`
	TotalTime   int         `json:"total_time"`
	TotalView   int         `json:"total_view" gorm:"index"`
	Difficulty  int         `json:"difficulty"`
	Author      *User       `json:"author" gorm:"foreignKey:UserID"`
	UserID      *int        `json:"user_id"`
	Ingredients []Including `json:"ingredients"`
	Steps       []Step      `json:"steps"`
	Photo       RecipePhoto `json:"photo"`
}

type Including struct {
	RecipeID     int        `json:"recipe_id" gorm:"primaryKey"`
	IngredientID int        `json:"ingredient_id" gorm:"primaryKey"`
	UnitID       int        `json:"unit_id" gorm:"not null"`
	Quantity     string     `json:"quantity"`
	Ingredient   Ingredient `json:"ingredient"`
	Unit         Unit       `json:"unit"`
}

type Step struct {
	ID         int         `json:"id"`
	Content    string      `json:"content" gorm:"types:text"`
	RecipeID   int         `json:"recipe_id"`
	StepPhotos []StepPhoto `json:"step_photos"`
}

type StepPhoto struct {
	ID     int    `json:"id"`
	URL    string `json:"url" gorm:"types:text"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
	StepID int    `json:"step_id"`
}

type RecipePhoto struct {
	ID       int
	URL      string `gorm:"types:text"`
	Height   int
	Width    int
	RecipeID int
}
