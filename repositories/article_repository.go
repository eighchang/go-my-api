package repositories

import (
	"database/sql"

	"github.com/eighchang/goapi/database"
	"github.com/eighchang/goapi/models"
)

const (
	articleNumPerPage = 5
)

func GetArticle(articleID int) (*models.Article, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	const sqlStr = `
		SELECT article_id, title, contents, username, nice, created_at
		FROM public.articles
		WHERE article_id = $1;
	`

	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var article models.Article
	var createdTime sql.NullTime

	err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		return nil, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return &article, nil
}

func GetArticles(page int) ([]models.Article, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	const sqlStr = `
		SELECT article_id, title, contents, username, nice, created_at
		FROM public.articles
		limit $1 offset $2;
	`

	rows, err := db.Query(sqlStr, articleNumPerPage, ((page - 1) * articleNumPerPage))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	articles := make([]models.Article, 0)

	for rows.Next() {
		var article models.Article
		var createdTime sql.NullTime

		err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
		if err != nil {
			return nil, err
		}

		if createdTime.Valid {
			article.CreatedAt = createdTime.Time
		}

		articles = append(articles, article)
	}

	return articles, nil
}

func InsertArticle(tx *sql.Tx, article models.Article) (models.Article, error) {
	const sqlStr = `
		INSERT INTO public.articles
		(title, contents, username, nice, created_at)
		VALUES($1, $2, $3, $4, now())
		returning article_id, created_at
		;
	`
	var newArticle = models.Article{
		Title:    article.Title,
		Contents: article.Contents,
		UserName: article.UserName,
	}

	var newID int
	var createdTime sql.NullTime
	tx.QueryRow(sqlStr, article.Title, article.Contents, article.UserName).Scan(&newID, &createdTime)

	err := tx.Commit()
	if err != nil {
		return models.Article{}, err
	}

	newArticle.ID = newID

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return newArticle, nil
}

func UpdateNiceNum(tx *sql.Tx, articleID int) error {
	const sqlGetNice = `
		select nice
		from articles
		where article_id = $1;
	`
	row := tx.QueryRow(sqlGetNice, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}

	var nicenum int
	err := row.Scan(&nicenum)
	if err != nil {
		tx.Rollback()
		return err
	}

	const sqlUpdateNice = `update articles set nice = ? where article_id = $1`

	_, err = tx.Exec(sqlUpdateNice, nicenum)
	if err != nil {
		tx.Rollback()
		return nil
	}

	tx.Commit()
	return nil
}
