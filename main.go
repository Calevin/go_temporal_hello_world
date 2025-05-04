package main

import (
	"context"
	"go.temporal.io/sdk/client"
	"log"
)

func main() {
	// Crear cliente Temporal
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Error al conectar con Temporal:", err)
	}
	defer c.Close()

	// Opciones del workflow
	optionsHW := client.StartWorkflowOptions{
		ID:        "workflow-hola-mundo",
		TaskQueue: "hola-mundo-task-queue",
	}

	// Ejecutar el workflow
	weHW, err := c.ExecuteWorkflow(context.Background(), optionsHW, HelloWorldWorkflow, "Sebastián")
	if err != nil {
		log.Fatalln("Error al ejecutar el workflow:", err)
	}

	log.Println("Workflow HelloWorld lanzado:", "ID:", weHW.GetID(), "RunID:", weHW.GetRunID())

	// Esperar el resultado
	var result string
	err = weHW.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Error al obtener el resultado workflow HelloWorld:", err)
	}

	log.Println("Resultado del workflow HelloWorld:", result)

	// NotificationWorkflow
	optionsNotification := client.StartWorkflowOptions{
		ID:        "workflow-envio-email",
		TaskQueue: "notificaciones-queue",
	}

	input := EmailInput{
		To:      "sebastian@ejemplo.com",
		Subject: "Bienvenido",
		Body:    "Gracias por unirte a Temporal.",
	}

	weNot, err := c.ExecuteWorkflow(context.Background(), optionsNotification, NotificationWorkflow, input)
	if err != nil {
		log.Fatalln("Error al lanzar el workflow:", err)
	}

	log.Println("Workflow Notification iniciado:", "ID:", weNot.GetID(), "RunID:", weNot.GetRunID())

	err = weNot.Get(context.Background(), nil)
	if err != nil {
		log.Fatalln("Workflow Notification falló:", err)
	}

	log.Println("Workflow Main finalizó correctamente")
}
