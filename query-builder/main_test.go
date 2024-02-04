package querybuilder

import (
	"testing"
)

func TestSimpleMatch(t *testing.T) {
	query := NewQuery().Match("n:Recipe").ReturnField("n").BuildQuery()

	if query != "MATCH (n:Recipe) RETURN n" {
		t.Error("Expected MATCH (n:Recipe) RETURN n, got", query)
	}
}

func TestSimpleMatchWithSkipAndLimit(t *testing.T) {
	query := NewQuery().Match("n:Recipe").ReturnField("n").Skip("1").Limit("2").BuildQuery()

	if query != "MATCH (n:Recipe) RETURN n SKIP 1 LIMIT 2" {
		t.Error("Expected MATCH (n:Recipe) RETURN n SKIP 1 LIMIT 2, got", query)
	}
}

func TestSimpleMatchWithWhere(t *testing.T) {
	query := NewQuery().Match("n:Recipe").Where("n.name = 'Pasta'").ReturnField("n").BuildQuery()

	if query != "MATCH (n:Recipe) WHERE n.name = 'Pasta' RETURN n" {
		t.Error("Expected MATCH (n:Recipe) WHERE n.name = 'Pasta' RETURN n, got", query)
	}
}

func TestSimpleMatchWrongOrder(t *testing.T) {
	// Attempt to build a query with parts out of order
	defer func() {
		if r := recover(); r != nil {
			expectedErrorMsg := "RETURN must follow WHERE"
			if r.(error).Error() != expectedErrorMsg {
				t.Errorf("Expected panic with message '%s', got '%s'", expectedErrorMsg, r.(error).Error())
			}
		} else {
			t.Error("Expected panic, got nil")
		}
	}()

	query := NewQuery().Match("n:Recipe").ReturnField("n").Skip("1").Where("n.name = 'Pasta'").BuildQuery()
	_ = query
}
