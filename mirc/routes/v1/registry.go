package v1

import (
	. "github.com/alimy/mir/v2"
	. "github.com/alimy/mir/v2/engine"
)

func init() {
	AddEntry(new(Registry))
}

// Registry crates registry api define
type Registry struct {
	Group           Group  `mir:"/api/v1"`
	Register        Post   `mir:"account/register"`
	Login           Post   `mir:"account/login"`
	TokensInfo      Post   `mir:"account/tokens"`
	GenerateTokens  Put    `mir:"account/tokens"`
	RevokeTokens    Delete `mir:"account/tokens"`
	TokenByName     Get    `mir:"account/tokens/{name}"`
	Categories      Get    `mir:"categories"`
	SearchCrates    Get    `mir:"crates"`
	PublishCrates   Put    `mir:"crates/new"`
	SuggestCrates   Get    `mir:"crates/suggest"`
	CrateInfo       Get    `mir:"crates/{name}"`
	CrateOwnersInfo Get    `mir:"crates/{name}/owners"`
	PutCrateOwners  Put    `mir:"crates/{name}/owners"`
	DelCrateOwners  Delete `mir:"crates/{name}/owners"`
	YankCrate       Delete `mir:"crates/{name}/{version}/yank"`
	UnyankCrate     Put    `mir:"crates/{name}/{version}/unyank"`
	DownloadCrate   Get    `mir:"crates/{name}/{version}/download"`
}
