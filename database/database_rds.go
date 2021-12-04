package database

import (
	//"fmt"
	//"crypto/rand"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectToRDS() *sql.DB {
	db, err := sql.Open("mysql", "admin:retrobot.123@tcp(yt-recommendation.cepuilwl2joi.us-east-2.rds.amazonaws.com:3306)/memebot")

    if err != nil {
        panic(err.Error())
    }
	return db
}

func GetMemes(db *sql.DB, platform string)[]string{
	results, err := db.Query("SELECT url FROM resource where platform='"+platform+"' order by rand() limit 10;")
    if err != nil {
        panic(err.Error()) 
    }

	type wrapper struct{
		URL string `json:"url"`
	}
	var wrapperObj wrapper
	var meme_urls []string
    for results.Next() {
        err = results.Scan(&wrapperObj.URL)
        if err != nil {
            panic(err.Error()) 
        }
		meme_urls = append(meme_urls,wrapperObj.URL)
    }
	return meme_urls
}

