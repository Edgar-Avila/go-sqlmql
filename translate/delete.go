package translate

import (
	"fmt"

	"github.com/Edgar-Avila/go-sqlmql/parser"
)

func TranslateDelete(deleteStmt *parser.DeleteStatement) (string, error) {
	name := deleteStmt.Name
	filter, err := TranslateWhere(deleteStmt.Where)
	if err != nil {
		return "", err
	}
	translated := fmt.Sprintf("db.%s.deleteMany(%s);", name, filter)
	return translated, nil
}
