package shortener

import (
	"context"
	"fmt"
	"sync"

	pb "github.com/musadhiek-p/url-shortener/api"
)

// temporary in memmory storage
var (
	urlStorage = make(map[string]string)
	mu         sync.Mutex
)

// shortener serviceimplement the grpc service

type shortenerService struct {
	pb.UnimplementedShortenerServiceServer
}

func (s *shortenerService) ShortenURL(ctx context.Context, req *pb.ShortenRequest) (*pb.ShortenResponse, error) {
	mu.Lock()
	defer mu.Unlock()
	shortURL := fmt.Sprintf("short.ly%d", len(urlStorage)+1)

	urlStorage[shortURL] = req.LongUrl

	return &pb.ShortenResponse{ShortUrl: shortURL}, nil
}

func (s *shortenerService) ResolveURL(ctx context.Context, req *pb.ResolveRequest) (*pb.ResolveResponse, error) {
	mu.Lock()
	defer mu.Unlock()
	longURL, exists := urlStorage[req.ShortUrl]
	if !exists {
		return nil, fmt.Errorf("URL not found")
	}

	return &pb.ResolveResponse{LongURL: longURL}, nil
}
