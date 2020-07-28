package gateway

func (p *postgresRepository) CreateUser(username, password, first, last, role string) error {
	_, err := p.client.db.Exec("INSERT INTO test1 (username, password, first, last, role) VALUES ($1, $2, $3, $4, $5)", username, string(password), first, last, role)
	if err != nil {
		return err
	}
	return nil
}
