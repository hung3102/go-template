package mail

import (
	"fmt"
	"testing"
)

func Test_Message(t *testing.T) {
	sendParams := &SendParams{
		ToAddress: "to@address\ncom",
		Subject:   "メールの\nタイトル",
		Body:      "メールの\n本文",
		File: &SendParamFile{
			Data:        []byte("ファイルの\n中身"),
			ContentType: "application\n/pdf",
			Filename:    "ファイル\n名.pdf",
		},
	}

	rawMessage, err := NewMessage().GetMessage(
		&GetMessageParams{
			FromAddress: "from@address\nmail",
			SendParams:  *sendParams,
		})
	if err != nil {
		t.Fatalf("err: %+v", err)
	}
	fmt.Println(string(*rawMessage))
}
