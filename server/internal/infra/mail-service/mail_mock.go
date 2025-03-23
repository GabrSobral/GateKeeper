package mailservice

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type MailServiceMock struct {
	From    string
	To      string
	Subject string
}

func (ms *MailServiceMock) sendMail(ctx context.Context, params SendMailParams) error {
	fmt.Printf("Mail sent successfully to: %v, from: %v, subject: %v", ms.To, ms.From, ms.Subject)

	return nil
}

func (ms *MailServiceMock) SendEmailConfirmationEmail(ctx context.Context, to, userName, token string) error {
	fmt.Printf("Mail sent successfully to: %v, from: %v, subject: %v", to, ms.From, ms.Subject)

	return nil
}

func (ms *MailServiceMock) SendMfaEmail(ctx context.Context, to, userName, token string) error {
	fmt.Printf("Mail sent successfully to: %v, from: %v, subject: %v", to, ms.From, ms.Subject)

	return nil
}

func (ms *MailServiceMock) SendForgotPasswordEmail(ctx context.Context, to, userName, token string, passwordResetID, applicationID uuid.UUID) error {
	fmt.Printf("Mail sent successfully to: %v, from: %v, subject: %v", to, ms.From, ms.Subject)

	return nil
}
