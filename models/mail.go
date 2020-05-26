package models

import (
	"context"
	"time"

	"github.com/mailgun/mailgun-go/v4"
	dt "github.com/nicolauscg/impensa/datatransfers"
)

type MailOrmer interface {
	SendMail(dt.MailParam) (*dt.MailSuccessResponse, error)
	CreateMailMessage(mailParam dt.MailParam) *mailgun.Message
	SendMailMessage(message *mailgun.Message) (*dt.MailSuccessResponse, error)
}

type mailOrm struct {
	mailgun *mailgun.MailgunImpl
}

func NewMailOrmer(mailgun *mailgun.MailgunImpl) *mailOrm {
	return &mailOrm{mailgun}
}

func (o *mailOrm) SendMail(mailParam dt.MailParam) (*dt.MailSuccessResponse, error) {
	message := o.mailgun.NewMessage(
		"mail@impensa.nicolauscg.me",
		mailParam.Subject,
		"",
		mailParam.Recipient,
	)
	message.SetHtml(mailParam.Body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, id, err := o.mailgun.Send(ctx, message)

	if err != nil {
		return nil, err
	}

	return &dt.MailSuccessResponse{Message: resp, Id: id}, nil
}

func (o *mailOrm) CreateMailMessage(mailParam dt.MailParam) *mailgun.Message {
	message := o.mailgun.NewMessage(
		"mail@impensa.nicolauscg.me",
		mailParam.Subject,
		"",
		mailParam.Recipient,
	)
	message.SetHtml(mailParam.Body)

	return message
}

func (o *mailOrm) SendMailMessage(message *mailgun.Message) (*dt.MailSuccessResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, id, err := o.mailgun.Send(ctx, message)
	if err != nil {
		return nil, err
	}
	return &dt.MailSuccessResponse{Message: resp, Id: id}, nil
}
