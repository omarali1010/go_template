package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Client wraps the pgxpool.Pool so you can extend it if needed
type Client struct {
	Pool *pgxpool.Pool
}

// NewClient creates a new PostgreSQL client connection pool
func NewClient(host, port, user, password, dbname string, maxConns int) (*Client, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, dbname)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse postgres config: %w", err)
	}

	config.MaxConns = int32(maxConns)
	// Optional: configure other settings like connection timeout here

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	// Ping the database to verify connection
	err = pool.Ping(ctx)
	if err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	log.Println("Connected to PostgreSQL successfully")

	return &Client{Pool: pool}, nil
}

// Close closes the connection pool
func (c *Client) Close() {
	if c.Pool != nil {
		c.Pool.Close()
	}
}
