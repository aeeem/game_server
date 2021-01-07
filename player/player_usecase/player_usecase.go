package playerusecase

import (
	"context"
	"game_server/domain"
	"time"
)

type playerUsecase struct {
	playerRepo     domain.PlayerRepository
	ContextTimeout time.Duration
}

//NewPlayerUsecase make new player usecase objects
func NewPlayerUsecase(pr domain.PlayerRepository, timeout time.Duration) domain.PlayerUsecase {
	return &playerUsecase{
		playerRepo: pr,
	}
}

func (pU *playerUsecase) Login(ctx context.Context, username string, password string) (pl domain.Player, err error) {
	return
}

func GenerateToken(data map[string]string) (Token string, err error) {

	return
}
