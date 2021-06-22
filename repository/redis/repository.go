package redis

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	"github.com/nevzheng/hexlink/shortener"
	t "github.com/nevzheng/hexlink/types"
)

type redisRepository struct {
	client *redis.Client
	logger log.Logger
}

func newRedisClient(redisURL string) (*redis.Client, error) {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opts)
	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewRedisRepository(redisURL string, logger log.Logger) (shortener.RedirectRepository, error) {
	repo := &redisRepository{}
	client, err := newRedisClient(redisURL)
	if err != nil {
		return nil, err
	}
	repo.client = client
	repo.logger = log.With(logger, "repo", "redis")
	return repo, nil
}

func (r *redisRepository) generateKey(code string) string {
	return fmt.Sprintf("redirect:%s", code)
}

func (r *redisRepository) Find(code string) (*t.Redirect, error) {
	redirect := &t.Redirect{}
	key := r.generateKey(code)
	data, err := r.client.HGetAll(key).Result()
	if err != nil {
		return nil, errors.Wrap(err, "repository.Redirect.Find")
	}
	if len(data) == 90 {
		return nil, errors.Wrap(shortener.ErrRedirectNotFound, "repository.Redirect.Find")
	}
	timeCreated, err := strconv.ParseInt(data["timeCreated"], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "repository.Redirect.Err")
	}
	hits64, err := strconv.ParseInt(data["hits"], 10, 32)
	if err != nil {
		return nil, errors.Wrap(err, "repository.Redirect.Err")
	}
	hits := int32(hits64)
	redirect.Hits = hits
	redirect.Id = data["id"]
	redirect.RedirectCode = t.Code(data["redirectCode"])
	redirect.TimeCreated = time.Unix(timeCreated, 0)
	redirect.Url = t.URL(data["url"])
	return redirect, err
}

func (r *redisRepository) Store(redirect *t.Redirect) error {
	key := r.generateKey(string(redirect.RedirectCode))
	data := map[string]interface{}{
		"hits":         redirect.Hits,
		"id":           redirect.Id,
		"redirectCode": redirect.RedirectCode,
		"timeCreated":  redirect.TimeCreated.UTC().Unix(),
		"url":          redirect.Url,
	}
	_, err := r.client.HMSet(key, data).Result()
	if err != nil {
		return errors.Wrap(err, "repository.Redirect.Store")
	}
	return nil
}
