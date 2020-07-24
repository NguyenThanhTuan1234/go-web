package gateway

import "go-web/models"

func (r *rdbClient) GetUser(un string) (*models.User, error) {
	var u *models.User

	rows, err := r.db.Query("SELECT username, password, first, last, role FROM test1 WHERE username = $1", un)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&u.UserName, &u.Password, &u.First, &u.Last, &u.Role)
		if err != nil {
			panic(err)
		}
	}
	return u, nil
}
