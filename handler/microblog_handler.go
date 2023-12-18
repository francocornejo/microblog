package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"microblog/models"
	"microblog/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type MicroBlogHandler interface {
	SendMessageHandler(ctx *gin.Context)
	FollowHandler(ctx *gin.Context)
	TimelineHandler(ctx *gin.Context)
}

type microBlogHandler struct {
	service service.Service
}

func NewMicroblogHandler(service service.Service) MicroBlogHandler {
	return &microBlogHandler{
		service: service,
	}
}

// @Description Envio de mensajes
// @Tags Endpoints
// @Accept  json
// @Produce  json
// @Param        Request  body      models.Message  true  "Body para el envío de mensajes"
// @Success 200 {object} models.Message
// @Failure 404 {object} models.ErrorMessage
// @Failure 503 {object} models.ErrorMessage
// @Router /send [post]
func (m *microBlogHandler) SendMessageHandler(ctx *gin.Context) {
	var (
		bodyMessage models.Message
	)

	err := ctx.BindJSON(&bodyMessage)
	if err != nil {
		errMessage := ManejarErroresBindJSON(err)
		ctx.JSON(errMessage.Status, errMessage)
		return
	}

	message, errResponse := m.service.SendMessageService(bodyMessage)
	if errResponse != nil {
		switch {
		case errResponse.Status == 404:
			ctx.JSON(http.StatusNotFound, errResponse)
			return
		default:
			ctx.JSON(http.StatusInternalServerError, errResponse)
			return
		}
	}
	ctx.JSON(http.StatusOK, message)
}

// @Description Seguir a otros usuarios
// @Tags Endpoints
// @Accept  json
// @Produce  json
// @Param        Request  body      models.UsernameFollower  true  "Body para seguir al usuario deseado"
// @Success 200 {object} models.Follower
// @Failure 400 {object} models.ErrorMessage
// @Failure 404 {object} models.ErrorMessage
// @Failure 503 {object} models.ErrorMessage
// @Router /follow [post]
func (m *microBlogHandler) FollowHandler(ctx *gin.Context) {
	var bodyFollowers models.UsernameFollower

	//Obtengo los 2 usuarios Username como seguidor y FollowUsername como el usuario seguido
	err := ctx.BindJSON(&bodyFollowers)
	if err != nil {
		errMessage := ManejarErroresBindJSON(err)
		ctx.JSON(errMessage.Status, errMessage)
		return
	}

	follow, errResponse := m.service.FollowService(bodyFollowers)
	if errResponse != nil {
		switch {
		case errResponse.Status == 404:
			ctx.JSON(http.StatusNotFound, errResponse)
			return
		case errResponse.Status == 400:
			ctx.JSON(http.StatusBadRequest, errResponse)
			return
		case errResponse.Status == 500:
			ctx.JSON(http.StatusInternalServerError, errResponse)
			return
		}
	}

	ctx.JSON(http.StatusOK, follow)
}

// @Description Obtener mensajes de los usuarios a los que sigue
// @Tags Endpoints
// @Accept  json
// @Produce  json
// @Param        Request  body      models.Timeline  true  "Body para obtener mensajes de los usuarios a los que sigue"
// @Success 200 {object} models.Follower
// @Failure 400 {object} models.ErrorMessage
// @Failure 404 {object} models.ErrorMessage
// @Failure 503 {object} models.ErrorMessage
// @Router /messages [get]
func (m *microBlogHandler) TimelineHandler(ctx *gin.Context) {
	var bodyTimeline models.Timeline

	err := ctx.BindJSON(&bodyTimeline)
	if err != nil {
		errMessage := ManejarErroresBindJSON(err)
		ctx.JSON(errMessage.Status, errMessage)
		return
	}

	timeline, errResponse := m.service.TimelineService(bodyTimeline)
	if errResponse != nil {
		switch {
		case errResponse.Status == 404:
			ctx.JSON(http.StatusNotFound, errResponse)
			return
		case errResponse.Status == 400:
			ctx.JSON(http.StatusBadRequest, errResponse)
			return
		case errResponse.Status == 500:
			ctx.JSON(http.StatusInternalServerError, errResponse)
			return
		}
	}
	ctx.JSON(http.StatusOK, timeline)
}

//Manejo de errores cuando parseo el json
func ManejarErroresBindJSON(err error) models.ErrorMessage {
	var (
		validationErrors []models.Error
	)

	if ute, ok := err.(*json.UnmarshalTypeError); ok {
		return models.ErrorResponse("Error de Tipo en el campo "+ute.Field, fmt.Sprintf("Expected: %s - Actual: %s", ute.Type, ute.Value), http.StatusBadRequest, nil)
	}
	if err, ok := err.(validator.ValidationErrors); ok {
		validationErrors = loopFor(err, validationErrors)
	}

	if validationErrors != nil {
		return models.ErrorResponse("Error de validacion", err.Error(), http.StatusBadRequest, validationErrors)
	}

	return models.ErrorResponse("Error de validacion", err.Error(), http.StatusBadRequest, nil)
}

//Procesa errores del validator y los convierte a lista de errores personalizados
func loopFor(e error, valErrs []models.Error) []models.Error {
	var valErr models.Error
	for _, err := range e.(validator.ValidationErrors) {
		valErr.Field = err.Field()
		if err.Param() != "" {
			valErr.Message = err.Tag() + ": " + err.Param()
		} else {
			valErr.Message = err.Tag()
		}
		valErrs = append(valErrs, valErr)
	}
	//devuelvo lista de errores
	return valErrs
}
