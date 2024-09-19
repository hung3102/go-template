// package mail - メール送信処理
package mail

import (
	"context"
)

// Mail - メール送信処理
type Mail interface {
	// Send - メールを送信する
	Send(ctx context.Context, params *SendParams) error
}

// SendParams - メール送信のパラメーター
type SendParams struct {
	ToAddress string         // 送信先メールアドレス
	Subject   string         // メールタイトル
	Body      string         // メール本文
	File      *SendParamFile // 添付ファイル
}

// SendParamFile - 添付ファイル
type SendParamFile struct {
	Data        []byte // ファイルの内容
	ContentType string // ContentType
	Filename    string // ファイル名
}
