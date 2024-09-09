package authentication

import (
	"net/http"

	gojwt "github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/topgate/gcim-temporary/back/app/internal/authentication"
	jwt "github.com/topgate/gcim-temporary/back/app/internal/authentication/jwtimpl"
	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	"github.com/topgate/gcim-temporary/back/app/internal/repositories"
	"golang.org/x/exp/slices"
	"golang.org/x/xerrors"
)

// Authenticator - Authentication する middleware
type Authenticator struct{}

// NewAuthenticatorParam - NewAuthenticator のパラメータ
type NewAuthenticatorParam struct {
	JWTSecret             []byte                                            // SigningKey
	CookieSessionName     string                                            // CookieSessionName
	SkipPaths             []string                                          // 認証を必要としないパス
	ContextKey            string                                            // ContextKey
	AuthenticationService authentication.Provider                           // 認証サービス
	UserSessionRepository repositories.BaseRepository[entities.UserSession] // セッションのリポジトリ
}

// NewAuthenticator - middleware.JWTConfig constructor
func NewAuthenticator(param NewAuthenticatorParam) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    param.JWTSecret,
		SigningMethod: gojwt.SigningMethodHS512.Name,
		TokenLookup:   "cookie:" + param.CookieSessionName,
		Skipper: func(c echo.Context) bool {
			return slices.Contains(param.SkipPaths, c.Request().URL.Path)
		},
		ContextKey:     param.ContextKey,
		ParseTokenFunc: parseTokenFunc(param.AuthenticationService, param.UserSessionRepository),
		ErrorHandler:   handleError,
		NewClaimsFunc:  func(echo.Context) gojwt.Claims { return new(jwt.Claims) },
	})
}

func tokenToClaims(provider authentication.Provider, authToken string) (*jwt.Claims, error) {
	claimsInterface, err := provider.VerifyAuthToken(authToken)
	if err != nil {
		return nil, xerrors.Errorf("failed to verify auth token: %w", err)
	}

	return claimsInterface.(*jwt.Claims), nil
}

func parseTokenFunc(provider authentication.Provider, _ repositories.BaseRepository[entities.UserSession]) func(c echo.Context, authToken string) (any, error) {
	return func(c echo.Context, authToken string) (any, error) {
		req := c.Request()
		ctx := req.Context()

		// トークンをClaims構造体に変換する
		claims, err := tokenToClaims(provider, authToken)
		if err != nil {
			return nil, xerrors.Errorf("failed to verify auth token: %w", err)
		}

		/* FIXME
		ctx = superctx.WithValue[ctxkey.UserID](ctx, "xxxxxxxxxxxxxxxxxxxx") // UserID
		*/
		c.SetRequest(req.WithContext(ctx))

		/* FIXME
		// セッションを取得する
		session, err := sessionRepository.GetByID(ctx, "yyyyyyyyyyyyyyyyyyyy") // SessionID
		if err != nil {
			return nil, xerrors.Errorf("session not exist: %w", err)
		}

		// セッションの有効期限のチェック
		expiresAt := synchro.In[tz.AsiaTokyo](session.ExpiresAt())
		// NOTE: now >= expiresAt
		if !synchro.Now[tz.AsiaTokyo]().Before(expiresAt) {
			return nil, xerrors.Errorf("expiresAt=%s(sessionID=%s): %w", expiresAt.String(), "yyyyyyyyyyyyyyyyyyyy", "session expired")
		}
		*/

		return claims, nil
	}
}

func handleError(c echo.Context, err error) error {
	herr := echo.NewHTTPError(http.StatusUnauthorized, "Missing JWT")
	return c.JSON(herr.Code, herr)
}
