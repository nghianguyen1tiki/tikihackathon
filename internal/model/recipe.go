package model

type Recipe struct {
	ID          int
	Name        string
	Description string `gorm:"types:text"`
	Rating      int
	Servings    int
	Author      *User
	Ingredients []Including
	Steps       []Step
}

type Including struct {
	RecipeID     int `gorm:"primaryKey"`
	IngredientID int `gorm:"primaryKey"`
	UnitID       int `gorm:"not null"`
	Quantity     int
	Ingredient   Ingredient
	Unit         Unit
}

type Step struct {
	ID         int
	Content    string `gorm:"types:text"`
	RecipeID   int
	StepPhotos []StepPhoto
}

type StepPhoto struct {
	ID     int
	URL    string `gorm:"types:text"`
	Height int
	Width  int
	StepID int
}
