package log2slack

import (
	"context"

	"github.com/ashwanthkumar/slack-go-webhook"
)

// Do - stream log to slack
func Do(_ context.Context, cfg *Config, attachments []slack.Attachment) error {
	payload := slack.Payload{
		Username:    cfg.SlackUserName,
		IconEmoji:   cfg.SlackIconEmoji,
		Attachments: attachments,
	}

	if errs := slack.Send(cfg.SlackWebhookURL, "", payload); len(errs) > 0 {
		return errs[0]
	}

	return nil
}

// Config は Do() に渡す構成情報
type Config struct {
	SlackUserName   string // 通知時のユーザー名
	SlackIconEmoji  string // 通知時のアイコン
	SlackWebhookURL string // 通知先のWebhook URL
}
