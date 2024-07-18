package usecase

import (
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/usecase/dto"
)

func (l *MessageUsecase) SaveUsers(input dto.Payload) (*dto.Payload, error) {

	err := l.Irepository.SaveUsers(&entity.Message{
		Message:  input.Message,
		Username: input.Username,
		Type:     input.Type,
		Time:     input.Time,
	})

	if err != nil {
		return nil, err
	}

	return &input, nil
}
