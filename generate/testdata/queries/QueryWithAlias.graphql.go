package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"github.com/Khan/genqlient/graphql"
	"github.com/me/mypkg"
)

// QueryWithAliasResponse is returned by QueryWithAlias on success.
type QueryWithAliasResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User QueryWithAliasUser `json:"User"`
}

// QueryWithAliasUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type QueryWithAliasUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	ID mypkg.ID `json:"ID"`
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	OtherID mypkg.ID `json:"otherID"`
}

func QueryWithAlias(
	client graphql.Client,
) (*QueryWithAliasResponse, error) {
	var retval QueryWithAliasResponse
	err := client.MakeRequest(
		nil,
		"QueryWithAlias",
		`
query QueryWithAlias {
	User: user {
		ID: id
		otherID: id
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}
