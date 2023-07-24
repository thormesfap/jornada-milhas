package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/thormesfap/jornada-milhas/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	erro := godotenv.Load(".env")
	if erro != nil {
		log.Panic("Erro ao carregar variáveis de ambiente (arquivo .env não encontrado)")
	}
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	Migrate()
}

func Migrate(){
	DB.AutoMigrate(&models.Depoimento{})
}

