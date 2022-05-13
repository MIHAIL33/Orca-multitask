package mailer

import (
	"bytes"
	"log"
	"net/smtp"
	"os"
	"text/template"

	"github.com/spf13/viper"
)

type TypeMail string

const (
	Successed TypeMail = "templates/mail-successed.html"
	Failed TypeMail  = "templates/mail-failed.html"
	Finished TypeMail  = "templates/mail-finished.html"
)

func (t TypeMail) String() string {
	switch t {
	case Successed :
		return "Successed"
	case Failed:
		return "Failed"
	case Finished:
		return "Finished"
	default: 
		return ""
	}
}

type DataMail struct {
	Hostname string
	Dir string
}

type Mailer struct {
	from string
	password string
	Addressees []string
	body string
	typeMail TypeMail
}

func NewMailer(dir string, typeMail TypeMail) *Mailer {
	var data DataMail
	data.Hostname = getHostname()
	data.Dir = dir
	body, err := parseTemlate(string(typeMail), data)
	if err != nil {
		log.Println(err)
	}
	return &Mailer{
		from: os.Getenv("FROM_MAILER"), 
		password: os.Getenv("PASS_MAILER"), 
		Addressees: viper.GetStringSlice("email.addressees"),
		body: body,
		typeMail: typeMail,
	}
}

func (m *Mailer) SendMail() error{
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + m.typeMail.String() + "!\n" 
	msg := []byte(subject + mime + "\n" + m.body)
	auth := smtp.PlainAuth("", m.from, m.password, "smtp.gmail.com")
	addr := "smtp.gmail.com:587"

	if err := smtp.SendMail(addr, auth, m.from, m.Addressees, msg); err != nil {
		return err
	}

	return nil
}

func parseTemlate(templateFileName string, data interface{}) (string, error) {
	template, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = template.Execute(buf, data); err != nil {
		return "", err
	}
	body := buf.String()
	return body, nil
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Println(err)
	}
	return hostname
}