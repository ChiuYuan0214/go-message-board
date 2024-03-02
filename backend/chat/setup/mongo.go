package setup

import (
	"chat/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func InitDynamo() *types.DynamoClient {
	sess := session.Must(session.NewSessionWithOptions(session.Options{Config: aws.Config{
		Region: aws.String("ap-northeast-3"),
		// LogLevel:    aws.LogLevel(aws.LogDebugWithHTTPBody),
	}}))

	return &types.DynamoClient{DB: dynamodb.New(sess)}
}
