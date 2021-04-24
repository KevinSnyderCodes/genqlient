package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"encoding/json"

	"github.com/Khan/genqlient/graphql"
)

// UnionNoFragmentsQueryRandomLeafArticle includes the requested fields of the GraphQL type Article.
type UnionNoFragmentsQueryRandomLeafArticle struct {
	Typename string `json:"__typename"`
}

func (v UnionNoFragmentsQueryRandomLeafArticle) implementsGraphQLInterfaceUnionNoFragmentsQueryRandomLeafLeafContent() {
}

// UnionNoFragmentsQueryRandomLeafLeafContent includes the requested fields of the GraphQL type LeafContent.
// The GraphQL type's documentation follows.
//
// LeafContent represents content items that can't have child-nodes.
type UnionNoFragmentsQueryRandomLeafLeafContent interface {
	implementsGraphQLInterfaceUnionNoFragmentsQueryRandomLeafLeafContent()
}

// UnionNoFragmentsQueryRandomLeafVideo includes the requested fields of the GraphQL type Video.
type UnionNoFragmentsQueryRandomLeafVideo struct {
	Typename string `json:"__typename"`
}

func (v UnionNoFragmentsQueryRandomLeafVideo) implementsGraphQLInterfaceUnionNoFragmentsQueryRandomLeafLeafContent() {
}

// UnionNoFragmentsQueryResponse is returned by UnionNoFragmentsQuery on success.
type UnionNoFragmentsQueryResponse struct {
	RandomLeaf UnionNoFragmentsQueryRandomLeafLeafContent `json:"-"`
}

func (v *UnionNoFragmentsQueryResponse) UnmarshalJSON(b []byte) error {
	var firstPass struct {
		*UnionNoFragmentsQueryResponse
		RandomLeaf json.RawMessage `json:"randomLeaf"`
	}
	firstPass.UnionNoFragmentsQueryResponse = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err = json.Unmarshal(firstPass.RandomLeaf, &tn)
	if err != nil {
		return err
	}
	switch tn.TypeName {

	case "Article":

		v.RandomLeaf = UnionNoFragmentsQueryRandomLeafArticle{}
		err = json.Unmarshal(
			firstPass.RandomLeaf, &v.RandomLeaf)

	case "Video":

		v.RandomLeaf = UnionNoFragmentsQueryRandomLeafVideo{}
		err = json.Unmarshal(
			firstPass.RandomLeaf, &v.RandomLeaf)

	}
	if err != nil {
		return err
	}

	return nil
}

func UnionNoFragmentsQuery(
	client graphql.Client,
) (*UnionNoFragmentsQueryResponse, error) {
	var retval UnionNoFragmentsQueryResponse
	err := client.MakeRequest(
		nil,
		"UnionNoFragmentsQuery",
		`
query UnionNoFragmentsQuery {
	randomLeaf {
		__typename
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}
