package volcago

import "time"

// Meta - meta information for optimistic exclusive lock
type Meta struct {
	CreatedAt time.Time  `` // 作成日時
	CreatedBy string     `` // 作成者
	UpdatedAt time.Time  `` // 更新日時
	UpdatedBy string     `` // 更新者
	DeletedAt *time.Time `` // 削除日時
	DeletedBy string     `` // 削除者
	Version   int        `` // バージョン
}
