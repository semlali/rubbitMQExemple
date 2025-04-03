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

	// Déclare la même file d'attente que celle du producteur
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

	// Consommation des messages
	msgs, err := channel.Consume(
		queue.Name, // Nom de la file d'attente
		"",         // Consumer tag : vide pour un consommateur unique
		true,       // Auto-ack : reconnaitre automatiquement les messages
		false,      // Exclusive : ce consommateur peut-il être exclusif ?
		false,      // No-local : les messages envoyés par ce consommateur sont-ils ignorés ?
		false,      // No-wait : n'attendre aucune réponse
		nil,
	)
	if err != nil {
		log.Fatalf("Erreur lors de la consommation des messages : %s", err)
	}

	// Traitement des messages reçus
	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
	for msg := range msgs {
		fmt.Printf(" [x] Received %s\n", msg.Body)
	}
}
