package newsapi

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewConnection(databaseConfig *Config) (*sql.DB, error) {
	// This function creates a new connection to the database
	// full dsn- username:password@protocol(address)/dbname?param=value
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		databaseConfig.Database.User, databaseConfig.Database.Password, databaseConfig.Database.Host, databaseConfig.Database.Port, databaseConfig.Database.Database))
	return db, err
}

func SaveNews(db *sql.DB, news *NewsResponse, table string) (int64, error) {
	// "INSERT IGNORE INTO news(uniqId, sourceid, sourcename, author, title, description, url, urlToImage, publishedAt, content) VALUES(?,?,?,?,?,?,?,?,?,?)", (sourceId+publishedAt+author+title, sourceId, sourceName, author, title, description, url, urlToImage, publishedAt, content)
	rawSqlStmt := "INSERT IGNORE INTO " + table + "(uniqId, sourceid, sourcename, author, title, description, url, urlToImage, publishedAt, content) VALUES "
	vals := []interface{}{}
	for _, rows := range news.Articles {
		rawSqlStmt += "(?,?,?,?,?,?,?,?,?,?),"
		vals = append(vals, rows.Source.ID+rows.PublishedAt+rows.Author+rows.Title, rows.Source.ID, rows.Source.Name, rows.Author, rows.Title, rows.Description, rows.URL, rows.URLToImage, rows.PublishedAt, rows.Content)
	}

	rawSqlStmt = rawSqlStmt[0:len(rawSqlStmt)-1] + ";"
	prepStmt, err := db.Prepare(rawSqlStmt)
	if err != nil {
		return 0, err
	}
	res, err := prepStmt.Exec(vals...)
	if err != nil {
		return 0, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return 0, nil // doesn't matter, it's only for logging purposes
	}
	defer prepStmt.Close()

	return rows, nil
}
