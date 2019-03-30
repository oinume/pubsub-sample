# pubsub-sample

Cloud Pub/Sub sample

## Requirements

- Go 1.11 or later
- Google Cloud Platform account

## How to run

### Publish a message

```shell
$ export GOOGLE_CLOUD_PROJECT=<your project name>
$ export GO111MODULE=on
$ go run cmd/publisher/publisher.go <topic> <message>
```

Example:

```shell
$ go run cmd/publisher/publisher.go first-topic "Hello world"
```

### Subscribe a topic

With GAE

```shell
$ cd gae
$ gcloud app deploy
```
