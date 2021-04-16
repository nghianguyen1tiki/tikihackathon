package model

type Recipe struct {
	ID          int
	Name        string `gorm:"unique"`
	Description string `gorm:"types:text"`
	Rating      float32
	Servings    int
	TotalTime   int
	Difficulty  int
	Author      *User `gorm:"foreignKey:UserID"`
	UserID      *int
	Ingredients []Including
	Steps       []Step
	Photo       RecipePhoto
}

type Including struct {
	RecipeID     int `gorm:"primaryKey"`
	IngredientID int `gorm:"primaryKey"`
	UnitID       int `gorm:"not null"`
	Quantity     string
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

type RecipePhoto struct {
	ID       int
	URL      string `gorm:"types:text"`
	Height   int
	Width    int
	RecipeID int
}
