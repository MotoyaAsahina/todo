package model

import (
	"context"
)

type Token struct {
	TokenID string `json:"tokenID" gorm:"size:400;primary_key"`
}

type ITokenRepository interface {
	IssueToken(ctx context.Context, googleToken string) (*Token, error)
	CertificateToken(ctx context.Context, token string) (bool, error)
}
