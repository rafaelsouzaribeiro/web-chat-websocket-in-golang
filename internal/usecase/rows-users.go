package usecase

func (l *MessageUsecase) GetUsersRows() (int64, error) {
	rows, err := l.Irepository.GetUsersRows()

	if err != nil {
		return 0, err
	}

	return rows, nil
}
