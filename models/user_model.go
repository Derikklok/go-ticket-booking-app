package models

type User struct {
	Name  string `json:"name" bson:"name" binding:"required"`
	Email string `json:"email" bson:"email" binding:"required,email"`
	Age   int    `json:"age" bson:"age"`
}
