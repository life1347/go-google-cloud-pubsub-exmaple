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

	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"github.com/satori/go.uuid"
	"tachingchen.com/googlePubSub/common"
)

func main() {
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		fmt.Fprintf(os.Stderr, "GOOGLE_CLOUD_PROJECT environment variable must be set.\n")
		os.Exit(1)
	}

	client := common.CreateClient(projectID)

	// Create topic if not exists
	const topicName string = "example-topic"
	topic := common.CreateTopicIfNotExists(client, topicName)

	// Create a new subscription.
	subName := fmt.Sprintf("example-sub-%s", uuid.NewV4().String())
	if err := common.CreateSub(client, subName, topic, nil); err != nil {
		log.Println(err)
	}

	// Pull messages via the subscription.
	if err := pullMsgs(client, subName, topic); err != nil {
		log.Println(err)
	}

	// Delete the subscription.
	if err := common.DeleteSub(client, subName); err != nil {
		log.Fatal(err)
	}
}

func pullMsgs(client *pubsub.Client, name string, topic *pubsub.Topic) error {
	ctx := context.Background()

	sub := client.Subscription(name)
	it, err := sub.Pull(ctx)
	if err != nil {
		return err
	}
	defer it.Stop()

	for i := 0; i < 10; i++ {
		msg, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		session := &common.Session{}
		err = json.Unmarshal(msg.Data, session)
		if err == nil {
			log.Printf("Got message: %s\n", session.SessionID)
		}
		msg.Done(true)
	}
	return nil
}
