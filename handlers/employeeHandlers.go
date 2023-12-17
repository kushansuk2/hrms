package handlers

import (
	"hrms/models"
	db "hrms/database"
	"log"
	"gofr.dev/pkg/gofr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var employeeCollection *mongo.Collection = db.OpenCollection("employee")

func GetAllEmployees() gofr.Handler{ 
	return func(ctx *gofr.Context) (interface{}, error){
		query := bson.D{{}}
		cursor, err := employeeCollection.Find(ctx.Context,query)
		if(err!=nil){
			log.Fatalln(err)
		}
		var employees []models.Employee = make([]models.Employee,0)
		if err := cursor.All(ctx.Context,&employees); err != nil {
			log.Fatalln(err)
			return nil, err
		}
		return employees, nil
	}
}

func CreateEmployee() gofr.Handler{ 
	return func(ctx *gofr.Context) (interface{}, error){
		var employee models.Employee
		if err := ctx.Bind(&employee); err != nil {
			log.Fatalln(err)
			return nil, err
		}
		employee.ID = ""
		insertionResult, err := employeeCollection.InsertOne(ctx.Context,employee)
		if(err != nil){
			log.Fatalln(err)
			return nil, err
		}
		filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
		createdRecord := employeeCollection.FindOne(ctx.Context,filter)
		createdEmployee := &models.Employee{}
		createdRecord.Decode(createdEmployee)
		return createdEmployee, nil
	}
}

func UpdateEmployee() gofr.Handler{
	return func(ctx *gofr.Context) (interface{},error){
		idParam := ctx.Param("id")
		employeeId, err := primitive.ObjectIDFromHex(idParam)
		if(err != nil){
			log.Fatalln(err)
			return nil, err
		}
		var employee models.Employee
		if err := ctx.Bind(&employee); err != nil {
			log.Fatalln(err)
			return nil, err
		}
		query := bson.D{{Key: "_id",Value: employeeId}}
		update := bson.D{
			{

			}
		}
	}
}