package models

type Employee struct{
	ID            string			 `json:"id" bson:"_id"`
	First_name    string             `json:"first_name" validate:"required,min=2,max=100"`
	Last_name     string             `json:"last_name" validate:"required, min=2,max=100"`
	Email         string             `json:"email" validate:"email,required"`
	Phone         string             `json:"phone" validate:"required"`
	Salary		  float64			 `json:"salary" validate:"required"`
}