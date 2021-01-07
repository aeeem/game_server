package domain

import (
	"context"
	"time"
)

type Player struct {
	ID        int64     `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"password" db:"password"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type PlayerRepository interface {
	Fetch(ctx context.Context, cursor string, num uint64) (pl []Player, nextCursor string, err error)
	GetByID(ctx context.Context, id uint64) (pl Player, err error)
	Store(ctx context.Context, play *Player) (err error)
	GetByUsername(ctx context.Context, username string) (pl Player, err error)
	// SetPosition(ctx context.Context, XPos float64, YPos float64, ZPos float64)
	// GetPosition(ctx context.Context, id uint64)
}

type PlayerUsecase interface {
	Login(ctx context.Context, username string, password string) (pl Player, err error)
}
