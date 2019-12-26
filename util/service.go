package util

import (
	"database/sql"
	"errors"
	"github.com/go-redis/redis"
	"github.com/sintanial/go-ip2location"
	"net/url"
	"strconv"
	"strings"
)

func NewMysqlClient(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn+"?parseTime=true&loc=Europe%2FMoscow&charset=utf8mb4")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func NewIP2LocationClient(fpath string, cached bool) (*ip2location.Reader, error) {
	return ip2location.FromFile(fpath, cached)
}

func NewRedisClient(dsn string) (*redis.Client, error) {
	u, err := url.Parse(dsn)
	if err != nil {
		return nil, errors.New("failed parse redis dsn: " + err.Error())
	}

	opts := &redis.Options{
		Network:  u.Scheme,
		Addr:     u.Host,
		Password: u.Query().Get("requirepass"),
	}

	if u.Path != "" {
		dbNum, err := strconv.Atoi(strings.Trim(u.Path, "/"))
		if err != nil {
			return nil, errors.New("invalid database format: " + u.Path)
		}

		opts.DB = dbNum
	}

	client := redis.NewClient(opts)
	return client, nil
}
