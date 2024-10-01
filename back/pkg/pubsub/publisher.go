package pubsub

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
	"golang.org/x/xerrors"
)

// PubsubPublisher - PubsubPublisherの構造体
type PubsubPublisher struct {
	client *pubsub.Client
}

// NewPubsubPublisher - PubsubPublisherの初期化
func NewPubsubPublisher(client *pubsub.Client) *PubsubPublisher {
	return &PubsubPublisher{client: client}
}

// Publisher - Publisherのインターフェース
type Publisher interface {
	PublishMessage(ctx context.Context, topicID string, attr map[string]string, msg string) error
}

// PublishMessage - メッセージをpubsubに送信する
func (p *PubsubPublisher) PublishMessage(ctx context.Context, topicID string, attr map[string]string, msg string) error {
	json, err := json.Marshal(msg)
	if err != nil {
		return xerrors.Errorf("PublishMessage - Input msg error : %w", err)
	}

	t := p.client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data:       []byte(json),
		Attributes: attr,
	})

	_, err = result.Get(ctx)
	if err != nil {
		return xerrors.Errorf("PublishMessage - Get: %w", err)
	}

	return nil
}
