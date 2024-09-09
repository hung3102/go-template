package jwtimpl

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/topgate/gcim-temporary/back/app/internal/authentication"
	"github.com/topgate/gcim-temporary/back/pkg/timeutil"
	"golang.org/x/xerrors"
)

type jwtProvider struct {
	secret        []byte
	signingMethod *jwt.SigningMethodHMAC
	issuer        string
}

// NewJWTProviderParam - NewJWTProvider のパラメータ
type NewJWTProviderParam struct {
	Secret        []byte
	SigningMethod *jwt.SigningMethodHMAC
	Issuer        string
}

// NewJWTProvider - jwtProvider のコンストラクタ
func NewJWTProvider(param NewJWTProviderParam) authentication.Provider {
	return &jwtProvider{secret: param.Secret, signingMethod: param.SigningMethod, issuer: param.Issuer}
}

// issueAuthTokenParam - issueAuthToken のパラメータ
type issueAuthTokenParam struct {
	UserID    string    // ユーザID
	SessionID string    // セッションID
	Now       time.Time // 現在時刻
	IsBot     bool      // ボットかどうか
}

// issueAuthToken - AuthToken を発行する
func (j *jwtProvider) issueAuthToken(param issueAuthTokenParam) (string, error) {
	claims := Claims{
		UserID:    param.UserID,
		SessionID: param.SessionID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(timeutil.EndOfDay(param.Now)),
			IssuedAt:  jwt.NewNumericDate(param.Now),
			NotBefore: jwt.NewNumericDate(param.Now),
			Issuer:    j.issuer,
		},
	}

	token := jwt.NewWithClaims(j.signingMethod, claims)
	ss, err := token.SignedString(j.secret)
	if err != nil {
		return "", xerrors.Errorf("failed to signing: %w", err)
	}

	return ss, nil
}

// IssueAuthToken - AuthToken を発行する
// revive:disable:confusing-naming
func (j *jwtProvider) IssueAuthToken(userID, sessionID string, now time.Time) (string, error) {
	param := issueAuthTokenParam{
		UserID:    userID,
		SessionID: sessionID,
		Now:       now,
		IsBot:     false,
	}

	return j.issueAuthToken(param)
}

// VerifyAuthToken - AuthToken の検証を行う
func (j *jwtProvider) VerifyAuthToken(tokenString string) (claims jwt.Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, new(Claims), func(*jwt.Token) (any, error) {
		return j.secret, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, xerrors.Errorf("failed to parse jwt: %w", authentication.ErrAuthTokenExpired)
		}
		return nil, xerrors.Errorf("failed to parse jwt: %w", err)
	}

	if !token.Valid {
		return nil, xerrors.New("invalid jwt")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, xerrors.New("failed to cast custom claims")
	}

	return claims, nil
}
