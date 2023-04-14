package connection

import (
	"log"
	"os"

	"github.com/Thrashy190/info-center-api/pkg/models"
	"github.com/Thrashy190/info-center-api/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ModelAutoMigrations() {
	var err error
	err = DB.AutoMigrate(models.Careers{}, models.Users{}, models.Book{}, models.Tesis{}, models.Projects{}, models.Request{}, models.RequestInfo{})

	if err != nil {
		utils.Error(err.Error())
	}
}

func Dbconnection() {

	host := " host=" + os.Getenv("DB_HOST")
	user := " user=" + os.Getenv("DB_USERNAME")
	password := " password=" + os.Getenv("DB_PASSWORD")
	dbname := " dbname=" + os.Getenv("DB_NAME")
	port := " port=" + os.Getenv("DB_PORT")

	var DSN = host + user + password + dbname + port
	var err error

	if DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{}); err != nil {
		utils.Error(err.Error())
		return
	}

	//TODO check close connection

	defer func() {
		sqlDB, err := DB.DB()
		if err != nil {
			log.Fatalln(err)
		}
		_ = sqlDB.Close()
	}()

	utils.Succes("DB connected successfully")
	log.Println(DB)

}
