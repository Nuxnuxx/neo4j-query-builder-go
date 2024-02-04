package neo4jquerybuildergo

import (
	"fmt"
	"strings"
)

type QueryBuilder interface {
	Build() string
}

type Query struct {
	queryParts []QueryBuilder
}

func NewQuery() *Query {
	return &Query{queryParts: []QueryBuilder{}}
}

func (q *Query) BuildQuery() string {
	var sb strings.Builder
	for _, part := range q.queryParts {
		sb.WriteString(part.Build())
		sb.WriteString(" ")
	}
	return strings.TrimSpace(sb.String())
}

func (q *Query) MatchNode(nodePattern string) *Query {
	q.queryParts = append(q.queryParts, &MatchBuilder{nodePattern: nodePattern})
	return q
}

func (q *Query) Where(condition string) *Query {
	q.queryParts = append(q.queryParts, &WhereBuilder{conditions: []string{condition}})
	return q
}

func (q *Query) ReturnField(field string) *Query {
	q.queryParts = append(q.queryParts, &ReturnBuilder{returnFields: []string{field}})
	return q
}

func (q *Query) Skip(skip int) *Query {
	q.queryParts = append(q.queryParts, &Skip{skip: skip})
	return q
}

func (q *Query) Limit(limit int) *Query {
	q.queryParts = append(q.queryParts, &Limit{limit: limit})
	return q
}

type MatchBuilder struct {
	nodePattern string
}

func (mb *MatchBuilder) Build() string {
	return "MATCH " + mb.nodePattern
}

type WhereBuilder struct {
	conditions []string
}

func (wb *WhereBuilder) Build() string {
	if len(wb.conditions) == 0 {
		return ""
	}
	return "WHERE " + strings.Join(wb.conditions, " AND ")
}

type ReturnBuilder struct {
	returnFields []string
}

func (rb *ReturnBuilder) Build() string {
	if len(rb.returnFields) == 0 {
		return ""
	}
	return "RETURN " + strings.Join(rb.returnFields, ", ")
}

type Skip struct {
	skip int
}

func (sb *Skip) Build() string {
	return fmt.Sprintf("SKIP %d", sb.skip)
}

type Limit struct {
	limit int
}

func (lb *Limit) Build() string {
	return fmt.Sprintf("LIMIT %d", lb.limit)
}
