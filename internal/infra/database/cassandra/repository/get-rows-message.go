package repository

func (r *MesssageRepository) GetMessageRows() (int64, error) {
	pagination := r.getPagination("pagination_messages")

	return int64(pagination.Page), nil

}
