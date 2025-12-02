package grpcclient

import (
	"context"
	"log"
	"time"

	pb "github.com/suhas-developer07/GuessVibe-Server/generated/proto"
	"google.golang.org/grpc"
)

type LLMClient struct {
	Client pb.LLMServiceClient
	conn   *grpc.ClientConn
}

func NewLLMClient(grpcURL string) *LLMClient {
	conn, err := grpc.Dial(grpcURL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to LLM gRPC server: %v", err)
	}

	return &LLMClient{
		Client: pb.NewLLMServiceClient(conn),
		conn:   conn,
	}
}

// ---------------------------
// Helper: request context
// ---------------------------
func (c *LLMClient) ctx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

// ---------------------------
// 1. Generate First Question
// ---------------------------
func (c *LLMClient) GenerateFirstQuestion(state *pb.SessionState) (string, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	resp, err := c.Client.GenerateFirstQuestion(ctx, state)
	if err != nil {
		return "", err
	}

	return resp.Question, nil
}

// ---------------------------
// 2. Generate Next Question
// ---------------------------
func (c *LLMClient) GenerateNextQuestion(state *pb.SessionState) (string, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	resp, err := c.Client.GenerateNextQuestion(ctx, state)
	if err != nil {
		return "", err
	}

	return resp.Question, nil
}

// ---------------------------
// 3. Generate Final Guess
// ---------------------------
func (c *LLMClient) GenerateFinalGuess(state *pb.SessionState) (*pb.LLMGuessResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.Client.GenerateFinalGuess(ctx, state)
}
