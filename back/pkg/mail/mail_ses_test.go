package mail

// import (
// 	"context"
// 	"testing"

// 	"github.com/aws/aws-sdk-go-v2/service/sesv2"
// )

// func Test_MailSES_Send(t *testing.T) {
// 	ctx := context.Background()
// 	sendParams := &SendParams{
// 		ToAddress: "to@address\ncom",
// 		Subject:   "メールの\nタイトル",
// 		Body:      "メールの\n本文",
// 		File: &SendParamFile{
// 			Data:        []byte("ファイルの\n中身"),
// 			ContentType: "application\n/pdf",
// 			Filename:    "ファイル\n名.pdf",
// 		},
// 	}
// 	mail := NewMailSES(&NewMailSESParams{
// 		SesService:  sesv2.New(sesv2.Options{}),
// 		FromAddress: "from@address.mail",
// 	})
// 	err := mail.Send(ctx, sendParams)
// 	if err != nil {
// 		t.Fatalf("err: %+v", err)
// 	}
// }
