package repository

func (r *MesssageRepository) GetUsersRows() (float64, error) {

	pagination := r.getPagination("pagination_users")

	return float64(pagination.Page), nil

}
