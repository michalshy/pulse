package store

import "strings"

type filter struct {
	clauses []string
	params  []any
}

func (f *filter) add(clause string, param any) {
	f.clauses = append(f.clauses, clause)
	f.params = append(f.params, param)
}

func (f *filter) where() (string, []any) {
	if len(f.clauses) == 0 {
		return "", nil
	}
	return "WHERE " + strings.Join(f.clauses, " AND "), f.params
}
