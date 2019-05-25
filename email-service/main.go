package main
import (
	userPb "github.com/howl-io/shippy/user-service/proto/user"
	"github.com/micro/go-micro"
	"log"
	"context"
)

const (
	TOPIC = "user.created"
)

type Subscriber struct {}

func main() {
	server := micro.NewService(
				micro.Name("go.micro.srv.email"),
				micro.Version("latest"),
			)
	server.Init()

	micro.RegisterSubscriber(TOPIC, server.Server(), new(Subscriber))


	if err := server.Run(); err != nil {
		log.Fatalf("Server run failed: %v\n", err)
	}
}

func (sub *Subscriber) Process(ctx context.Context, user *userPb.User) error {
	log.Printf("[Picked up a new message]\n")
	log.Printf("[Sending email to]: %s\n", user.Name)
	return nil
}

func sendEmail(user *userPb.User) error {
	log.Printf("[SENDING A EMAIL to %s...]\n", user.Name)
	return nil
}

