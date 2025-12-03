package session

import (
	pb "github.com/suhas-developer07/GuessVibe-Server/generated/proto"
	//models "github.com/suhas-developer07/GuessVibe-Server/internals/models/RedisSession_model"
)

func (s *SessionState) ToProto() *pb.SessionState {
	history := []*pb.QA{}
	for _, h := range s.History {
		history = append(history, &pb.QA{
			Question: h.Question,
			Answer:   h.Answer,
		})
	}

	// Convert candidateScores (float64 → float32)
	candidate := map[string]float32{}
	for k, v := range s.CandidateScores {
		candidate[k] = float32(v)
	}

	return &pb.SessionState{
		SessionId:       s.SessionID,
		QuestionCount:   int32(s.QuestionCount),
		History:         history,
		CandidateScores: candidate,
	}
}

