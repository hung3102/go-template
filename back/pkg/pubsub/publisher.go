package pubsub

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/topgate/gcim-temporary/back/pkg/environ"
)

// PublishMsg - メッセージをpubsubに送信する
func PublishMsg(ctx context.Context, topicID string, attr map[string]string, msg string) error {
	projectID := environ.ProjectID()

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}

	json, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("Input msg error : %v", err)
	}

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data:       []byte(json),
		Attributes: attr,
	})

	_, err = result.Get(ctx)
	if err != nil {
		return fmt.Errorf("Get: %v", err)
	}

	return nil
}
