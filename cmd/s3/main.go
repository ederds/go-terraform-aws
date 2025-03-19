package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {

	ctx := context.Background()

	// Obtém o nome do bucket e o endpoint do LocalStack das variáveis de ambiente
	bucketName := os.Getenv("S3_BUCKET")
	localstackEndpoint := os.Getenv("S3_LOCALSTACK_ENDPOINT")

	if bucketName == "" || localstackEndpoint == "" {
		log.Fatal("Erro: As variáveis de ambiente S3_BUCKET e LOCALSTACK_ENDPOINT devem estar definidas.")
	}

	// Carrega a configuração do AWS SDK
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal("Erro ao carregar configuração AWS:", err)
	}

	// Cria um cliente S3 apontando para o LocalStack
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(localstackEndpoint)
	})

	// Tenta obter um objeto do bucket
	output, err := client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("go.mod"),
	})
	if err != nil {
		log.Fatal("Erro ao obter objeto do S3:", err)
	}
	defer output.Body.Close()

	// Lê o conteúdo do objeto
	b, err := io.ReadAll(output.Body)
	if err != nil {
		log.Fatal("Erro ao ler o conteúdo do objeto:", err)
	}

	// Exibe o conteúdo do arquivo
	fmt.Println(string(b))
}
