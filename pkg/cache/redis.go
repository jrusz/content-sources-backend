package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/RedHatInsights/rbac-client-go"
	"github.com/content-services/content-sources-backend/pkg/config"
	"github.com/redhatinsights/platform-go-middlewares/identity"
	"github.com/redis/go-redis/v9"
)

type redisCache struct {
	client *redis.Client
}

func NewRedisCache() *redisCache {
	c := config.Get()
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisUrl(),
		Username: c.Clients.Redis.Username,
		Password: c.Clients.Redis.Password,
		DB:       c.Clients.Redis.DB,
	})
	return &redisCache{
		client: client,
	}
}

// authKey constructs the cache key for AccessList caching
func authKey(ctx context.Context) string {
	identity := identity.Get(ctx)
	return fmt.Sprintf("auth:%v,%v", identity.Identity.User.Username, identity.Identity.OrgID)
}

// GetAccessList uses the request context to read user information, and then tries to retrieve the role AccessList from the cache
func (c *redisCache) GetAccessList(ctx context.Context) (rbac.AccessList, error) {
	accessList := rbac.AccessList{}
	cmd := c.client.Get(ctx, authKey(ctx))
	if errors.Is(cmd.Err(), redis.Nil) {
		return nil, NotFound
	} else if cmd.Err() != nil {
		return nil, fmt.Errorf("redis error: %w", cmd.Err())
	}

	buf, err := cmd.Bytes()
	if err != nil {
		return nil, fmt.Errorf("redis bytes conversion error: %w", err)
	}

	err = json.Unmarshal(buf, &accessList)
	if err != nil {
		return nil, fmt.Errorf("redis unmarshal error: %w", err)
	}
	return accessList, nil
}

// SetAccessList uses the request context to read user information, and loads the cache with role access list
func (c *redisCache) SetAccessList(ctx context.Context, accessList rbac.AccessList) error {
	buf, err := json.Marshal(accessList)
	if err != nil {
		return fmt.Errorf("unable to marshal for Redis cache: %w", err)
	}

	c.client.Set(ctx, authKey(ctx), string(buf), config.Get().Clients.Redis.Expiration)
	return nil
}
