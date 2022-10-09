package model

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

type Url struct {
	ID       uint64 `json:"id"`
	LongUrl  string `json:"longUrl" gorm:"not null;default:null"`
	ShortUrl string `json:"shortUrl" gorm:"unique;not null"`
}

type UrlClick struct {
	UrlID     uint64 `json:"urlId"`
	Url       Url
	ClickedAt time.Time `json:"clickedAt"`
}

type UrlTag struct {
	UrlID uint64 `json:"urlId"`
	Url   Url
	Key   string `json:"key"`
	Value string `json:"value"`
}

func Setup(config *Config) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Url{}, &UrlClick{}, &UrlTag{})
	if err != nil {
		fmt.Println(err)
	}
}
