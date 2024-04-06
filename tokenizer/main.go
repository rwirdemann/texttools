package main

import (
	"fmt"
	"strings"
)

type Expectation struct {
	tokens      []string
	ignoreDiffs []int
}

func NewExpectation(expectation string, verification string) Expectation {
	e := Expectation{tokens: tokenize(expectation), ignoreDiffs: buildDiff(expectation, verification)}
	return e
}

func (e Expectation) Equal(actual string) bool {
	equal := true
	actualTokens := tokenize(actual)
	for i, v := range e.tokens {
		if v != actualTokens[i] {
			fmt.Printf("Token %d differs: %s / %s...", i, v, actualTokens[i])
			if contains(e.ignoreDiffs, i) {
				fmt.Printf("thats OK\n")
			} else {
				fmt.Printf("thats NOT OK\n")
				equal = false
			}
		}
	}
	return equal
}

func main() {
	e := "insert into job (description, publish_at, publish_trials, published_timestamp, tags, title, id) values ('World', '2024-04-02 08:37:37', 0, null, '', 'Hello', 39)"
	v := "insert into job (description, publish_at, publish_trials, published_timestamp, tags, title, id) values ('World', '2024-04-02 08:37:38', 0, null, '', 'Hello', 40)"

	expectation := NewExpectation(e, v)
	a1 := "insert into job (description, publish_at, publish_trials, published_timestamp, tags, title, id) values ('World', '2024-04-02 08:37:39', 0, null, '', 'Hello', 41)"
	a2 := "update job (description, publish_at, publish_trials, published_timestamp, tags, title, id) values ('World', '2024-04-02 08:37:39', 0, null, '', 'Hello', 41)"
	expectation.Equal(a1)
	expectation.Equal(a2)
}

func contains(values []int, value int) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

func buildDiff(expectation, verification string) []int {
	t1 := tokenize(expectation)
	t2 := tokenize(verification)
	diffs := []int{}
	for i, v := range t1 {
		if v != t2[i] {
			diffs = append(diffs, i)
		}
	}
	return diffs
}

func tokenize(s string) []string {
	split := strings.Split(s, ",")
	var tokens = []string{}
	for _, v := range split {
		tokens = append(tokens, strings.Trim(v, " )"))
	}

	return tokens
}
