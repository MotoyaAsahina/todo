package model

import (
	"context"
	"encoding/base64"
	"time"

	"gorm.io/gorm"
)

type Token struct {
	TokenID string `json:"tokenID" gorm:"size:300;primary_key"`
}

func IssueToken(ctx context.Context, googleToken string) (*Token, error) {
	token := &Token{
		TokenID: base64.RawURLEncoding.EncodeToString([]byte(time.Now().String() + googleToken)),
	}
	if err := GetDB(ctx).Create(token).Error; err != nil {
		return nil, err
	}
	return token, nil
}

func CertificateToken(ctx context.Context, token string) (bool, error) {
	err := GetDB(ctx).First(&Token{TokenID: token}).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
