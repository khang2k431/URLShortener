package service

import (
    "math/rand"
    "sync"
    "time"
)

type URLService struct {
    store      map[string]string
    visitCount map[string]int
    mu         sync.Mutex
}

func NewURLService() *URLService {
    rand.Seed(time.Now().UnixNano())
    return &URLService{
        store:      make(map[string]string),
        visitCount: make(map[string]int),
    }
}

func (s *URLService) GenerateShortURL() string {
    const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    b := make([]byte, 6)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func (s *URLService) SaveURL(short, long string) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.store[short] = long
    s.visitCount[short] = 0
}

func (s *URLService) GetURL(short string) (string, bool) {
    s.mu.Lock()
    defer s.mu.Unlock()
    url, exists := s.store[short]
    if exists {
        s.visitCount[short]++
    }
    return url, exists
}

func (s *URLService) Stats() map[string]int {
    s.mu.Lock()
    defer s.mu.Unlock()
    return s.visitCount
}
