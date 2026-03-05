package repository

import (
	"JWT_REST_Gin_MySQL/configuration"
	"JWT_REST_Gin_MySQL/model"
)

func ListCategories() ([]model.MCategory, error) {
	db := configuration.DB
	rows, err := db.Query(`
		SELECT ID, name, created_at, updated_at
		FROM categories
		ORDER BY ID ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]model.MCategory, 0)
	for rows.Next() {
		var c model.MCategory
		if err := rows.Scan(&c.ID, &c.Name, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}
		res = append(res, c)
	}
	return res, rows.Err()
}

func CreateCategory(name string) (model.MCategory, error) {
	db := configuration.DB
	_, err := db.Exec(`INSERT INTO categories(name) VALUES (?)`, name)
	if err != nil {
		return model.MCategory{}, err
	}
	// 取最新插入的一条（简单做法）
	var c model.MCategory
	err = db.QueryRow(`
		SELECT ID, name, created_at, updated_at
		FROM categories
		WHERE name = ?
	`, name).Scan(&c.ID, &c.Name, &c.CreatedAt, &c.UpdatedAt)
	return c, err
}

func CategoryExists(id int64) (bool, error) {
	db := configuration.DB
	var x int
	err := db.QueryRow(`SELECT 1 FROM categories WHERE ID = ? LIMIT 1`, id).Scan(&x)
	if err != nil {
		// sql.ErrNoRows 也会到这里，你可以在 service 里统一处理为 false
		return false, err
	}
	return true, nil
}
