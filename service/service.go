package service

import (
	"microblog/models"
	"microblog/repository"
)

type Service interface {
	SendMessageService(bodyMessage models.Message) (*models.Message, *models.ErrorMessage)
	FollowService(bodyFollowers models.UsernameFollower) (*models.Follower, *models.ErrorMessage)
	TimelineService(bodyTimeline models.Timeline) ([]models.Feed, *models.ErrorMessage)
}

type service struct {
	sqlRepository repository.SQLRepository
}

func NewService(sqlRepository repository.SQLRepository) Service {
	return service{
		sqlRepository: sqlRepository,
	}
}

func (s service) SendMessageService(bodyMessage models.Message) (*models.Message, *models.ErrorMessage) {
	body, errResponse := s.sqlRepository.SendMessageRepository(&bodyMessage)
	if errResponse != nil {
		return nil, errResponse
	}
	return body, nil
}

func (s service) FollowService(bodyFollowers models.UsernameFollower) (*models.Follower, *models.ErrorMessage) {
	followUser, errResponse := s.sqlRepository.FollowRepository(&bodyFollowers)
	if errResponse != nil {
		return nil, errResponse
	}
	return followUser, nil
}

func (s service) TimelineService(bodyTimeline models.Timeline) ([]models.Feed, *models.ErrorMessage) {
	messageFollows, err := s.sqlRepository.TimelineRepository(&bodyTimeline)
	if err != nil {
		return nil, err
	}
	return messageFollows, nil
}
