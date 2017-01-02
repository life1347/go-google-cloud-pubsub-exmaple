// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Command topics is a tool to manage Google Cloud Pub/Sub topics by using the Pub/Sub API.
// See more about Google Cloud Pub/Sub at https://cloud.google.com/pubsub/docs/overview.

// This file is modified by Ta-Ching Chen (https://tachingchen.com/) on 2017.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"tachingchen.com/googlePubSub/common"
)

func main() {
	ctx := context.Background()

	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		fmt.Fprintf(os.Stderr, "GOOGLE_CLOUD_PROJECT environment variable must be set.\n")
		os.Exit(1)
	}

	client := common.CreateClient(projectID)

	const topicName string = "example-topic"
	topic := common.CreateTopicIfNotExists(client, topicName)

	for i := 0; i < 10; i++ {
		msgUuid := uuid.NewV4().String()
		// message we want to send to subscriber
		session, _ := json.Marshal(&common.Session{
			SessionID: msgUuid,
			TimeStamp: time.Now().Unix(),
		})
		// publish message to Cloud Pub/Sub
		_, _ = topic.Publish(ctx, &pubsub.Message{
			Data: session,
		})
		log.Printf("%s send", msgUuid)
		time.Sleep(1 * time.Second)
	}
	// DO NOT DELETE the topic before messages are consumed by client
	// or you will lost the messages left
	// topic.Delete(ctx)
}
