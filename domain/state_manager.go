package domain

import "context"

//StateManager is a abstract struct
type StateManager struct {
	Players []Player
	PosX    float64 `json:"pos_x"`
	PosY    float64 `json:"pos_y"`
	PosZ    float64 `json:"pos_z"`
}

type StateRepository interface {
	Store(ctx context.Context, play *StateManager) (err error)
	GetByID(ctx context.Context, cursor string, num uint64) (pl []Player, nextCursor string, err error)
	Update(ctx context.Context, state *StateManager) (err error)
}
