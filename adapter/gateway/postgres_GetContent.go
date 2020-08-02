package gateway

import "go-web/models"

func (p *postgresRepository) GetContent() (*models.Content, error) {
	var c models.Content

	rows, err := p.client.db.Query("SELECT title, content, uid FROM content1")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&c.Title, &c.Content, &c.Uid)
		if err != nil {
			panic(err)
		}
	}
	return &c, nil
}
