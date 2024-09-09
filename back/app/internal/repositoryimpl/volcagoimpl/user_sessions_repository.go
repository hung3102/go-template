package volcagoimpl

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	"github.com/topgate/gcim-temporary/back/app/internal/repositories"
	"github.com/topgate/gcim-temporary/back/app/internal/repositoryerrors"
	"github.com/topgate/gcim-temporary/back/app/internal/volcago"
	"github.com/topgate/gcim-temporary/back/app/internal/volcago/infrastructures"
)

type userSessionImpl struct {
	infra infrastructures.UserSessionRepository
}

// NewUserSession - UserSession リポジトリを生成
func NewUserSession(client *firestore.Client) repositories.BaseRepository[entities.UserSession] {
	return &userSessionImpl{infra: infrastructures.NewUserSessionRepository(client)}
}

// Create - イベントを作成
func (u *userSessionImpl) Create(ctx context.Context, entity *entities.UserSession) error {
	_, err := u.infra.Insert(ctx, &volcago.UserSession{
		ID:        entity.ID(),
		UserID:    entity.UserID(),
		ExpiresAt: entity.ExpiresAt(),
		Meta: volcago.Meta{
			CreatedAt: entity.Meta().CreatedAt(),
			CreatedBy: entity.Meta().CreatedBy(),
			UpdatedAt: entity.Meta().UpdatedAt(),
			UpdatedBy: entity.Meta().UpdatedBy(),
			DeletedAt: entity.Meta().DeletedAt(),
			DeletedBy: entity.Meta().DeletedBy(),
		},
	})
	if err != nil {
		return repositoryerrors.NewUnknownError("failed to create entity", err)
	}

	return nil
}

// GetByID - ID から取得
func (u *userSessionImpl) GetByID(ctx context.Context, id string) (*entities.UserSession, error) {
	entity, err := u.infra.Get(ctx, id)
	if err != nil {
		if errors.Is(err, infrastructures.ErrNotFound) {
			return nil, repositoryerrors.NewNotFoundError(err)
		}
		return nil, repositoryerrors.NewUnknownError("failed to get entity", err)
	}

	return entities.NewUserSession(&entities.NewUserSessionParam{
		ID:        entity.ID,
		UserID:    entity.UserID,
		ExpiresAt: entity.ExpiresAt,
		Meta: entities.NewMeta(&entities.NewMetaParam{
			CreatedAt: entity.CreatedAt,
			CreatedBy: entity.CreatedBy,
			UpdatedAt: entity.UpdatedAt,
			UpdatedBy: entity.UpdatedBy,
			DeletedAt: entity.DeletedAt,
			DeletedBy: entity.DeletedBy,
		}),
	}), nil
}
