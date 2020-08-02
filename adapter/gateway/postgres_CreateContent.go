package gateway

func (p *postgresRepository) CreateContent(content, title string, uid int) error {
	_, err := p.client.db.Exec("INSERT INTO content1(content,title,uid) VALUES($1, $2, $3)", content, title, uid)
	if err != nil {
		return err
	}
	return nil
}
