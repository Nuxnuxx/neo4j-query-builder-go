package querybuilder

import (
	"errors"
	"fmt"
	"strings"
)

func (q *Query) Match(nodePattern string) *Query {
	q.queryParts = append(q.queryParts, &MatchBuilder{nodePattern: fmt.Sprintf("(%s)", nodePattern)})
	q.lastPart = "MATCH"
	return q
}

func (q *Query) Where(condition string) *Query {
	if q.lastPart != "MATCH" && q.lastPart != "WHERE" {
		panic(errors.New("WHERE must follow MATCH"))
	}
	q.queryParts = append(q.queryParts, &WhereBuilder{conditions: []string{condition}})
	q.lastPart = "WHERE"
	return q
}

func (q *Query) ReturnField(field string) *Query {
	if q.lastPart != "WHERE" && q.lastPart != "RETURN" {
		panic(errors.New("RETURN must follow WHERE"))
	}
	q.queryParts = append(q.queryParts, &ReturnBuilder{returnFields: []string{field}})
	q.lastPart = "RETURN"
	return q
}

func (q *Query) Skip(skip string) *Query {
	if q.lastPart != "RETURN" && q.lastPart != "SKIP" {
		panic(errors.New("SKIP must follow RETURN"))
	}
	q.queryParts = append(q.queryParts, &Skip{skip: skip})
	q.lastPart = "SKIP"
	return q
}

func (q *Query) Limit(limit string) *Query {
	if q.lastPart != "SKIP" && q.lastPart != "LIMIT" {
		panic(errors.New("LIMIT must follow SKIP"))
	}
	q.queryParts = append(q.queryParts, &Limit{limit: limit})
	q.lastPart = "LIMIT"
	return q
}

func (q *Query) BuildQuery() string {
	var sb strings.Builder
	for _, part := range q.queryParts {
		sb.WriteString(part.Build())
		sb.WriteString(" ")
	}
	return strings.TrimSpace(sb.String())
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
	skip string
}

func (sb *Skip) Build() string {
	return fmt.Sprintf("SKIP %s", sb.skip)
}

type Limit struct {
	limit string
}

func (lb *Limit) Build() string {
	return fmt.Sprintf("LIMIT %s", lb.limit)
}

