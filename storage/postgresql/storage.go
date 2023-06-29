package storage


import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// type user_sites struct {
// 	chatid int64
// 	site   string
// }

type DB struct {
	DB *pgxpool.Pool
}

func NewDB(db *pgxpool.Pool) *DB {
	return &DB{db}
}

func (d *DB) Save(chatID int64, site string) error {
	query := fmt.Sprintf("insert into user_sites values(%v, '%s')", chatID, site)
	_, err := d.DB.Query(context.Background(), query)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) Get(chatID int64) ([]string, error) {
	query := fmt.Sprintf("select site from user_sites where chatid=%v", chatID)
	rows, err := d.DB.Query(context.Background(), query)
	if err != nil {
		return []string{}, err
	}

	var webSites []string
	for rows.Next() {
		var webSite string

		err = rows.Scan(&webSite)
		if err != nil {
			return []string{}, err
		}

		webSites = append(webSites, webSite)
	}

	return webSites, nil
}

func (d *DB) Delete(chatID int64, site string) error {
	query := fmt.Sprintf("delete from user_sites where chatid=%v and site='%s'", chatID, site)

	_, err := d.DB.Query(context.Background(), query)
	if err != nil {
		return err
	}

	return nil
}
