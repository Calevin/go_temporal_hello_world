package main

import (
	"go.temporal.io/sdk/workflow"
	"time"
)

// HelloWorldWorkflow primer workflow
func HelloWorldWorkflow(ctx workflow.Context, name string) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Workflow HelloWorld iniciado")

	// Simula algo asincrónico
	_ = workflow.Sleep(ctx, 2*time.Second)

	result := "Hola, " + name + "!"
	logger.Info("Resultado de HelloWorld", "mensaje", result)
	return result, nil
}
