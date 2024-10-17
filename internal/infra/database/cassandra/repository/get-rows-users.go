package repository

func (r *MesssageRepository) GetUsersRows() (int64, error) {

	pagination := r.getPagination("pagination_users")

	return int64(pagination.Page), nil

}
