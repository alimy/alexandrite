// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package conf

import (
	"fmt"
	"net/url"
	"strings"
)

// Config application config model
type Config struct {
	Application Application
	Runtime     Runtime
	Server      Server
	Database    Database
	Cache       Cache
	Store       Store
}

// Application indicate application section config
type Application struct {
	Name        string
	Authors     []string
	Description string
}

// Runtime indicate runtime section config
type Runtime struct {
	RunMode      string `toml:"run_mode"`
	MockDatabase bool   `toml:"mock_database"`
	MockStore    bool   `toml:"mock_store"`
	MockCache    bool   `toml:"mock_cache"`
}

// Serve indicate server section config
type Server struct {
	Addr string
}

// Database indicate database section config
type Database struct {
	Type         string
	Host         string
	Name         string
	User         string
	Password     string
	Path         string
	SSLMode      string `toml:"ssl_mode"`
	MaxIdleConns int    `toml:"max_idle_conns"`
	MaxOpenConns int    `toml:"max_open_conns"`
}

// Cache indicate cache section config
type Cache struct {
	Type string
}

// Store indicate store section config
type Store struct {
	Type string
	Path string
}

func (c *Application) String() string {
	return fmt.Sprintf("{name:%q, authors:%v,description:%q}", c.Name, c.Authors, c.Description)
}

func (c *Server) String() string {
	return fmt.Sprintf("{addr:%q}", c.Addr)
}

func (c *Runtime) String() string {
	return fmt.Sprintf("{run_mode:%q, mock_database:%t, mock_store:%t, mock_cache:%t}", c.RunMode, c.MockDatabase, c.MockStore, c.MockCache)
}

func (c *Runtime) InProdMode() bool {
	return strings.ToLower(c.RunMode) == "prod"
}

func (c *Database) String() string {
	return fmt.Sprintf("{type:%q, host:%q, name:%q, user:%q, password:%q, path:%q, ssl_mode:%q, max_idle_conns:%d, max_open_conns:%d}",
		c.Type, c.Host, c.Name, c.User, c.Password, c.Path, c.SSLMode, c.MaxIdleConns, c.MaxOpenConns)
}

func (c *Database) Dsn() (dsn string, err error) {
	// In case the database name contains "?" with some parameters
	concate := "?"
	if strings.Contains(c.Name, concate) {
		concate = "&"
	}

	switch c.Type {
	case "mysql":
		if c.Host[0] == '/' { // Looks like a unix socket
			dsn = fmt.Sprintf("%s:%s@unix(%s)/%s%scharset=utf8mb4&parseTime=true",
				c.User, c.Password, c.Host, c.Name, concate)
		} else {
			dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s%scharset=utf8mb4&parseTime=true",
				c.User, c.Password, c.Host, c.Name, concate)
		}

	case "postgres":
		host, port := c.parsePostgreSQLHostPort(c.Host)
		if host[0] == '/' { // looks like a unix socket
			dsn = fmt.Sprintf("postgres://%s:%s@:%s/%s%ssslmode=%s&host=%s",
				url.QueryEscape(c.User), url.QueryEscape(c.Password), port, c.Name, concate, c.SSLMode, host)
		} else {
			dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s%ssslmode=%s",
				url.QueryEscape(c.User), url.QueryEscape(c.Password), host, port, c.Name, concate, c.SSLMode)
		}

	case "mssql":
		host, port := c.parseMSSQLHostPort(c.Host)
		dsn = fmt.Sprintf("server=%s; port=%s; database=%s; user id=%s; password=%s;",
			host, port, c.Name, c.User, c.Password)

	case "sqlite3":
		dsn = "file:" + c.Path + "?cache=shared&mode=rwc"

	default:
		return "", fmt.Errorf("unrecognized dialect: %s", c.Type)
	}

	return dsn, nil
}

// parsePostgreSQLHostPort parses given input in various forms defined in
// https://www.postgresql.org/docs/current/static/libpq-connect.html#LIBPQ-CONNSTRING
// and returns proper host and port number.
func (c *Database) parsePostgreSQLHostPort(info string) (host, port string) {
	host, port = "127.0.0.1", "5432"
	if strings.Contains(info, ":") && !strings.HasSuffix(info, "]") {
		idx := strings.LastIndex(info, ":")
		host = info[:idx]
		port = info[idx+1:]
	} else if len(info) > 0 {
		host = info
	}
	return host, port
}

func (c *Database) parseMSSQLHostPort(info string) (host, port string) {
	host, port = "127.0.0.1", "1433"
	if strings.Contains(info, ":") {
		host = strings.Split(info, ":")[0]
		port = strings.Split(info, ":")[1]
	} else if strings.Contains(info, ",") {
		host = strings.Split(info, ",")[0]
		port = strings.TrimSpace(strings.Split(info, ",")[1])
	} else if len(info) > 0 {
		host = info
	}
	return host, port
}

func (c *Cache) String() string {
	return fmt.Sprintf("{type:%q}", c.Type)
}

func (c *Store) String() string {
	return fmt.Sprintf("{type:%q, path:%q}", c.Type, c.Path)
}

func (c *Config) String() string {
	return fmt.Sprintf("{application:%s, runtime:%s, server:%s, database:%s, cache:%s, store:%s}",
		&c.Application, &c.Runtime, &c.Server, &c.Database, &c.Cache, &c.Store)
}
