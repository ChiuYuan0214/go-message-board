package infra

import (
	"chat/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var _ Dynamo = (*DynamoDB)(nil)

type DynamoDB struct {
	client *types.DynamoClient
}

func (d *DynamoDB) Run() (err error) {
	var sess *session.Session
	sess, err = session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String("ap-northeast-3"),
		},
	})
	if err != nil {
		return err
	}
	d.client = &types.DynamoClient{DB: dynamodb.New(sess)}
	return
}

func (d *DynamoDB) Stop() {}

func (d *DynamoDB) Client() *types.DynamoClient {
	return d.client
}
