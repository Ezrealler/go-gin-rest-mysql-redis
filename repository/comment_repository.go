package repository

import (
	"JWT_REST_Gin_MySQL/configuration"
	"JWT_REST_Gin_MySQL/model"
)

// CreateComment 往 comments 表插入一条评论
func CreateComment(postID, userID int64, content string) (model.MComment, error) {
	db := configuration.DB

	res, err := db.Exec(`
		INSERT INTO comments (post_id, user_id, content, created_at)
		VALUES (?, ?, ?, NOW())
	`, postID, userID, content)
	if err != nil {
		return model.MComment{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return model.MComment{}, err
	}

	return GetCommentByID(id)
}

func GetCommentByID(id int64) (model.MComment, error) {
	db := configuration.DB

	var cmt model.MComment
	err := db.QueryRow(`
		SELECT ID, post_id, user_id, content, created_at
		FROM comments
		WHERE ID = ?
	`, id).Scan(&cmt.ID, &cmt.PostID, &cmt.UserID, &cmt.Content, &cmt.CreatedAt)

	return cmt, err
}

// ListCommentsByPostID 查询某篇文章的评论列表
func ListCommentsByPostID(postID int64, limit, offset int) ([]model.MComment, error) {
	db := configuration.DB

	rows, err := db.Query(`
		SELECT ID, post_id, user_id, content, created_at
		FROM comments
		WHERE post_id = ?
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?
	`, postID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := make([]model.MComment, 0, limit)
	for rows.Next() {
		var c model.MComment
		if err := rows.Scan(&c.ID, &c.PostID, &c.UserID, &c.Content, &c.CreatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, rows.Err()
}
