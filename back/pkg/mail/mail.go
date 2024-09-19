// package mail - メール送信処理
package mail

import (
	"context"
)

// Mail - メール送信処理
type Mail interface {
	// メールを送信する
	Send(ctx context.Context, params *SendParams) error
}

// SendParams - メール送信のパラメーター
type SendParams struct {
	// 送信先メールアドレス
	ToAddress string
	// メールタイトル
	Subject string
	// メール本文
	Body string
	// 添付ファイル
	File *SendParamFile
}

// SendParamFile - 添付ファイル
type SendParamFile struct {
	// Data - ファイルの内容
	Data []byte
	// ContentType - ContentType
	ContentType string
	// Filename - ファイル名
	Filename string
}
