package player_postgres_test

import (
	"context"
	"game_server/domain"
	"game_server/player"
	"game_server/player/player_repository/player_postgres"
	"time"

	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestFetch(t *testing.T) {
	//opening the database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	//initial rows mock value
	mockUsers := []domain.Player{
		domain.Player{
			ID: 1, Email: "email 1", Password: "password 1", Username: "aniis", CreatedAt: time.Now(),
		},
		domain.Player{
			ID: 2, Email: "email 2", Password: "password 2", Username: "aeeem", CreatedAt: time.Now(),
		},
	}

	//adding rows into for mocking purpose
	rows := sqlmock.NewRows([]string{"id", "email", "password", "username", "updated_at", "created_at"}).
		AddRow(mockUsers[0].ID, mockUsers[0].Email, mockUsers[0].Password, mockUsers[0].Username,
			mockUsers[0].UpdatedAt, mockUsers[0].CreatedAt).
		AddRow(mockUsers[1].ID, mockUsers[1].Email, mockUsers[1].Password, mockUsers[1].Username,
			mockUsers[1].UpdatedAt, mockUsers[1].CreatedAt)

	query := "SELECT id,email,password,username , updated_at, created_at FROM users WHERE created_at > \\? ORDER BY created_at LIMIT \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)

	a := player_postgres.NewPlayerRepository(sqlxDB)
	cursor := player.EncodeCursor(mockUsers[1].CreatedAt)
	num := int64(2)
	list, nextCursor, err := a.Fetch(context.TODO(), cursor, uint64(num))
	assert.NotEmpty(t, nextCursor)
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	now := time.Now()
	pl := &domain.Player{
		Username:  "hehe",
		Password:  "hahaha",
		Email:     "oakwekaowekaew",
		UpdatedAt: now,
		CreatedAt: now,
	}
	query := `INSERT INTO users`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(pl.Username, pl.Password, pl.Email).WillReturnResult(sqlmock.NewResult(1, 1))
	a := player_postgres.NewPlayerRepository(sqlxDB)
	err = a.Store(context.TODO(), pl)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), pl.ID)

}

func TestGetByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	Rows := sqlmock.NewRows([]string{
		"id",
		"username",
		"email",
		"password",
		"created_at",
		"updated_at",
	}).AddRow(1, "aem", "arifmaulanaa@gmail.com", "asdasd", time.Now(), time.Now())
	query := "SELECT id,username,email,password,created_at,updated_at from users where username=\\? limit \\?"
	mock.ExpectQuery(query).WillReturnRows(Rows)
	p := player_postgres.NewPlayerRepository(sqlxDB)
	player, err := p.GetByUsername(context.TODO(), "aem")
	assert.NoError(t, err)
	assert.NotNil(t, player)
}
func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	createdAt := time.Now()
	updated_at := createdAt
	Rows := sqlmock.NewRows([]string{
		"id",
		"username",
		"email",
		"password",
		"created_at",
		"updated_at",
	}).AddRow(1, "aem", "arifmaulanaa@gmail.com", "asdasd", createdAt, updated_at)
	query := "SELECT id,username,email,password,created_at,updated_at FROM users WHERE id=\\?"
	mock.ExpectQuery(query).WillReturnRows(Rows)
	p := player_postgres.NewPlayerRepository(sqlxDB)
	player, err := p.GetByID(context.TODO(), 1)
	assert.NoError(t, err)
	assert.Equal(t, player, domain.Player{
		ID:        1,
		Username:  "aem",
		Email:     "arifmaulanaa@gmail.com",
		Password:  "asdasd",
		CreatedAt: createdAt,
		UpdatedAt: updated_at,
	})
	assert.NotNil(t, player)
}
