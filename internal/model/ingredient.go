package model

type Ingredient struct {
	ID   int
	Name string `gorm:"unique"`
	CatID string
}

type Unit struct {
	ID   int
	Name string `gorm:"unique"`
}

type TikiCate struct {
	Id int
	Name string
	CateID string
}
