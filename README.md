Infraestrutura com Terraform e AWS (S3 + SQS) 🚀

Este projeto implementa uma infraestrutura AWS usando Terraform, além de um serviço em Go para interagir com S3 e SQS dentro do LocalStack.

📌 Tecnologias Utilizadas
Go 🐹 → Aplicação backend
Terraform 🌍 → Infraestrutura como código (IaC)
AWS S3 ☁️ → Armazenamento de objetos
AWS SQS 📩 → Mensageria
LocalStack 🏠 → Simulação de serviços AWS localmente
Git & GitHub 🔗 → Versionamento de código
⚡ Pré-requisitos
Antes de iniciar, certifique-se de ter instalado:

Go (v1.21+)
Terraform (v1.5+)
Docker (Para rodar o LocalStack)
Git
LocalStack
🔧 Instalação
1️⃣ Clone o repositório
sh
Copiar
Editar
git clone git@github.com:ederds/go-terraform-aws.git
cd go-terraform-aws
2️⃣ Configure as variáveis de ambiente
Crie um arquivo .env e adicione as credenciais fake do LocalStack:

sh
Copiar
Editar
echo 'AWS_ACCESS_KEY_ID=dummy
AWS_SECRET_ACCESS_KEY=dummy
AWS_DEFAULT_REGION=us-east-1
LOCALSTACK_ENDPOINT=http://localhost:4566
S3_BUCKET=my-teste-bucket' > .env
Carregue as variáveis no ambiente:

sh
Copiar
Editar
export $(cat .env | xargs)
3️⃣ Inicie o LocalStack
sh
Copiar
Editar
docker run -d --name localstack -p 4566:4566 localstack/localstack
4️⃣ Inicialize o Terraform
sh
Copiar
Editar
terraform init
terraform apply -auto-approve
5️⃣ Rode o serviço em Go
sh
Copiar
Editar
go run cmd/s3/main.go
📂 Estrutura do Projeto
bash
Copiar
Editar
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
🎯 Como Usar
Criar um Bucket S3
sh
Copiar
Editar
aws --endpoint-url=http://localhost:4566 s3 mb s3://my-teste-bucket
Listar Buckets S3
sh
Copiar
Editar
aws --endpoint-url=http://localhost:4566 s3 ls
Criar uma Fila SQS
sh
Copiar
Editar
aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name my-queue
Enviar uma Mensagem para SQS
sh
Copiar
Editar
aws --endpoint-url=http://localhost:4566 sqs send-message --queue-url http://localhost:4566/000000000000/my-queue --message-body "Hello World"
🛠 Contribuindo
Sinta-se à vontade para contribuir! Basta seguir os passos:

Faça um fork do repositório
Crie uma branch (git checkout -b minha-feature)
Commit suas mudanças (git commit -m "Adicionei uma nova feature")
Faça um push (git push origin minha-feature)
Abra um Pull Request
📜 Licença
Este projeto está licenciado sob a MIT License – veja o arquivo LICENSE para mais detalhes.

📢 Contato
📧 Email: ederds@example.com
🐙 GitHub: ederds

