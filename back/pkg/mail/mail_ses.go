package mail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"golang.org/x/xerrors"
)

var _ Mail = (*mailSES)(nil)

// mailSES - SESでメールを送信する
type mailSES struct {
	sesService  *sesv2.Client // SESクライアント
	fromAddress string        // 送信元メールアドレス
}

// NewMailSESParams - NewMailSESのパラメーター
type NewMailSESParams struct {
	SesService  *sesv2.Client // SESクライアント
	FromAddress string        // 送信元メールアドレス
}

// NewMailSES - MailSESを作成する
func NewMailSES(params *NewMailSESParams) *mailSES {
	return &mailSES{
		sesService:  params.SesService,
		fromAddress: params.FromAddress,
	}
}

// Send - メールを送信する
func (m *mailSES) Send(ctx context.Context, params *SendParams) error {
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
func (m *mailSES) sendEmailInput(message *[]byte) *sesv2.SendEmailInput {
	return &sesv2.SendEmailInput{
		Content: &types.EmailContent{
			Raw: &types.RawMessage{
				Data: *message,
			},
		},
	}
}
