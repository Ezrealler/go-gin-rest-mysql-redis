package repository

import (
	"JWT_REST_Gin_MySQL/configuration"
	"JWT_REST_Gin_MySQL/model"
	"database/sql"
	"errors"
	"log"

	// Use prefix blank identifier _ when importing driver for its side
	// effect and not use it explicity anywhere in our code.
	// When a package is imported prefixed with a blank identifier,the init
	// function of the package will be called. Also, the GO compiler will
	// not complain if the package is not used anywhere in the code

	_ "github.com/go-sql-driver/mysql"
)

// GetPostByID ...
func GetPostByID(id int64) (model.MPost, error) {
	db := configuration.DB

	var post model.MPost

	err := db.QueryRow(
		`SELECT p.ID, p.user_id, p.category_id, c.name, p.title, p.description, p.status, p.created_at, p.updated_at
		FROM posts p
		LEFT JOIN categories c ON p.category_id = c.ID
		WHERE p.ID = ?
		`, id).Scan(
		&post.ID,
		&post.UserID,
		&post.CategoryID,
		&post.Category,
		&post.Title,
		&post.Description,
		&post.Status,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return post, err
	}

	return post, nil
}

// 分页版GetPostAll ...
func ListPosts(offset, limit int, categoryID int64) ([]model.MPost, error) {
	db := configuration.DB

	var (
		rows *sql.Rows
		err  error
	)

	if categoryID > 0 {
		rows, err = db.Query(`
			SELECT p.ID, p.user_id, p.category_id, c.name, p.title, p.description, p.status, p.created_at, p.updated_at
			FROM posts p
			LEFT JOIN categories c ON p.category_id = c.ID
			WHERE p.category_id = ?
			ORDER BY p.created_at DESC
			LIMIT ? OFFSET ?
		`, categoryID, limit, offset)
	} else {
		rows, err = db.Query(`
			SELECT p.ID, p.user_id, p.category_id, c.name, p.title, p.description, p.status, p.created_at, p.updated_at
			FROM posts p
			LEFT JOIN categories c ON p.category_id = c.ID
			ORDER BY p.created_at DESC
			LIMIT ? OFFSET ?
		`, limit, offset)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := make([]model.MPost, 0, limit)
	for rows.Next() {
		var p model.MPost
		if err := rows.Scan(
			&p.ID, &p.UserID, &p.CategoryID, &p.Category,
			&p.Title, &p.Description, &p.Status, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, rows.Err()
}

// CreatePost ...
func CreatePost(mPost model.MPost) (model.MPost, error) {
	db := configuration.DB

	var err error

	if mPost.CategoryID == 0 {
		mPost.CategoryID = 1
	}

	crt, err := db.Prepare(`
		INSERT INTO posts (user_id, category_id, title, description, status)
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		log.Panic(err)
		return mPost, err
	}

	res, err := crt.Exec(mPost.UserID, mPost.CategoryID, mPost.Title, mPost.Description, mPost.Status)
	if err != nil {
		//log.Panic(err)
		return mPost, err
	}

	rowID, err := res.LastInsertId()
	if err != nil {
		log.Panic(err)
		return mPost, err
	}

	mPost.ID = int64(rowID)

	// find post by id
	resval, err := GetPostByID(mPost.ID)
	if err != nil {
		log.Panic(err)
		return mPost, err
	}

	return resval, nil
}

// UpdatePost ...
func UpdatePost(mPost model.MPost) (model.MPost, error) {
	db := configuration.DB

	var err error

	crt, err := db.Prepare("update posts set title =?, description =?, status =? where id=?")
	if err != nil {
		return mPost, err
	}
	_, queryError := crt.Exec(mPost.Title, mPost.Description, mPost.Status, mPost.ID)
	if queryError != nil {
		return mPost, queryError
	}

	// find post by id
	res, err := GetPostByID(mPost.ID)
	if err != nil {
		return mPost, err
	}

	return res, nil
}

// DeletePostByID ...
func DeletePostByID(id int64) error {
	db := configuration.DB

	res, err := db.Exec("DELETE FROM posts WHERE ID = ?", id)
	if err != nil {
		return err
	}

	aff, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if aff == 0 {
		return errors.New("post not found")
	}

	return nil
}
