package mail

import (
	"github.com/jhillyerd/enmime"
	"golang.org/x/xerrors"
)

// Message - メールの本文を作成
type Message struct {
	// メール本文を格納する変数
	msg []byte
}

// NewMessage - メールの本文を作成
func NewMessage() *Message {
	return &Message{}
}

// GetMessageParams - メール本文作成用のパラメーター
type GetMessageParams struct {
	FromAddress string
	SendParams
}

// GetMessage - メール本文を作成する
func (m *Message) GetMessage(params *GetMessageParams) (*[]byte, error) {
	builder := m.builder(params)
	err := builder.Send(m) // 本文を作成(Sendの中で、m.Sendが実行される)
	if err != nil {
		return nil, xerrors.Errorf("error in Message.GetMessage: %w", err)
	}
	return &m.msg, nil
}

// builder - builderを作成する
func (m *Message) builder(params *GetMessageParams) *enmime.MailBuilder {
	builder := enmime.Builder().
		From("", params.FromAddress).
		To("", params.ToAddress).
		Subject(params.Subject).
		Text([]byte(params.Body))
	if params.File != nil {
		builder = m.attachment(builder, params.File)
	}
	return &builder
}

// attachment - builder作成の添付ファイルの設定をする
func (m *Message) attachment(builder enmime.MailBuilder, file *SendParamFile) enmime.MailBuilder {
	builder = builder.AddAttachment(
		file.Data,
		file.ContentType,
		file.Filename,
	)
	return builder
}

// Send - メッセージ情報を取得する
func (m *Message) Send(from string, tos []string, msg []byte) error {
	m.msg = msg
	return nil
}
