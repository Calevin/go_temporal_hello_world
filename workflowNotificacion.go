package main

import (
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"time"
)

// NotificationWorkflow segundo workflow
func NotificationWorkflow(ctx workflow.Context, email EmailInput) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("Workflow Notification iniciado")

	options := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second * 1,  // cuánto espera entre reintentos al principio
			BackoffCoefficient: 2.0,              // factor multiplicador (exponencial)
			MaximumInterval:    time.Second * 30, // cuánto es lo máximo que puede esperar entre reintentos
			MaximumAttempts:    3,                // cantidad total de intentos (el primero + reintentos)
		},
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	// Ejecutar SendEmail como Activity
	err := workflow.ExecuteActivity(ctx, SendEmail, email).Get(ctx, nil)
	if err != nil {
		logger.Error("Error al enviar email", "err", err)
		return err
	}

	// Ejecutar SaveLog como Activity
	logMessage := "Email enviado correctamente a: " + email.To
	err = workflow.ExecuteActivity(ctx, SaveLog, logMessage).Get(ctx, nil)
	if err != nil {
		logger.Error("Error al guardar log", "err", err)
		return err
	}

	logger.Info("Workflow Notification finalizado con éxito")
	return nil
}
