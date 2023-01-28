package translate

import (
	"fmt"
	"github.com/Edgar-Avila/go-sqlmql/parser"
	"strings"
)

// TranslateProjection translates the projection 
// of a select statement
func TranslateProjection(cols []string) (string, error) {
	projArr := make([]string, 0)
	if cols[0] != "*" {
		projArr = append(projArr, "_id: 0")
		for _, col := range cols {
			projArr = append(projArr, fmt.Sprintf("%s: 1", col))
		}
	}
	return fmt.Sprintf("{%s}", strings.Join(projArr, ",")), nil
}

// TranslateSort translates the sort specification of an 
// ORDER BY clause
func TranslateSort(orderBy []parser.SortSpec) (string, error) {
	if len(orderBy) > 0 {
		sortArr := make([]string, 0)
		for _, spec := range orderBy {
			dir := 1
			if spec.Dir == "DESC" {
				dir = -1
			}
			sortArr = append(sortArr, fmt.Sprintf("%s: %v", spec.Col, dir))
		}
		return fmt.Sprintf(".sort({%s})", strings.Join(sortArr, ", ")), nil
	}
	return "", nil
}

// TranslateLimit translates a limit clause
func TranslateLimit(limit int64) (string, error) {
	if limit > 0 {
		return fmt.Sprintf(".limit(%v)", limit), nil
	}
	return "", nil
}

// TranslateSelect translates a select statement
func TranslateSelect(selectStmt *parser.SelectStatement) (string, error) {
	// db.students.find({}, {_id: 0, name: 1}).sort({name: -1})
	name := selectStmt.From
	proj, err := TranslateProjection(selectStmt.Cols)
	if err != nil {
		return "", err
	}
	find, err := TranslateWhere(selectStmt.Where)
	if err != nil {
		return "", err
	}
	sort, err := TranslateSort(selectStmt.OrderBy)
	if err != nil {
		return "", err
	}
	limit, err := TranslateLimit(selectStmt.Limit)
	if err != nil {
		return "", err
	}
	translated := fmt.Sprintf("db.%s.find(%s, %s)%s%s;", name, find, proj, sort, limit)
	return translated, nil
}
