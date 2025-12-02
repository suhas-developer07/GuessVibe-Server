package session

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
	models "github.com/suhas-developer07/GuessVibe-Server/internals/models/RedisSession_model"
)

type RedisRepo struct {
	Client *redis.Client
}

func NewRedisRepo(client *redis.Client) *RedisRepo {
	return &RedisRepo{Client: client}
}

func (r *RedisRepo) SaveSession(ctx context.Context, state *models.SessionState) error {
	data, _ := json.Marshal(state)
	return r.Client.Set(ctx, "session:"+state.SessionID, data, 0).Err()
}

func (r *RedisRepo) GetSession(ctx context.Context, sessionID string) (*models.SessionState, error) {
	data, err := r.Client.Get(ctx, "session:"+sessionID).Bytes()
	if err != nil {
		return nil, fmt.Errorf("session not found")
	}

	var state models.SessionState
	json.Unmarshal(data, &state)
	return &state, nil
}

func (r *RedisRepo) UpdateSession(ctx context.Context, state *models.SessionState) error {
	data, _ := json.Marshal(state)
	return r.Client.Set(ctx, "session:"+state.SessionID, data, 0).Err()
}

func (r *RedisRepo) DeleteSession(ctx context.Context, sessionID string) error {
	return r.Client.Del(ctx, "session:"+sessionID).Err()
}
