package utils

// import (
// 	"log"
// 	"os"
// 	"regexp"

// 	"github.com/joho/godotenv"
// )

// const projectDirName = "Database-Transactions-Simulation"

// func LoadEnv() {
// 	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
// 	currentWorkDirectory, _ := os.Getwd()
// 	rootPath := projectName.Find([]byte(currentWorkDirectory))

// 	err := godotenv.Load(string(rootPath) + "/.env")

// 	if err != nil {
// 		log.Fatalf("Error loading .env file", err)
// 	}
// }
