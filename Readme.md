# GoMailSender API - Integração com Google Cloud Pub/Sub

A GoMailSender é uma API em Golang desenvolvida para facilitar o envio de e-mails, com suporte à integração via Google Cloud Pub/Sub. Essa integração permite que você receba notificações em tempo real sobre o status dos e-mails enviados.

## Pré-requisitos

Antes de começar, certifique-se de ter os seguintes pré-requisitos instalados:

- Go (Golang) >= 1.14
- Conta no Google Cloud Platform (GCP) com o Pub/Sub habilitado
- Configurar arquivo .env com variáveis de conexão smtp

## Configuração do Google Cloud Pub/Sub

1. Crie um tópico e uma assinatura no Google Cloud Pub/Sub para receber notificações.

2. Configure as credenciais da sua conta GCP para autenticar a aplicação. Você pode definir as credenciais através de variáveis de ambiente ou um arquivo de configuração.

## Instalação

1. Clone o repositório:

```bash
git clone https://github.com/MeDeyvid/GoMailSender.git
cd gomailsender
```

2. Instale as dependências:

```bash
go mod download
```

3. Configure as variáveis de ambiente ou crie um arquivo de configuração para as credenciais do Google Cloud Pub/Sub.

## Uso

1. Inicie o serviço da GoMailSender:

```bash
go run main.go
```

2. Envie uma requisição para a API para enviar e-mails:

```bash
curl -X POST \
  http://localhost:8080/send \
  -H 'Content-Type: application/json' \
  -d '{
    "message": {
        "data": "ewogICAiZnJvbSI6ICJlbWFpbEBlbWFpbC5jb20iLAogICAidG8iOiBbCiAgICAgICAiZW1haWxAZW1haWwuY29tIgogICBdLAogICAidGVtcGxhdGUiOiB7CiAgICAgICAiU3ViamVjdCI6ICJzZW5kIG1haWwgd2l0aCB0ZW1wbGF0ZSBzdWJqZWN0IiwKICAgICAgICJodG1sIjogYGNvbnRlbnQgdGVtcGxhdGUuaHRtbGAsCiAgICAgICAiYm9keSI6IHsKICAgICAgICAgICAiTmFtZSI6ICJuYW1lIiwKICAgICAgICAgICAiTWVzc2FnZSI6ICJuYW1lIgogICAgICAgfQogICB9Cn0=",
    },
}'
```

3. Acompanhe os logs do serviço para verificar as notificações do Google Cloud Pub/Sub sobre o status do e-mail enviado.