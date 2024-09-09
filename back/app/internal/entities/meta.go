package entities

import "time"

// Meta - エンティティのメタ情報
type Meta struct {
	createdAt time.Time  // 作成日時
	createdBy string     // 作成者
	updatedAt time.Time  // 更新日時
	updatedBy string     // 更新者
	deletedAt *time.Time // 削除日時
	deletedBy string     // 削除者
}

// NewMetaParam - NewMeta のパラメータ
type NewMetaParam struct {
	CreatedAt time.Time  // 作成日時
	CreatedBy string     // 作成者
	UpdatedAt time.Time  // 更新日時
	UpdatedBy string     // 更新者
	DeletedAt *time.Time // 削除日時
	DeletedBy string     // 削除者
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
func NewMeta(param *NewMetaParam) *Meta {
	return &Meta{
		createdAt: param.CreatedAt,
		createdBy: param.CreatedBy,
		updatedAt: param.UpdatedAt,
		updatedBy: param.UpdatedBy,
		deletedAt: param.DeletedAt,
		deletedBy: param.DeletedBy,
	}
}

// CreatedAt - CreatedAt のゲッター
func (m *Meta) CreatedAt() time.Time {
	return m.createdAt
}

// CreatedBy - CreatedBy のゲッター
func (m *Meta) CreatedBy() string {
	return m.createdBy
}

// UpdatedAt - UpdatedAt のゲッター
func (m *Meta) UpdatedAt() time.Time {
	return m.updatedAt
}

// UpdatedBy - UpdatedBy のゲッター
func (m *Meta) UpdatedBy() string {
	return m.updatedBy
}

// DeletedAt - DeletedAt のゲッター
func (m *Meta) DeletedAt() *time.Time {
	return m.deletedAt
}

// DeletedBy - DeletedBy のゲッター
func (m *Meta) DeletedBy() string {
	return m.deletedBy
}
