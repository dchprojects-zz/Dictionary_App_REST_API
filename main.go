package main

import (
	"fmt"
	"net/http"
	"os"
	"path"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {

	server := echo.New()
	server.GET(path.Join("/"), Version)

	godotenv.Load()
	port := os.Getenv("PORT")

	address := fmt.Sprintf("%s:%s", "0.0.0.0", port)

	server.Start(address)
	// database, err := db.CreateDatabase()
	// if err != nil {
	// 	log.Fatal("Database connection failed: %s", err.Error())
	// }

	// app := &app.App{
	// 	Router:   mux.NewRouter().StrictSlash(true),
	// 	Database: database,
	// }

	// app.SetupRouter()

}

func Version(context echo.Context) error {
	return context.JSON(http.StatusOK, map[string]interface{}{"version": 1})
}
