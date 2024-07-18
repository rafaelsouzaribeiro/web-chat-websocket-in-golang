package usecase

import "github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"

type MessageUsecase struct {
	Irepository entity.Irepository
}

func NewMessageUseCase(IRepository entity.Irepository) *MessageUsecase {
	return &MessageUsecase{
		Irepository: IRepository,
	}
}
