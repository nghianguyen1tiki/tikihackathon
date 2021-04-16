package model

type Ingredient struct {
	ID   int
	Name string `gorm:"unique"`
}

type Unit struct {
	ID   int
	Name string `gorm:"unique"`
}
