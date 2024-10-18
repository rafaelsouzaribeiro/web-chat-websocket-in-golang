package entity

type Irepository interface {
	GetFromMessageIndex() (*[]Message, error)
	GetFromUsersIndex() (*[]Message, error)
	SaveMessage(msg *Message) error
	SaveUsers(msg *Message) error
	GetInitMessages() (*[]Message, error)
	GetInitUsers() (*[]Message, error)
	GetMessageRows() (float64, error)
	GetUsersRows() (float64, error)
}
