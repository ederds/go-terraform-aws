package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func main() {

	ctx := context.Background()
	queueName := os.Getenv("SQS_QUEUE")
	localstackEndpoint := os.Getenv("LOCALSTACK_ENDPOINT")

	// Carrega a configuração do AWS SDK
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal("Erro ao carregar configuração AWS:", err)
	}

	fmt.Println("queuName: %s endpoint: %s\n", queueName, localstackEndpoint)

	// Cria um cliente S3 apontando para o LocalStack
	client := sqs.NewFromConfig(cfg, func(o *sqs.Options) {
		o.BaseEndpoint = aws.String(localstackEndpoint)
	})

	queue, err := client.GetQueueUrl(ctx, &sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody: aws.String("hello, sqs!"),
		QueueUrl:    queue.QueueUrl,
	})

	if err != nil {
		log.Fatal(err)
	}

}
