package models

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

type User struct {
	UserID       int       `json:"user_id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password"`
	CreatedAt    time.Time `json:"created_at"`
}

type Ticket struct {
	TicketID    int       `json:"ticket_id"`
	CreatorID   int       `json:"creator_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Task struct {
	TaskID      int       `json:"task_id"`
	TicketID    int       `json:"ticket_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TicketAssignment struct {
	AssignmentID int       `json:"assignment_id"`
	TicketID     int       `json:"ticket_id"`
	AssigneeID   int       `json:"assignee_id"`
	AssignedAt   time.Time `json:"assigned_at"`
}

const (
	defaultHost     = "localhost"
	defaultPort     = 5432
	defaultUser     = "postgres"
	defaultPassword = "postgres"
	defaultDbname   = "practice"
)

func getEnvOrDefault(envVar string, defaultValue string) string {
	value := os.Getenv(envVar)
	if value == "" {
		return defaultValue
	}
	return value
}

func InitializeDB() {
	var err error
	host := getEnvOrDefault("PGHOST", defaultHost)
	port := getEnvOrDefault("PGPORT", fmt.Sprintf("%d", defaultPort))
	user := getEnvOrDefault("PGUSER", defaultUser)
	password := getEnvOrDefault("PGPASSWORD", defaultPassword)
	dbname := getEnvOrDefault("PGDATABASE", defaultDbname)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
}
