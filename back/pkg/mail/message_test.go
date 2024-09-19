package mail_test

import (
	"fmt"
	"testing"

	"github.com/topgate/gcim-temporary/back/pkg/mail"
)

func Test_Message(t *testing.T) {
	sendParams := &mail.SendParams{
		ToAddress: "to@address\ncom",
		Subject:   "メールの\nタイトル",
		Body:      "メールの\n本文",
		File: &mail.SendParamFile{
			Data:        []byte("ファイルの\n中身"),
			ContentType: "application\n/pdf",
			Filename:    "ファイル\n名.pdf",
		},
	}

	rawMessage, err := mail.NewMessage().GetMessage(
		&mail.GetMessageParams{
			FromAddress: "from@address\nmail",
			SendParams:  *sendParams,
		})
	if err != nil {
		t.Fatalf("err: %+v", err)
	}
	fmt.Println(string(*rawMessage))
}
