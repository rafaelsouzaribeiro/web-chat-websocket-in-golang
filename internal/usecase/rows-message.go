package usecase

func (l *MessageUsecase) GetMessageRows() (int64, error) {
	rows, err := l.Irepository.GetMessageRows()

	if err != nil {
		return 0, err
	}

	return rows, nil
}
