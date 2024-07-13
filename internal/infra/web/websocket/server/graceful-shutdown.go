package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/usecase/dto"
)

func gracefulShutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Error shutting down server: %v\n", err)
	}

	for _, user := range users {
		disconnectionMessage := dto.Payload{
			Username: "<strong>info</strong>",
			Message:  fmt.Sprintf("User <strong>%s</strong> disconnected", user.username),
		}

		saveMessageToRedis(disconnectionMessage, "users")
		deleteUserByConn(user.conn, true)
	}

	fmt.Println("Server gracefully stopped")

	os.Exit(0)
}
