// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package assets

import (
	"embed"
	"fmt"
	"io"

	"github.com/alimy/embedx"
	"github.com/alimy/hori/internal/utils"
	"github.com/aymerick/raymond/v2"
	"github.com/sirupsen/logrus"
)

//go:embed templates
var tmplFS embed.FS

type Template interface {
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
}

type raymondTmpl struct {
	tmpls map[string]*raymond.Template
}

func (t *raymondTmpl) ExecuteTemplate(w io.Writer, name string, data interface{}) error {
	tmpl, exist := t.tmpls[name]
	if !exist {
		return fmt.Errorf("template named %s is not exist", name)
	}
	result, err := tmpl.Exec(data)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(result))
	return err
}

// NewTemplate new template.Template instance from templates files.
func NewTemplate() Template {
	embedFS := embedx.ChangeRoot(tmplFS, "templates")
	raymond.RegisterNamer(raymond.NamerFunc(utils.Naming))
	if err := raymond.RegisterPartialFS(embedFS, "partials/*.hbs"); err != nil {
		logrus.Fatal(err)
	}
	filenames, err := utils.FileGlob(embedFS, "*.hbs", "account/*.hbs")
	if err != nil {
		logrus.Fatal(err)
	}
	tmpls := make(map[string]*raymond.Template, len(filenames))
	for _, filename := range filenames {
		tmpl, err := raymond.ParseWith(embedFS, filename)
		if err != nil {
			logrus.Fatal(err)
		}
		tmpls[utils.Naming(filename)] = tmpl
	}
	return &raymondTmpl{
		tmpls: tmpls,
	}
}
