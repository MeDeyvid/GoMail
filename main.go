package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

type PubSubMessage struct {
	Message struct {
		Data []byte `json:"data" validate:"required"`
		ID   string `json:"id"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

type Template struct {
  Html string `json:"html"`
  Subject string `json:"subject" validate:"required"`
  Body interface {} `json:"body"`
}

type DataMessage struct {
  From string `json:"from" validate:"required"`
  To []string `json:"to" validate:"required"`
  Template Template `json:"template,omitempty" `
}

func SendMail(w http.ResponseWriter, r *http.Request) {
  var data PubSubMessage
  errParse := json.NewDecoder(r.Body).Decode(&data)
  if errParse != nil {
    http.Error(w, errParse.Error(), http.StatusBadRequest)
    return
  }

  var subMessage DataMessage
  json.Unmarshal(data.Message.Data, &subMessage)

  user := os.Getenv("EMAIL_USER")
  password := os.Getenv("EMAIL_PASSWORD")
  smtpHost := os.Getenv("SMTP_HOST")
  smtpPort := os.Getenv("SMTP_PORT")

  auth := smtp.PlainAuth("", user, password, smtpHost)

  t := template.Must(template.New("myname").Parse(subMessage.Template.Html))

  var body bytes.Buffer

  mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
  body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", subMessage.Template.Subject , mimeHeaders)))

  if len([]rune(subMessage.Template.Html)) != 0 {
    t.Execute(&body, subMessage.Template.Body)
  }

  err := smtp.SendMail(smtpHost+":"+smtpPort, auth, subMessage.From, subMessage.To, body.Bytes())
  if err != nil {
    http.Error(w, err.Error(), http.StatusPreconditionFailed)
    return
  }
}

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Printf("Error loading .env file")
  }

  r := chi.NewRouter()
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)

  r.Use(middleware.Timeout(60 * time.Second))

  r.Post("/send", SendMail)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}