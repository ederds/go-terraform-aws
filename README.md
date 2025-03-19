# 🚀 Go Terraform AWS Project
## Infraestrutura com Terraform e AWS (S3 + SQS) 🚀

![Terraform](https://img.shields.io/badge/Terraform-v1.5-blue?logo=terraform)
![Go](https://img.shields.io/badge/Go-1.21-blue?logo=go)
![AWS](https://img.shields.io/badge/AWS-S3%20%26%20SQS-orange?logo=amazonaws)

Este projeto implementa uma **infraestrutura AWS usando Terraform**, além de um serviço em **Go** para interagir com **S3 e SQS** dentro do **LocalStack**.

---

## 📌 Tecnologias Utilizadas
- **Go** 🐹 → Aplicação backend
- **Terraform** 🌍 → Infraestrutura como código (IaC)
- **AWS S3** ☁️ → Armazenamento de objetos
- **AWS SQS** 📩 → Mensageria
- **LocalStack** 🏠 → Simulação de serviços AWS localmente
- **Git & GitHub** 🔗 → Versionamento de código

---

## ⚡ Pré-requisitos
Antes de iniciar, certifique-se de ter instalado:
- [Go](https://go.dev/doc/install) (v1.21+)
- [Terraform](https://developer.hashicorp.com/terraform/downloads) (v1.5+)
- [Docker](https://www.docker.com/) (Para rodar o LocalStack)
- [Git](https://git-scm.com/)
- [LocalStack](https://github.com/localstack/localstack)

---

## 🔧 Instalação

### 1️⃣ Clone o repositório
```sh
git clone git@github.com:ederds/go-terraform-aws.git
cd go-terraform-aws
```


### 2️⃣ Configure as variáveis de ambiente
Crie um arquivo .env e adicione as credenciais fake do LocalStack:
```
echo 'AWS_ACCESS_KEY_ID=dummy
AWS_SECRET_ACCESS_KEY=dummy
AWS_DEFAULT_REGION=us-east-1
LOCALSTACK_ENDPOINT=http://localhost:4566
S3_BUCKET=my-teste-bucket' > .env
```

Carregue as variáveis no ambiente:

```export $(cat .env | xargs)```

### 3️⃣ Inicie o LocalStack

`docker run -d --name localstack -p 4566:4566 localstack/localstack`

### 4️⃣ Inicialize o Terraform
```
terraform init
terraform apply -auto-approve
```

### 5️⃣ Rode o serviço em Go

`go run cmd/s3/main.go`

---

## 📂 Estrutura do Projeto

```
go-terraform-aws/
│── cmd/               # Código principal da aplicação
│   ├── s3/            # Serviço que interage com S3
│   ├── sqs/           # Serviço que interage com SQS
│── terraform/         # Infraestrutura como código
│   ├── main.tf        # Configuração do Terraform
│   ├── variables.tf   # Variáveis do Terraform
│── .gitignore         # Arquivos ignorados no versionamento
│── README.md          # Documentação do projeto
│── go.mod             # Gerenciador de dependências do Go
```

---

## 🎯 Como Usar

Criar um Bucket S3
`aws --endpoint-url=http://localhost:4566 s3 mb s3://my-teste-bucket`

Listar Buckets S3
`aws --endpoint-url=http://localhost:4566 s3 ls`

Criar uma Fila SQS
`aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name my-queue`

Enviar uma Mensagem para SQS
`aws --endpoint-url=http://localhost:4566 sqs send-message --queue-url http://localhost:4566/000000000000/my-queue --message-body "Hello World"`

---

### 📢 Contato
📧 Email: ederaxl@gmail.com
🐙 GitHub: ederds