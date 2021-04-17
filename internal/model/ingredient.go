package model

type Ingredient struct {
	ID           int           `json:"id" gorm:"primaryKey"`
	Name         string        `json:"name" gorm:"unique"`
	TikiCategory *TikiCategory `json:"tiki_category" gorm:"foreignKey:tiki_cate_id"`
	TikiCateID   int           `json:"tiki_cate_id"`
}

type Unit struct {
	ID   int    `json:"id"`
	Name string `json:"name" gorm:"unique"`
}

type TikiCategory struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	CatID string `json:"cat_id"`
}
