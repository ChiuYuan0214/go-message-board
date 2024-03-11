package types

import (
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

type DynamoChat struct {
	SenderId   uint64    `json:"senderId"`
	ReceiverId uint64    `json:"receiverId"`
	Content    string    `json:"content"`
	Time       time.Time `json:"time"`
}

type DynamoClient struct {
	DB *dynamodb.DynamoDB
}

func (dc *DynamoClient) GetAllWithFilters(senderId, receiverId uint64, startTime, endTime time.Time) ([]DynamoChat, error) {
	filterExpression := "senderId = :senderId AND receiverId = :receiverId AND #time < :endTime"

	expressionAttributeValues := map[string]*dynamodb.AttributeValue{
		":senderId":   {N: aws.String(fmt.Sprint(senderId))},
		":receiverId": {N: aws.String(fmt.Sprint(receiverId))},
		":endTime":    {S: aws.String(endTime.Format(time.RFC3339))},
	}

	expressionAttributeNames := map[string]*string{
		"#time": aws.String("time"),
	}

	input := &dynamodb.ScanInput{
		TableName:                 aws.String("chat-messages"),
		FilterExpression:          aws.String(filterExpression),
		ExpressionAttributeValues: expressionAttributeValues,
		ExpressionAttributeNames:  expressionAttributeNames,
	}

	if startTime.After(time.Now()) {
		filterExpression += " AND #time >= :startTime"
		expressionAttributeValues[":startTime"] = &dynamodb.AttributeValue{S: aws.String(startTime.Format(time.RFC3339))}
		(*input).Limit = aws.Int64(20)
	}

	var chats []DynamoChat

	result, err := dc.DB.Scan(input)
	if err != nil {
		return chats, err
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &chats)
	if err != nil {
		return chats, err
	}

	sort.Slice(chats, func(i, j int) bool {
		return chats[i].Time.After(chats[j].Time)
	})

	return chats, nil
}

func (dc *DynamoClient) BatchInsert(chatList []DynamoChat) bool {
	for _, chat := range chatList {
		log.Printf("senderId = %d. receiverId = %d. content = %s. time = %d", chat.SenderId, chat.ReceiverId, chat.Content, chat.Time.Unix())
	}
	writeRequests := make([]*dynamodb.WriteRequest, len(chatList))
	for i, chat := range chatList {
		writeRequests[i] = &dynamodb.WriteRequest{
			PutRequest: &dynamodb.PutRequest{
				Item: map[string]*dynamodb.AttributeValue{
					"chatId": {
						S: aws.String(uuid.New().String()),
					},
					"senderId": {
						N: aws.String(fmt.Sprint(chat.SenderId)),
					},
					"receiverId": {
						N: aws.String(fmt.Sprint(chat.ReceiverId)),
					},
					"content": {
						S: aws.String(chat.Content),
					},
					"time": {
						S: aws.String(chat.Time.Format(time.RFC3339)),
					},
				},
			},
		}
	}

	input := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			"chat-messages": writeRequests,
		},
	}

	output, err := dc.DB.BatchWriteItem(input)
	if err != nil {
		log.Println("error when insert into dynamo:", err)
	} else {
		log.Println("insertion output:", output)
	}

	return err == nil
}
