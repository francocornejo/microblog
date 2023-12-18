package service_test

import (
	"net/http"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"microblog/models"
	mock_repository "microblog/repository/mock"
	"microblog/service"
)

var _ = Describe("Service", func() {
	var (
		mockCtrl         *gomock.Controller
		mockSql          *mock_repository.MockSQLRepository
		serviceMicroblog service.Service
		bodyOK           models.Message
		bodyErr          models.Message
		bodyFollowOK     models.UsernameFollower
		bodyFollowErr    models.UsernameFollower
		bodyFollowers    models.Follower
		viewMessagesOK   models.Timeline
		viewMessagesErr  models.Timeline
		feed             []models.Feed
	)

	bodyOK = models.Message{
		Username:  "usuarioExistente",
		Text:      "Tengo de prueba OK",
		Timestamp: "2023-12-17 21:16:51",
	}

	bodyErr = models.Message{
		Username:  "usuarioInexistente",
		Text:      "Tengo de prueba OK",
		Timestamp: "2023-12-17 21:16:51",
	}

	bodyFollowOK = models.UsernameFollower{
		Username:         "usuarioSeguidor",
		FollowerUsername: "usuarioSeguido",
	}

	bodyFollowers = models.Follower{
		UserID:     1,
		FollowerID: 2,
	}

	bodyFollowErr = models.UsernameFollower{
		Username:         "usuarioSeguidor-ERROR",
		FollowerUsername: "usuarioSeguido-ERROR",
	}

	viewMessagesOK = models.Timeline{
		Username: "usuarioExistente",
	}

	viewMessagesErr = models.Timeline{
		Username: "usuarioInexistente",
	}

	feed = []models.Feed{
		{
			Username:  "usuario1",
			Text:      "Texto Nuevo",
			Timestamp: "2023-12-17 21:16:51",
		},
		{
			Username:  "usuario2",
			Text:      "Texto Viejo",
			Timestamp: "2023-12-13 21:16:51",
		},
	}

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockSql = mock_repository.NewMockSQLRepository(mockCtrl)
		serviceMicroblog = service.NewService(mockSql)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("ENVIAR MENSAJES", func() {
		It("Se envia el mensaje con Exito", func() {
			mockSql.
				EXPECT().
				SendMessageRepository(&bodyOK).
				Times(1).
				Return(&bodyOK, nil)

			sendMsg, errMessage := serviceMicroblog.SendMessageService(bodyOK)
			Expect(errMessage).To(BeNil())
			Expect(sendMsg).NotTo(BeNil())
		})

		It("No se puede publicar el mensaje", func() {
			expectedErrMessage := models.ErrorResponse("| Error | ", "No se encontró el usuario", http.StatusNotFound, nil)
			mockSql.
				EXPECT().
				SendMessageRepository(&bodyErr).
				Times(1).
				Return(nil, &expectedErrMessage)

			sendMsg, errMessage := serviceMicroblog.SendMessageService(bodyErr)
			Expect(errMessage).NotTo(BeNil())
			Expect(sendMsg).To(BeNil())
		})

		It("Error al publicar mensaje. Error en el servidor", func() {
			expectedErrMessage := models.ErrorResponse(" | Error | ", "Ocurrio un error al publicar mensaje", http.StatusInternalServerError, nil)
			mockSql.
				EXPECT().
				SendMessageRepository(&bodyErr).
				Times(1).
				Return(nil, &expectedErrMessage)

			sendMsg, errMessage := serviceMicroblog.SendMessageService(bodyErr)
			Expect(errMessage).NotTo(BeNil())
			Expect(sendMsg).To(BeNil())
		})
	})

	Describe("SEGUIR USUARIOS", func() {
		It("Sigo un usuario con Exito", func() {
			mockSql.
				EXPECT().
				FollowRepository(&bodyFollowOK).
				Times(1).
				Return(&bodyFollowers, nil)

			sendFollow, errMessage := serviceMicroblog.FollowService(bodyFollowOK)
			Expect(errMessage).To(BeNil())
			Expect(sendFollow).NotTo(BeNil())
		})

		It("No es posible seguir el usuario", func() {
			expectedErrMessage := models.ErrorResponse("| Error | ", "No se encontró el usuario", http.StatusNotFound, nil)
			mockSql.
				EXPECT().
				FollowRepository(&bodyFollowErr).
				Times(1).
				Return(nil, &expectedErrMessage)

			sendFollow, errMessage := serviceMicroblog.FollowService(bodyFollowErr)
			Expect(errMessage).NotTo(BeNil())
			Expect(sendFollow).To(BeNil())
		})

		It("Error al seguir usuario. Error en el servidor", func() {
			expectedErrMessage := models.ErrorResponse(" | Error | ", "Ocurrio un error al seguir un usuario", http.StatusInternalServerError, nil)
			mockSql.
				EXPECT().
				FollowRepository(&bodyFollowErr).
				Times(1).
				Return(nil, &expectedErrMessage)

			sendFollow, errMessage := serviceMicroblog.FollowService(bodyFollowErr)
			Expect(errMessage).NotTo(BeNil())
			Expect(sendFollow).To(BeNil())
		})
	})

	Describe("RECIBIR MENSAJES", func() {
		It("Recibo mensajes con exito", func() {
			mockSql.
				EXPECT().
				TimelineRepository(&viewMessagesOK).
				Times(1).
				Return(feed, nil)

			getFeed, errMessage := serviceMicroblog.TimelineService(viewMessagesOK)
			Expect(errMessage).To(BeNil())
			Expect(getFeed).NotTo(BeNil())
		})

		It("Error al recibir mensajes. El usuario no existe", func() {
			expectedErrMessage := models.ErrorResponse("| Error | ", "No se encontró el usuario", http.StatusNotFound, nil)
			mockSql.
				EXPECT().
				TimelineRepository(&viewMessagesErr).
				Times(1).
				Return(nil, &expectedErrMessage)

			getFeed, errMessage := serviceMicroblog.TimelineService(viewMessagesErr)
			Expect(errMessage).NotTo(BeNil())
			Expect(getFeed).To(BeNil())
		})

		It("Error al recibir mensajes. Error en el servidor", func() {
			expectedErrMessage := models.ErrorResponse(" | Error | ", "Ocurrio un error al obtener mensajes", http.StatusInternalServerError, nil)
			mockSql.
				EXPECT().
				TimelineRepository(&viewMessagesErr).
				Times(1).
				Return(nil, &expectedErrMessage)

			getFeed, errMessage := serviceMicroblog.TimelineService(viewMessagesErr)
			Expect(errMessage).NotTo(BeNil())
			Expect(getFeed).To(BeNil())
		})
	})
})
