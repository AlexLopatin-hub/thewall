package core_conn

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Connection interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Close()

	OpTimeout() time.Duration
}

type Conn struct {
	*pgx.Conn
	opTimeout time.Duration
}

func NewConn(
	ctx context.Context,
	config Config,
) (*Conn, error) {
	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	conn, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		return &Conn{}, fmt.Errorf("create connection: %w", err)
	}

	if err = conn.Ping(ctx); err != nil {
		return &Conn{}, fmt.Errorf("ping connection: %w", err)
	}

	return &Conn{
		Conn:      conn,
		opTimeout: config.Timeout,
	}, nil
}

func (c *Conn) OpTimeout() time.Duration {
	return c.opTimeout
}
