package model

type Blog struct{
	ID unit `json:"id" gorm: "primaryKey"`
}