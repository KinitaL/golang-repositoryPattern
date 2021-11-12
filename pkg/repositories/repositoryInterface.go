package repositories

import (
	"log"
	"os"
	"repository/pkg/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type NoteRepositoryInterface interface {
	GetOne(id int) models.Note
	Get() []models.Note
	Save(note models.Note) string
}

func envLoad() {
	if err := godotenv.Load("./repo.env"); err != nil {
		log.Fatal("No .env file was found")
	}
}
func envManage(key string) string {

	value, exists := os.LookupEnv(key)
	if exists {
		return value
	} else {
		log.Fatal("No conn")
		return ""
	}
}

func connectToDB() *gorm.DB {
	envLoad()
	host := envManage("POSTGRES_HOST")
	port := envManage("POSTGRES_PORT")
	user := envManage("POSTGRES_USER")
	pass := envManage("POSTGRES_PASSWORD")
	dbname := envManage("POSTGRES_DB")
	helper := "host=" + host + " port=" + port + " user=" + user + " password=" + pass + " dbname=" + dbname + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(helper), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
