#!/usr/bin/env sh

REPO="asia.gcr.io/$GOOGLE_CLOUD_PROJECT/"
IMG="pubsub-pull-subscription"
TAG=`date +%s`

if CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v pullSubscription.go ; then
    if ! docker build . --tag $REPO$IMG:$TAG; then
        echo "\nFailed to build image!\n"
        exit 1
    fi
    if ! gcloud docker -- push $REPO$IMG:$TAG; then
        echo "\nFailed to push image: $REPO$IMG:$TAG\n"
        exit 1
    fi
    echo "\nImage: $REPO$IMG:$TAG is pushed to gcloud\n"
else
    echo "\nCompilation ERROR\n"
fi