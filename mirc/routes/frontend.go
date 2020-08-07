// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package routes

import (
	. "github.com/alimy/mir/v2"
	. "github.com/alimy/mir/v2/engine"
)

func init() {
	AddEntry(new(Frontend))
}

// Frontend frontend api define
type Frontend struct {
	Chain          Chain `mir:"-"`
	Index          Get   `mir:"/"`
	Me             Get   `mir:"/me"`
	Search         Get   `mir:"/search"`
	MostDownloaded Get   `mir:"/most-downloaded"`
	LastUpdated    Get   `mir:"/last-updated"`
	Crate          Get   `mir:"/crates/{crate}"`
	Login          Get   `mir:"/account/login"`
	LoginPost      Post  `mir:"/account/login"`
	Logout         Get   `mir:"/account/logout"`
	Register       Get   `mir:"/account/register"`
	RegisterPost   Post  `mir:"/account/register"`
	Manage         Get   `mir:"/account/manage"`
	Password       Post  `mir:"/account/manage/password"`
	Tokens         Post  `mir:"/account/manage/tokens"`
	RevokeToken    Get   `mir:"/account/manage/tokens/{token-id}/revoke"`
	Assets         Get   `mir:"/assets/"`
}
