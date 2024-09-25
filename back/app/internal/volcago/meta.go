package volcago

import "time"

// Meta - meta information for optimistic exclusive lock
type Meta struct {
	CreatedAt time.Time  `firestore:"created_at"` // 作成日時
	CreatedBy string     `firestore:"created_by"` // 作成者
	UpdatedAt time.Time  `firestore:"updated_at"` // 更新日時
	UpdatedBy string     `firestore:"updated_by"` // 更新者
	DeletedAt *time.Time `firestore:"deleted_at"` // 削除日時
	DeletedBy string     `firestore:"deleted_by"` // 削除者
	Version   int        `firestore:"version"`    // バージョン
}
