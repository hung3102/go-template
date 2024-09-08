package entities

import "time"

// Meta - エンティティのメタ情報
type Meta struct {
	createdAt time.Time  // 作成日時
	createdBy string     // 作成者
	updatedAt time.Time  // 更新日時
	updatedBy string     // 更新者
	deletedAt *time.Time // 削除日時
	deletedBy *string    // 削除者
}

// NewMetaParam - NewMeta のパラメータ
type NewMetaParam struct {
	CreatedAt time.Time  // 作成日時
	CreatedBy string     // 作成者
	UpdatedAt time.Time  // 更新日時
	UpdatedBy string     // 更新者
	DeletedAt *time.Time // 削除日時
	DeletedBy *string    // 削除者
}

// Update - 更新
func (m Meta) Update(now time.Time, userID string) *Meta {
	return &Meta{
		createdAt: m.createdAt,
		createdBy: m.createdBy,
		updatedAt: now,
		updatedBy: userID,
		deletedAt: m.deletedAt,
		deletedBy: m.deletedBy,
	}
}

// NewMeta - Meta のコンストラクタ
func NewMeta(param NewMetaParam) *Meta {
	return &Meta{
		createdAt: param.CreatedAt,
		createdBy: param.CreatedBy,
		updatedAt: param.UpdatedAt,
		updatedBy: param.UpdatedBy,
		deletedAt: param.DeletedAt,
		deletedBy: param.DeletedBy,
	}
}
