package domain

import "context"

type Sessions struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UUID         string `json:"uuid"`
}

type SessionRepository interface {
	GetByID(ctx context.Context, UUID string) (sessions *Sessions, err error) //for validating
	Delete(ctx context.Context, UUID string) (err error)                      //for revoking token
	Store(ctx context.Context, ses *Sessions) (err error)                     //for storing new tokens
}
