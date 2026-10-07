package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init(envPath string) {
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading .env file from %s: %v", envPath, err)
	}

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", connect)

	if err != nil {
		log.Fatalf("Failed to open DB connection: %v", err)
	}
}

func PrepareInsert(data [][]string) string {
	baseQuery := `INSERT INTO test (id, customer_id, first_name, last_name, company, city, country, email, sub_date, website)
						VALUES %s`
	values := []string{}
	for _, lineData := range data {
		lineData = formatRow(lineData)
		values = append(values, "(" + strings.Join(lineData, ",") + ")")
	}

	finalQuery := fmt.Sprintf(baseQuery, strings.Join(values, ",")) + ";"
	return finalQuery
}

func formatRow(lineData []string) []string {
	for index, data := range lineData {
		if index != 0 {
			lineData[index] = "'" + strings.ReplaceAll(data, "'", "") + "'"
		}
	}
	return lineData
}