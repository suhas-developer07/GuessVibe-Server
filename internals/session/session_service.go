package session

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	//"github.com/suhas-developer07/GuessVibe-Server/internals/RedisSession_model"
)

type SessionState struct {
    SessionID       string             `json:"sessionId"`
    UserID          string             `json:"userId"`
    QuestionCount   int                `json:"questionCount"`
    History         []QA               `json:"history"`
    CandidateScores map[string]float64 `json:"candidateScores"`
    State           string             `json:"state"`  
}


type QA struct {
    Question string `json:"question"`
    Answer   string `json:"answer"`
}

type SessionRepository interface {
	SaveSession(ctx context.Context, state *SessionState) error
	GetSession(ctx context.Context, sessionID string) (*SessionState, error)
	UpdateSession(ctx context.Context, state *SessionState) error
	DeleteSession(ctx context.Context, sessionID string) error
}

type Service struct {
	Repo SessionRepository
}

func NewService(repo SessionRepository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) CreateSession(ctx context.Context, userID string) (*SessionState, error) {

	sessionID := uuid.New().String()

	state := &SessionState{
		SessionID:       sessionID,
		UserID:          userID,
		QuestionCount:   0,
		History:         []QA{},
		CandidateScores: make(map[string]float64),
		State:           "ASKING",
	}

	err := s.Repo.SaveSession(ctx, state)
	if err != nil {
		return nil, fmt.Errorf("failed to save session: %w", err)
	}

	return state, nil
}
