package usecase

func (l *MessageUsecase) GetMessageRows() (float64, error) {
	rows, err := l.Irepository.GetMessageRows()

	if err != nil {
		return 0, err
	}

	return rows, nil
}
