package telegram

import (
	"sync"
	"time"
)

type UserRateLimiter struct {
	mu       sync.Mutex
	requests map[int]int       // Количество запросов пользователя
	lastSeen map[int]time.Time // Время последнего запроса
	limit    int               // Максимальное количество запросов
	interval time.Duration     // Интервал времени для сброса счетчика
}

func NewUserRateLimiter(limit int, interval time.Duration) *UserRateLimiter {
	return &UserRateLimiter{
		requests: make(map[int]int),
		lastSeen: make(map[int]time.Time),
		limit:    limit,
		interval: interval,
	}
}

func (rl *UserRateLimiter) Allow(userID int) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()

	// Если прошло больше времени, чем интервал, сбрасываем счетчик
	if last, ok := rl.lastSeen[userID]; ok && now.Sub(last) > rl.interval {
		delete(rl.requests, userID)
		delete(rl.lastSeen, userID)
	}

	rl.requests[userID]++
	rl.lastSeen[userID] = now

	// Проверяем, не превышен ли лимит
	return rl.requests[userID] <= rl.limit
}
