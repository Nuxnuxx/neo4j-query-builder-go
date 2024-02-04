package querybuilder

import ()

type QueryBuilder interface {
	Build() string
}

type Query struct {
	queryParts []QueryBuilder
	lastPart   string
}

func NewQuery() *Query {
	return &Query{queryParts: []QueryBuilder{}, lastPart: ""}
}
