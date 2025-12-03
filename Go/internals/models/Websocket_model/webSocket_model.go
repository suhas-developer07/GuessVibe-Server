package wsModel

type ClientMessage struct {
    Type      string `json:"type"`      
    UserID    string `json:"userId,omitempty"`
    Answer    string `json:"answer,omitempty"` 
    SessionID string `json:"sessionId,omitempty"`
}

type ServerMessage struct {
    Type          string      `json:"type"`  
    SessionID     string      `json:"sessionId,omitempty"`
    Question      string      `json:"question,omitempty"`
    QuestionNumber int        `json:"questionNumber,omitempty"`
    Guess         string      `json:"guess,omitempty"`
    Confidence    float32     `json:"confidence,omitempty"`
    Alternatives  []string    `json:"alternatives,omitempty"`
    ErrorMessage  string      `json:"error,omitempty"`
}
