package api

import (
	"context"
	"sort"
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
	if _, ok := s.WordList[word]; ok {
		// update the search count
		s.WordList[word] += int32(1)
		return &Words{Count: s.WordList[word], Word: []string{word}}, nil
	}
	// Word not found
	return &Words{Count: int32(-1), Word: []string{word}}, nil
}

// UpdateWords
func (s *Server) UpdateWords(ctx context.Context, in *Words) (*Words, error) {
	tmpWords := make(map[string]int)
	for _, word := range in.Word {
		word = strings.ToLower(word)
		tmpWords[word] = 0
		if _, ok := s.WordList[word]; !ok {
			s.WordList[word] = 0
		}
	}
	// Delete words not in update set
	for word, _ := range s.WordList {
		if _, ok := tmpWords[word]; !ok {
			delete(s.WordList, word)
		}
	}
	return in, nil
}

// TopFive -
func (s *Server) TopFive(ctx context.Context, e *Empty) (*WordSlice, error) {
	var ss []*Words

	for k, v := range s.WordList {
		ss = append(ss, &Words{v, []string{k}})
	}
	// Sort the slice by count, highest to lowest
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Count > ss[j].Count
	})
	var ws WordSlice
	if len(ss) != 0 {
		// Grab the first 5 from the sorted slice
		max := 5
		if len(ss) < max {
			max = len(ss) - 1
		}
		ws.Slice = ss[:max]
	}
	return &ws, nil

}
