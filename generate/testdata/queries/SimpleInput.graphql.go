package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"github.com/Khan/genqlient/graphql"
	"github.com/me/mypkg"
)

// SimpleInputQueryResponse is returned by SimpleInputQuery on success.
type SimpleInputQueryResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User SimpleInputQueryUser `json:"user"`
}

// SimpleInputQueryUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type SimpleInputQueryUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id mypkg.ID `json:"id"`
}

func SimpleInputQuery(
	client graphql.Client,
	name string,
) (*SimpleInputQueryResponse, error) {
	variables := map[string]interface{}{
		"name": name,
	}

	var retval SimpleInputQueryResponse
	err := client.MakeRequest(
		nil,
		"SimpleInputQuery",
		`
query SimpleInputQuery ($name: String!) {
	user(query: {name:$name}) {
		id
	}
}
`,
		&retval,
		variables,
	)
	return &retval, err
}
