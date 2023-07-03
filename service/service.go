package service

import (
	"fmt"

	"github.com/fishkaoff/monitor-system-database-microservice/messages"
	"github.com/fishkaoff/monitor-system-database-microservice/storage"
)

type Service interface {
	SaveSite(chatID int64, site string) string
	GetSites(chatID int64) []string
	DeleteSite(chatID int64, site string) string
	SaveUser(chatID int64, token string) string
}

type service struct {
	storage storage.Storage
}

func New(s storage.Storage) *service {
	return &service{storage: s}
}

func (s *service) SaveSite(chatID int64, site string) string {
	err := s.storage.Save(chatID, site)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return messages.URLNOTADDED
	}

	return messages.URLADDED
}

func (s *service) GetSites(chatID int64) []string {
	sites, err := s.storage.Get(chatID)
	if err != nil {
		return []string{}
	}

	return sites
}

func (s *service) DeleteSite(chatID int64, site string) string {
	err := s.storage.Delete(chatID, site)
	if err != nil {
		return messages.URLNOTDELETED
	}

	return messages.URLDELETED
}

func (s *service) SaveUser(chatID int64, token string) string {
	err := s.storage.SaveUser(chatID, token)
	if err != nil {
		return err.Error()
	}
	return token
}