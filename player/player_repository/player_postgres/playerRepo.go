package player_postgres

import (
	"context"
	"game_server/domain"
	"game_server/player"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type playerPostgresRepository struct {
	Conn *sqlx.DB
}

func NewPlayerRepository(db *sqlx.DB) (playerRepos domain.PlayerRepository) {
	playerRepos = &playerPostgresRepository{
		Conn: db,
	}
	return
}

//handling queries result from databases
func (pr *playerPostgresRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Player, err error) {
	rows, err := pr.Conn.QueryxContext(ctx, query, args...)
	if err != nil {
		logrus.Print(err)
		return
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	for rows.Next() {
		p := domain.Player{}
		err = rows.StructScan(&p)
		if err != nil {
			return
		}
		result = append(result, p)
	}
	return
}

//Get all user from databases
func (pr *playerPostgresRepository) Fetch(ctx context.Context, cursor string, num uint64) (pl []domain.Player, nextCursor string, err error) {
	query := `SELECT id,email,password,username , updated_at, created_at FROM User WHERE created_at > ? ORDER BY created_at LIMIT ?`
	decodeCursor, err := player.DecodeCursor(cursor)
	if err != nil {
		return
	}

	pl, err = pr.fetch(ctx, query, decodeCursor, num) //fetching queries into databases
	if err != nil {
		return
	}
	//if result length zero then its last pages,
	if len(pl) > 0 {
		nextCursor = player.EncodeCursor(pl[len(pl)-1].CreatedAt)
	}

	return
}

func (pr *playerPostgresRepository) Store(ctx context.Context, play *domain.Player) (err error) {
	query := `INSERT INTO users (username,password,email) VALUES ($1,$2,$3)`
	stmt, err := pr.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	res, err := stmt.ExecContext(ctx, play.Username, play.Password, play.Email)
	if err != nil {
		logrus.Error(err)
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	play.ID = lastID
	return
}

func (pr *playerPostgresRepository) GetByUsername(ctx context.Context, username string) (pl domain.Player, err error) {
	query := `SELECT id,username,email,password,created_at,updated_at from users where username=? limit ?`
	list, err := pr.fetch(ctx, query, username, 1)
	if len(list) > 0 {
		pl = list[0]
	} else {
		err = domain.ErrNotfound
		return
	}
	return
}

//not implemented yet there is no test for this things
// func (pr *playerPostgresRepository) GetByID(ctx context.Context, id uint64) (pl domain.Player, err error) {
// 	query := `select id,username,email,password,created_at,updated_at from user where id=$1 limit 1`
// 	list, err := pr.fetch(ctx, query, id)
// 	if len(list) > 0 {
// 		pl = list[0]
// 	} else {
// 		err = domain.ErrNotfound
// 		return
// 	}
// 	return
// }

// func (pr *playerPostgresRepository) SetPosition(ctx context.Context, XPos float64, YPos float64, ZPos float64) {

// }
// func (pr *playerPostgresRepository) GetPosition(ctx context.Context, id uint64) {

// }
