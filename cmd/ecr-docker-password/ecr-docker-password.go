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

	ecrAwsAccessKeyId := action.GetInput("ecr_aws_access_key_id")
	if ecrAwsAccessKeyId == "" {
		action.Fatalf("ecr_aws_access_key_id is required")
	}

	ecrAwsSecretAccessKey := action.GetInput("ecr_aws_secret_access_key")
	if ecrAwsSecretAccessKey == "" {
		action.Fatalf("ecr_aws_secret_access_key is required")
	}

	ecrAwsRegion := action.GetInput("ecr_aws_region")

	ecrClient := ecr.New(ecrAwsAccessKeyId, ecrAwsSecretAccessKey, ecrAwsRegion)
	_, err := ecrClient.GetPassword(ctx)
	if err != nil {
		action.Fatalf("error getting ECR password: %s", err)
	}

	password := "megakekurepassword1234"

	action.AddMask(password)
	action.SetOutput("ecr-docker-password", password)
	action.SaveState("ecr-docker-password", password)
}
