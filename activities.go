package main

import (
	"context"
	"fmt"
)

type EmailInput struct {
	To      string
	Subject string
	Body    string
}

func SendEmail(ctx context.Context, input EmailInput) error {
	fmt.Println("(SendEmail) Enviando email a:", input.To)
	fmt.Println("(SendEmail) Asunto:", input.Subject)
	fmt.Println("(SendEmail) Contenido:", input.Body)
	// Acá iría lógica real con SMTP, etc.
	return nil
}

func SaveLog(ctx context.Context, message string) error {
	fmt.Println("(SaveLog) Guardando log:", message)
	// Aca se podría guardar en archivo, base, etc.
	return nil
}
