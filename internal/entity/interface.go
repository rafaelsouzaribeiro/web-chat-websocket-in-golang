package entity

type Irepository interface {
	GetFromMessageIndex() (*[]Message, error)
	GetFromUsersIndex() (*[]Message, error)
	SaveMessage(msg *Message) error
	SaveUsers(msg *Message) error
	GetInitMessages() (*[]Message, error)
	GetInitUsers() (*[]Message, error)
	GetMessageRows() (int64, error)
	GetUsersRows() (int64, error)
}
