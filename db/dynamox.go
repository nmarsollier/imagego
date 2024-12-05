package db

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/nmarsollier/imagego/tools/env"
	"github.com/nmarsollier/imagego/tools/errs"
)

var (
	dynamo_once     sync.Once
	dynamo_instance *DynamoDao
)

var tableName = "imagego"

func getDynamoDb() ImageDao {
	dynamo_once.Do(func() {
		customCreds := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(
			env.Get().AwsAccessKeyId,
			env.Get().AwsSecret,
			"",
		))

		cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(env.Get().AwsRegion),
			config.WithCredentialsProvider(customCreds),
		)
		if err != nil {
			log.Fatalf("Error cargando la configuración: %v", err)
		}

		client := dynamodb.NewFromConfig(cfg)

		dynamo_instance = &DynamoDao{
			client: client,
		}
	})

	return dynamo_instance
}

type DynamoDao struct {
	client *dynamodb.Client
}

func (r *DynamoDao) Get(key string) (string, error) {
	image := Image{ID: key, Image: ""}
	imageId, err := attributevalue.Marshal(image.ID)
	if err != nil {
		return "", err
	}

	response, err := r.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{"img": imageId}, TableName: &tableName,
	})

	if err != nil {
		return "", err
	}

	if response == nil || response.Item == nil {
		return "", errs.NotFound
	}

	err = attributevalue.UnmarshalMap(response.Item, &image)
	if err != nil {
		return "", err
	}

	return image.Image, nil
}

func (r *DynamoDao) Set(key string, value string, expiration time.Duration) (string, error) {
	image, err := attributevalue.MarshalMap(Image{ID: key, Image: value})
	if err != nil {
		return "", err
	}

	_, err = r.client.PutItem(
		context.TODO(),
		&dynamodb.PutItemInput{
			TableName: &tableName,
			Item:      image,
		},
	)
	if err != nil {
		return "", err
	}
	return key, nil
}
