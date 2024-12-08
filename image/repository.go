package image

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/nmarsollier/imagego/tools/db"
	"github.com/nmarsollier/imagego/tools/errs"
	"github.com/nmarsollier/imagego/tools/log"
)

var tableName = "images"

// Insert adds an image to the db
func Insert(image *Image, deps ...interface{}) (imageId string, err error) {
	if err = image.ValidateSchema(deps...); err != nil {
		log.Get(deps...).Error(err)
		return
	}

	client := db.Get(deps...)

	imageData, err := attributevalue.MarshalMap(image)
	if err != nil {
		return
	}

	_, err = client.PutItem(
		context.TODO(),
		&dynamodb.PutItemInput{
			TableName: &tableName,
			Item:      imageData,
		},
	)

	if err != nil {
		log.Get(deps...).Error(err)
		return
	}

	return image.ID, nil
}

// Find finds and returns an image from the database
func find(imageID string, deps ...interface{}) (image *Image, err error) {
	client := db.Get(deps...)

	response, err := client.GetItem(
		context.TODO(),
		&dynamodb.GetItemInput{
			Key: map[string]types.AttributeValue{
				"id": &types.AttributeValueMemberS{
					Value: imageID,
				}},
			TableName: &tableName,
		},
	)

	if err != nil || response == nil || response.Item == nil {
		log.Get(deps...).Error(err)

		return nil, errs.NotFound
	}

	err = attributevalue.UnmarshalMap(response.Item, &image)
	if err != nil {
		log.Get(deps...).Error(err)

		return nil, errs.NotFound
	}

	return
}
