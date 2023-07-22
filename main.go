package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/delordemm1/devemm-go/internal/cmd"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
	appDebug := os.Getenv("APP_DEBUG") == "true"
	// dbName := flag.String("dbname", "devemm", "Database name")
	// dbUser := flag.String("dbuser", "devemm", "Database user")
	// dbPassword := flag.String("dbpassword", "devemm", "Database password")
	// dbHost := flag.String("dbhost", "localhost", "Database host")
	// dbPort := flag.String("dbport", "5432", "Database port")
	appPort := flag.String("port", "4000", "Application port")
	flag.Parse()
	log.Printf("appDebug %v", appDebug)
	switch os.Args[1] {
	case "serve":
		cmd.WebServe(appDebug, fmt.Sprintf(":%s", *appPort), fmt.Sprintf("http://localhost:%s", *appPort))

	default:
		printUsage()
		os.Exit(1)
	}

}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  serve - Run web server")
}
