package main

import(
	"os"
	"fmt"
	"log"
	"hrms/routes"
	"github.com/joho/godotenv"
	"gofr.dev/pkg/gofr"
)

func main(){
	app := gofr.New()
	err := godotenv.Load("configs/.env")
	if err != nil {
		log.Fatalln("error in loading .env")
	}
	port := os.Getenv("HTTP_PORT")
	routes.EmployeeRoutes(app)
	fmt.Println("server starting at port ", port)
	app.Start()
}