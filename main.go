package main

import(
	"os"
	"fmt"
	"log"
	"github.com/joho/godotenv"
	"gofr.dev/pkg/gofr"
)

func main(){
	app := gofr.New()
	err := godotenv.Load("configs/.env")
	if err != nil {
		log.Fatal("error in loading .env")
	}
	port := os.Getenv("HTTP_PORT")
	fmt.Println("server starting at port ", port)
	app.Start()
}