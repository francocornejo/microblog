package repository

import (
	"fmt"
	"log"
	"microblog/models"
	"net/http"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SQLRepository interface {
	SendMessageRepository(bodyMessage *models.Message) (*models.Message, *models.ErrorMessage)
	FollowRepository(bodyFollowers *models.UsernameFollower) (*models.Follower, *models.ErrorMessage)
	TimelineRepository(bodyTimeline *models.Timeline) ([]models.Feed, *models.ErrorMessage)
}

type sqlRepository struct {
	BDMicroblog *gorm.DB
}

func NewSQLConnection(driver, url string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Verifica la conexión a la base de datos utilizando Ping
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	log.Println("[SQL CONNECTION] Conexion exitosa!")
	return db, nil
}

func NewSQLRepository() SQLRepository {
	db, err := NewSQLConnection("mysql", os.Getenv("SQL_CONNECTION"))
	if err != nil {
		log.Println(fmt.Sprintf("[SQL CONNECTION] %v", err))
	}
	return &sqlRepository{
		BDMicroblog: db,
	}
}

func (s *sqlRepository) SendMessageRepository(bodyMessage *models.Message) (*models.Message, *models.ErrorMessage) {
	tableMessage := os.Getenv("TABLE_MESSAGES")
	tableUser := os.Getenv("TABLE_USERS")
	var user models.User

	validateUser := s.BDMicroblog.Table(tableUser).Find(&user, "username = ?", bodyMessage.Username)
	if validateUser.RowsAffected == 0 {
		errMessage := models.ErrorResponse("| Error | ", "No se encontró el usuario", http.StatusNotFound, nil)
		log.Println(errMessage)
		return nil, &errMessage
	}

	if validateUser.Error != nil {
		errMessage := models.ErrorResponse("| Error | ", "Ocurrio un error validando el usuario", http.StatusInternalServerError, nil)
		log.Println(errMessage)
		return nil, &errMessage
	}

	newMessage := &models.Message{
		Username:  bodyMessage.Username,
		Text:      bodyMessage.Text,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	}

	res := s.BDMicroblog.Table(tableMessage).Create(newMessage)
	if res.Error != nil {
		errMessage := models.ErrorResponse("| Error | ", "No se pudo crear el mensaje |", http.StatusNotFound, nil)
		log.Printf(fmt.Sprintf("[InsertarUsuario] %v", res.Error))
		return nil, &errMessage
	}
	return newMessage, nil
}

func (s *sqlRepository) FollowRepository(bodyFollowers *models.UsernameFollower) (*models.Follower, *models.ErrorMessage) {
	tableFollowers := os.Getenv("TABLE_FOLLOWERS")
	tableUsers := os.Getenv("TABLE_USERS")
	var existingFollower models.Follower
	var userCountId models.CountID

	//Verifico que el usuario exista
	result := s.BDMicroblog.Table(tableUsers).Where("username = ?", bodyFollowers.Username).First(&userCountId)
	if result.RowsAffected == 0 {
		userNull := fmt.Sprintf("El usuario %s no existe", bodyFollowers.Username)
		errMessage := models.ErrorResponse("Error", userNull, http.StatusNotFound, nil)
		log.Println(errMessage)
		return nil, &errMessage
	}

	//Guardo el ID del usuario
	existingFollower.UserID = uint(userCountId.ID)

	//reinicio el contador
	userCountId.ID = 0

	//Verifico que el usuario exista
	result = s.BDMicroblog.Table(tableUsers).Where("username = ?", bodyFollowers.FollowerUsername).First(&userCountId)
	if result.RowsAffected == 0 {
		userNull := fmt.Sprintf("El usuario %s no existe", bodyFollowers.FollowerUsername)
		errMessage := models.ErrorResponse("Error", userNull, http.StatusBadRequest, nil)
		log.Println(errMessage)
		return nil, &errMessage
	}

	existingFollower.FollowerID = uint(userCountId.ID)

	//Reviso que los usuarios no se estén siguiendo
	result = s.BDMicroblog.Table(tableFollowers).Where("user_id = ? AND follower_id = ?", existingFollower.UserID, existingFollower.FollowerID).First(&existingFollower)
	if result.RowsAffected > 0 {
		// El usuario ya está siguiendo al otro, no es necesario hacer nada
		return &existingFollower, nil
	}

	//Se crea el seguidor
	result = s.BDMicroblog.Table(tableFollowers).Create(&existingFollower)
	if result.RowsAffected > 0 {
		return &existingFollower, nil
	}

	errMessage := models.ErrorResponse(" | Error | ", "Ocurrio un error al dar follow |", http.StatusInternalServerError, nil)
	log.Println(errMessage)
	return nil, &errMessage
}

func (s *sqlRepository) TimelineRepository(bodyTimeline *models.Timeline) ([]models.Feed, *models.ErrorMessage) {
	tableFollowers := os.Getenv("TABLE_FOLLOWERS")
	tableMessages := os.Getenv("TABLE_MESSAGES")
	tableUser := os.Getenv("TABLE_USERS")
	var timeline models.Timeline
	var followersPerUser []models.FollowerPerUser
	var messagesFeed []models.Feed

	//Verifico que el usuario exista
	result := s.BDMicroblog.Table(tableUser).Where("username= ?", &bodyTimeline.Username).First(&timeline)
	if result.RowsAffected == 0 {
		errMessage := models.ErrorResponse(" | Error | ", "El usuario no existe", http.StatusNotFound, nil)
		log.Println(errMessage)
		return nil, &errMessage
	}

	//Obtengo informacion de los seguidores del usuario
	result = s.BDMicroblog.
		Table(tableFollowers).
		Select("followers.id, followers.follower_id, followers_users.username AS follower_username").
		Joins("JOIN users ON followers.user_id = users.id").
		Joins("JOIN users AS followers_users ON followers.follower_id = followers_users.id").
		Where("users.username = ?", &bodyTimeline.Username).
		Find(&followersPerUser)

	if result.RowsAffected > 0 {
		var usernames []string
		for _, follower := range followersPerUser {
			usernames = append(usernames, follower.FollowerUsername)
		}

		//Obtengo los mensajes y los ordeno del mas nuevo al mas antiguo
		result = s.BDMicroblog.Table(tableMessages).Where("username IN (?)", usernames).Order("timestamp DESC").Find(&messagesFeed)
		if result.RowsAffected > 0 {
			return messagesFeed, nil
		}
		errMessage := models.ErrorResponse(" | Error | ", "Ocurrio un error al obtener mensajes", http.StatusInternalServerError, nil)
		log.Println(errMessage)
		return nil, &errMessage
	}

	if result.Error != nil {
		errMessage := models.ErrorResponse(" | Error | ", "Ocurrio un error al obtener mensajes", http.StatusInternalServerError, nil)
		log.Println(errMessage)
		return nil, &errMessage
	}

	return []models.Feed{}, nil
}
