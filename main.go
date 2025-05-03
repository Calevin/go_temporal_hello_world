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
	options := client.StartWorkflowOptions{
		ID:        "workflow-hola-mundo",
		TaskQueue: "hola-mundo-task-queue",
	}

	// Ejecutar el workflow
	we, err := c.ExecuteWorkflow(context.Background(), options, HelloWorldWorkflow, "Sebastián")
	if err != nil {
		log.Fatalln("Error al ejecutar el workflow:", err)
	}

	log.Println("Workflow lanzado:", "ID:", we.GetID(), "RunID:", we.GetRunID())

	// Esperar el resultado
	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Error al obtener el resultado:", err)
	}

	log.Println("Resultado del workflow:", result)
}
