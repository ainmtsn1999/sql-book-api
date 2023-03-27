package query

const (
	AddBook = `
		INSERT INTO
			db_go_sql.books
		(
			title,
			author,
			"desc"
		)
		VALUES ($1, $2, $3)
		RETURNING *;
	`

	UpdateBook = `
		UPDATE
			db_go_sql.books
		SET
			title = $2,
			author = $3,
			"desc" = $4
		WHERE id = $1;
	`

	GetAllBook = `
		SELECT *
		from db_go_sql.books
	`

	GetBookById = `
		SELECT *
		from db_go_sql.books
		WHERE id = $1;
	`

	DeleteBook = `
		DELETE from db_go_sql.books
		WHERE id = $1;
	`
)
