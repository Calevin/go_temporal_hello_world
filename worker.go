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
		log.Fatalln("Error al conectar con Temporal:", err)
	}
	defer c.Close()

	// Crear worker y registrar el workflow
	w := worker.New(c, "hola-mundo-task-queue", worker.Options{})
	w.RegisterWorkflow(HelloWorldWorkflow)

	log.Println("Esperando workflows...")
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Error al correr el worker:", err)
	}
}
