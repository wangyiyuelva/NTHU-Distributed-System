package dao

import (
	"context"
	"time"

	"github.com/NTHU-LSALAB/NTHU-Distributed-System/pkg/rediskit"
	"github.com/go-redis/cache/v8"
	"github.com/google/uuid"
)

type redisCommentDAO struct {
	cache   *cache.Cache
	baseDAO CommentDAO
}

var _ CommentDAO = (*redisCommentDAO)(nil)

const (
	commentDAOLocalCacheSize     = 1024
	commentDAOLocalCacheDuration = 1 * time.Minute
	commentDAORedisCacheDuration = 3 * time.Minute
)

// create redisCommentDAO by cache and baseDAO
func NewRedisCommentDAO(client *rediskit.RedisClient, baseDAO CommentDAO) *redisCommentDAO {
	// Redis TODO
}

// List all the comments related to the videoID
// Notice that all the comments will be stored as a single value in the cache
// The key is generated by listCommentKey function in comment.go
// The implementation should handle both cache miss and cache hit scenarios
func (dao *redisCommentDAO) ListByVideoID(ctx context.Context, videoID string, limit, offset int) ([]*Comment, error) {
	// Redis TODO
}

// The operation are not cacheable, just pass down to baseDAO
func (dao *redisCommentDAO) Create(ctx context.Context, comment *Comment) (uuid.UUID, error) {
	// Redis TODO
}

// The following operations are not cacheable, just pass down to baseDAO
func (dao *redisCommentDAO) Update(ctx context.Context, comment *Comment) error {
	// Redis TODO
}

// The following operations are not cacheable, just pass down to baseDAO
func (dao *redisCommentDAO) Delete(ctx context.Context, id uuid.UUID) error {
	// Redis TODO
}

// The following operations are not cacheable, just pass down to baseDAO
func (dao *redisCommentDAO) DeleteByVideoID(ctx context.Context, videoID string) error {
	// Redis TODO
}
