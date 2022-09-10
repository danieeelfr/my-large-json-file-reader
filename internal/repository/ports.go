package repository

import (
	"fmt"
	"time"

	"github.com/danieeelfr/my-large-json-file-reader/internal/config"
	"github.com/go-redis/redis"
)

// PortsRepo is the repository interface
type PortsRepo interface {
	Get(lek string, pageSize int) (map[string]interface{}, error)
	Set(map[string]interface{}) error
}

// Repository holds the database configs and clients
type Repository struct {
	redisClient *redis.Client
	redisConfig *config.RedisConfig
}

// NewRepository Return a implementation of PortsRepo
func NewRepository(cfg *config.Config) PortsRepo {
	return &Repository{
		redisClient: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
			Password: cfg.Redis.Password,
			DB:       0, // using the default database
		}),
		redisConfig: cfg.Redis,
	}
}

// Get get ports with pagination
func (r *Repository) Get(lek string, pageSize int) (map[string]interface{}, error) {
	// TODO get with pagination
	return nil, nil
}

// Set save or update the port values according the map passed as parameter
func (r *Repository) Set(ports map[string]interface{}) error {
	start := time.Now()

	errs := 0
	for k, v := range ports {
		s := r.redisClient.Set(k, fmt.Sprintf("%v", v), 0) // wont expire

		r, e := s.Result()
		if e != nil {
			fmt.Printf("error when saving=%v \n result=%v \n", e, r)
			errs++
			// TODO: improve this error handling
			// I choice to do not stop the process here, but if the error is a connection error
			// Is better to stop.
		}
	}

	elapsed := time.Since(start)

	fmt.Printf("The saving on database process took [%v] and saved [%d] of [%d] records\n", elapsed, len(ports)-errs, len(ports))

	return nil
}
