// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package templates

import "embed"

// TmplFS export an embed templates FS
//go:embed account partials
//go:embed crate.hbs error.hbs index.hbs last-updated.hbs most-downloaded.hbs search.hbs
var TmplFS embed.FS
