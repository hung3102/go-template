package mail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"golang.org/x/xerrors"
)

var _ Mail = (*MailSES)(nil)

// MailSES - SESでメールを送信する
type MailSES struct {
	// SESクライアント
	sesService *sesv2.Client
	// 送信元メールアドレス
	fromAddress string
}

// NewMailSESParams - NewMailSESのパラメーター
type NewMailSESParams struct {
	// SESクライアント
	SesService *sesv2.Client
	// 送信元メールアドレス
	FromAddress string
}

// NewMailSES - MailSESを作成する
func NewMailSES(params *NewMailSESParams) *MailSES {
	return &MailSES{
		sesService:  params.SesService,
		fromAddress: params.FromAddress,
	}
}

// Send - メールを送信する
func (m *MailSES) Send(ctx context.Context, params *SendParams) error {
	rawMessage, err := NewMessage().GetMessage(&GetMessageParams{
		FromAddress: m.fromAddress,
		SendParams:  *params,
	})
	if err != nil {
		return xerrors.Errorf("error in MailSES.Send: %w", err)
	}

	_, err = m.sesService.SendEmail(ctx, m.sendEmailInput(rawMessage))
	if err != nil {
		return xerrors.Errorf("error in MailSES.Send: %w", err)
	}
	return nil
}

// sendEmailInput - sesのSendEmailのパラメーターを作成する
func (m *MailSES) sendEmailInput(message *[]byte) *sesv2.SendEmailInput {
	return &sesv2.SendEmailInput{
		Content: &types.EmailContent{
			Raw: &types.RawMessage{
				Data: *message,
			},
		},
	}
}
