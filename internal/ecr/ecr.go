package ecr

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	serviceEcr "github.com/aws/aws-sdk-go-v2/service/ecr"
)

type ecr struct {
	ecrAwsAccessKeyId     string
	ecrAwsSecretAccessKey string
	ecrAwsRegion          region
}

type region string

func (e *region) String() string {
	if e == nil || *e == "" {
		return "us-east-2"
	}

	return string(*e)
}

type ECR interface {
	GetPassword(ctx context.Context) (string, error)
}

func New(AwsAccessKeyId, AwsSecretAccessKey, AwsRegion string) ECR {
	return &ecr{
		ecrAwsAccessKeyId:     AwsAccessKeyId,
		ecrAwsSecretAccessKey: AwsSecretAccessKey,
		ecrAwsRegion:          region(AwsRegion),
	}
}

func (e *ecr) GetPassword(ctx context.Context) (string, error) {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithDefaultRegion(e.ecrAwsRegion.String()),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				e.ecrAwsAccessKeyId,
				e.ecrAwsSecretAccessKey,
				"",
			),
		),
	)
	if err != nil {
		return "", err
	}

	ecrClient := serviceEcr.NewFromConfig(cfg)

	ecrToken, err := ecrClient.GetAuthorizationToken(ctx, &serviceEcr.GetAuthorizationTokenInput{})
	if err != nil {
		return "", err
	}

	if len(ecrToken.AuthorizationData) != 1 {
		return "", fmt.Errorf("no authorization data found or something went wrong")
	}

	return aws.ToString(ecrToken.AuthorizationData[0].AuthorizationToken), nil
}
