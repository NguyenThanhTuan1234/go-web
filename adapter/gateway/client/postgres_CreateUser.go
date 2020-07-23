package client

func (r *rdbClient) CreateUser(username, first, last, role string, password []byte) error {
	_, err := r.db.Exec("INSERT INTO test1 (username, password, first, last, role) VALUES ($1, $2, $3, $4, $5)", username, string(password), first, last, role)
	if err != nil {
		return err
	}
	return nil
}
