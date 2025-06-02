package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/joqd/go-russian/internal/core/domain"
	"github.com/joqd/go-russian/internal/core/port"
	"github.com/joqd/go-russian/pkg/redisdb"

	"github.com/redis/go-redis/v9"
)

type wordCache struct {
	rdb  *redis.Client
	xlog port.Logger
	ttl  time.Duration
}

func NewWordCache(redis redisdb.Redis, xlog port.Logger) port.WordCache {
	return &wordCache{
		rdb:  redis.Client,
		xlog: xlog,
		ttl:  redis.TTL,
	}
}

func (w *wordCache) Set(ctx context.Context, word *domain.Word) error {
	key := fmt.Sprintf("word:%s", word.ID)

	data, err := json.Marshal(word)
	if err != nil {
		w.xlog.Error("marshal word failed, word_id=%s, err=%v", word.ID, err)
		return err
	}

	if _, err := w.rdb.Set(ctx, key, data, w.ttl).Result(); err != nil {
		w.xlog.Error("redis set error, key=%s, err=%v", key, err)
		return err
	}

	return nil
}

func (w *wordCache) Get(ctx context.Context, id string) (*domain.Word, error) {
	key := fmt.Sprintf("word:%s", id)

	val, err := w.rdb.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, domain.ErrDataNotFound
		}

		w.xlog.Error("redis get error, key=%s, err=%v", key, err)
		return nil, err
	}

	var word domain.Word
	if err := json.Unmarshal([]byte(val), &word); err != nil {
		w.xlog.Error("unmarshal word failed, key=%s, er%v", key, err)
		return nil, err
	}

	return &word, nil
}

func (w *wordCache) Del(ctx context.Context, id string) error {
	key := fmt.Sprintf("word:%s", id)
	if err := w.rdb.Del(ctx, key).Err(); err != nil {
		w.xlog.Error("redis delete error, key=%s, err=%v", key, err)
		return err
	}
	return nil
}
