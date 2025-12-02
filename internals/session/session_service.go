package session

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	models "github.com/suhas-developer07/GuessVibe-Server/internals/models/RedisSession_model"
)

type SessionRepository interface {
	SaveSession(ctx context.Context, state *models.SessionState) error
	GetSession(ctx context.Context, sessionID string) (*models.SessionState, error)
	UpdateSession(ctx context.Context, state *models.SessionState) error
	DeleteSession(ctx context.Context, sessionID string) error
}

type Service struct {
	Repo SessionRepository
}

func NewService(repo SessionRepository) *Service {
	return &Service{Repo: repo}
}

// CreateSession generates a new session with UUID and default values
func (s *Service) CreateSession(ctx context.Context, userID string) (*models.SessionState, error) {

	sessionID := uuid.New().String()

	state := &models.SessionState{
		SessionID:       sessionID,
		UserID:          userID,
		QuestionCount:   0,
		History:         []models.QA{},
		CandidateScores: make(map[string]float64),
		State:           "ASKING",
	}

	err := s.Repo.SaveSession(ctx, state)
	if err != nil {
		return nil, fmt.Errorf("failed to save session: %w", err)
	}

	return state, nil
}
