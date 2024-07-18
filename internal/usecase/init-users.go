package usecase

import (
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/usecase/dto"
)

func (l *MessageUsecase) GetInitusers() (*[]dto.Payload, error) {
	list, err := l.Irepository.GetInitUsers()

	if err != nil {
		return nil, err
	}

	var payloads []dto.Payload
	for _, v := range *list {
		payloads = append(payloads, dto.Payload{
			Message:  v.Message,
			Username: v.Username,
			Type:     v.Type,
			Time:     v.Time,
		})

	}

	return &payloads, nil
}
