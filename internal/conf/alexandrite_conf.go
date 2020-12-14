// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package conf

import (
	"fmt"
)

// Config application config model
type Config struct {
	Application Application
	Server      Server
}

// Application indicate application section config
type Application struct {
	Name        string
	Authors     []string
	Description string
}

// Serve indicate server section config
type Server struct {
	Addr string
}

func (c *Application) String() string {
	return fmt.Sprintf("{name:%q, authors:%v,description:%q}", c.Name, c.Authors, c.Description)
}

func (c *Server) String() string {
	return fmt.Sprintf("{addr:%q}", c.Addr)
}

func (c *Config) String() string {
	return fmt.Sprintf("{application:%s, server:%s}", &c.Application, &c.Server)
}
