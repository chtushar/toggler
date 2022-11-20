package users

import (
	"go.uber.org/zap"
)

type Handler struct {
	log        *zap.Logger
	repository *Repository
}

func NewHandler(log *zap.Logger, repository *Repository) *Handler {
	return &Handler{
		log:        log,
		repository: repository,
	}
}
