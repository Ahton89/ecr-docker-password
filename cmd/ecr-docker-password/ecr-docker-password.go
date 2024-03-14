package main

import (
	"context"
	"ecr-docker-password/internal/ecr"
	"github.com/sethvargo/go-githubactions"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	action := githubactions.New()

	ecrAwsAccessKeyId := action.GetInput("access_key_id")
	if ecrAwsAccessKeyId == "" {
		action.Fatalf("access_key_id is required")
	}

	ecrAwsSecretAccessKey := action.GetInput("secret_access_key")
	if ecrAwsSecretAccessKey == "" {
		action.Fatalf("secret_access_key is required")
	}

	ecrAwsRegion := action.GetInput("region")

	ecrClient := ecr.New(ecrAwsAccessKeyId, ecrAwsSecretAccessKey, ecrAwsRegion)
	_, err := ecrClient.GetPassword(ctx)
	if err != nil {
		action.Fatalf("error getting ECR password: %s", err)
	}

	password := "megakekurepassword1234"

	action.SetOutput("ecr-docker-password", password)
	action.SaveState("ecr-docker-password", password)
}
