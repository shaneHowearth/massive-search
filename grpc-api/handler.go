package api

import (
	"context"
	"log"
	"strings"
)

// Server -
type Server struct {
	WordList map[string]int32
}

// NewServer -
func NewServer() Server {
	server := Server{}
	server.WordList = make(map[string]int32)

	server.WordList["hello"] = 0
	server.WordList["goodbye"] = 0
	server.WordList["simple"] = 0
	server.WordList["list"] = 0
	server.WordList["search"] = 0
	server.WordList["filter"] = 0
	server.WordList["yes"] = 0
	server.WordList["no"] = 0

	return server
}

// GetWord -
func (s *Server) GetWord(ctx context.Context, in *Words) (*Words, error) {

	word := strings.ToLower(in.Word[0])
	log.Printf("Searched for %s", word)
	if _, ok := s.WordList[word]; ok {
		// update the search count
		s.WordList[word] += int32(1)
		return &Words{Count: s.WordList[word], Word: []string{word}}, nil
	}
	log.Printf("Searched for %s", word)
	// Word not found
	return &Words{Count: int32(-1), Word: []string{word}}, nil
}
