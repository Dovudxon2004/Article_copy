package postgres

import (
	"database/sql"

	"github.com/Dovudxon2004/Api/Article/models"
	"github.com/Dovudxon2004/Api/Article/storage"
	"github.com/jmoiron/sqlx"
)

type articleRepo struct {
	db *sqlx.DB
}

func NewArticleRepo(db *sqlx.DB) storage.ArticleRepoI {
	return articleRepo{
		db: db,
	}
}

func (r articleRepo) Create(entity models.ArticleCreateModel) (err error) {
	insertQuery := `INSERT INTO article (
		title,
		body,
		author_id
	) VALUES (
		$1,
		$2,
		$3
	)`

	_, err = r.db.Exec(insertQuery, entity.Title, entity.Body, entity.AuthorID)

	return err
}

func (r articleRepo) GetList(query models.Query) (resp []models.ArticleListItem, err error) {
	var rows *sql.Rows

	if len(query.Search) > 0 {
		rows, err = r.db.Query(
			`SELECT
			ar.id, ar.title, ar.body, ar.created_at, ar.updated_at,
			au.id, au.firstname, au.lastname, au.created_at, au.updated_at
			FROM article AS ar JOIN author AS au ON ar.author_id = au.id
			WHERE title ILIKE '%' || $3 || '%'
			OFFSET $1 LIMIT $2`,
			query.Offset,
			query.Limit,
			query.Search,
		)
	} else {
		rows, err = r.db.Query(`SELECT
			ar.id, ar.title, ar.body, ar.created_at, ar.updated_at,
			au.id, au.firstname, au.lastname, au.created_at, au.updated_at
			FROM article AS ar JOIN author AS au ON ar.author_id = au.id
			OFFSET $1 LIMIT $2`,
			query.Offset,
			query.Limit,
		)
	}

	if err != nil {
		return resp, err
	}

	defer rows.Close()

	for rows.Next() {
		var a models.ArticleListItem

		err = rows.Scan(
			&a.ID, &a.Title, &a.Body, &a.CreatedAt, &a.UpdatedAt,
			&a.Author.ID, &a.Author.Firstname, &a.Author.Lastname, &a.Author.CreatedAt, &a.Author.UpdatedAt,
		)
		resp = append(resp, a)
		if err != nil {
			return resp, err
		}

	}
	return resp, err
}

func (r articleRepo) GetByID(ID int) (resp models.Article, err error) {

	return
}

func (r articleRepo) Update(entity models.ArticleUpdateModel) (effectedRowsNum int, err error) {

	return
}

func (r articleRepo) Delete(ID int) (effectedRowsNum int, err error) {

	return
}
