package rubbitMQExemple

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	// Crée une connexion à RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Erreur de connexion à RabbitMQ : %s", err)
	}
	defer conn.Close()

	// Crée un canal de communication
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Erreur lors de l'ouverture du canal : %s", err)
	}
	defer channel.Close()

	// Déclare une file d'attente
	queue, err := channel.QueueDeclare(
		"hello", // Nom de la file d'attente
		false,   // Durable : la file d'attente ne survit pas au redémarrage
		false,   // Exclu : la file d'attente est accessible seulement par le créateur
		false,   // AutoDelete : la file d'attente est supprimée lorsqu'elle est vide
		false,   // Arguments : aucun argument supplémentaire
		nil,
	)
	if err != nil {
		log.Fatalf("Erreur lors de la déclaration de la file d'attente : %s", err)
	}

	// Message à envoyer
	message := "Hello RabbitMQ!"

	// Envoi du message à la file d'attente
	err = channel.Publish(
		"",         // Exchange : n'envoie pas via un exchange
		queue.Name, // Routing key : nom de la file d'attente
		false,      // Mandatory : non
		false,      // Immediate : non
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Fatalf("Erreur lors de l'envoi du message : %s", err)
	}
	fmt.Printf(" [x] Sent %s\n", message)
}
