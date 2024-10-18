package repository

func (r *MesssageRepository) GetMessageRows() (float64, error) {
	pagination := r.getPagination("pagination_messages")

	return float64(pagination.Page), nil

}
