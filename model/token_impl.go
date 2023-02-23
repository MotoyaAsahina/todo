package model

import (
	"context"
	"encoding/base64"
	"gorm.io/gorm"
	"time"
)

type TokenRepository struct {
	db *DB
}

func (repo *TokenRepository) IssueToken(ctx context.Context, googleToken string) (*Token, error) {
	token := &Token{
		TokenID: base64.RawURLEncoding.EncodeToString([]byte(time.Now().String() + googleToken)),
	}
	if err := repo.db.GetDB(ctx).Create(token).Error; err != nil {
		// FIXME: おそらく既に登録されていてエラー
		//  Docker 再起動すると新しいトークンが生成されるのはなぜ？
		return nil, err
	}
	return token, nil
}

func (repo *TokenRepository) CertificateToken(ctx context.Context, token string) (bool, error) {
	err := repo.db.GetDB(ctx).First(&Token{TokenID: token}).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
