package uuid

import (
	_uuid "github.com/google/uuid"
	"golang.org/x/xerrors"
)

// UUID - UUIDを作成
type UUID struct{}

// GetUUID - UUIDを作成する
func (u *UUID) GetUUID() (string, error) {
	uuid, err := _uuid.NewUUID()
	if err != nil {
		return "", xerrors.Errorf("error in UUID %w", err)
	}
	return uuid.String(), nil
}
