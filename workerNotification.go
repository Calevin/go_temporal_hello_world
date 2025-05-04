package main

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
)

func main() {
	// Crear cliente Temporal
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("No se pudo conectar a Temporal:", err)
	}
	defer c.Close()

	// Crear worker y registrar el workflow
	w := worker.New(c, "notificaciones-queue", worker.Options{})
	w.RegisterWorkflow(NotificationWorkflow)
	w.RegisterActivity(SendEmail)
	w.RegisterActivity(SaveLog)

	log.Println("Worker esperando tareas...")
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Fallo al correr el worker:", err)
	}
}
