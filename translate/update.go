package translate

import (
	"fmt"
	"strings"

	"github.com/Edgar-Avila/go-sqlmql/parser"
)

func TranslateUpdate(updateStmt *parser.UpdateStatement) (string, error) {
	name := updateStmt.Name
	filter, err := TranslateWhere(updateStmt.Where)
	if err != nil {
		return "", err
	}
	changeArr := make([]string, len(updateStmt.Changes))
	for i, change := range updateStmt.Changes {
		lit, err := TranslateLiteral(change.Lit)
		if err != nil {
			return "", err
		}
		changeArr[i] = fmt.Sprintf("%s: %s", change.Name, lit)
	}
	update := fmt.Sprintf("{$set: {%s}}", strings.Join(changeArr, ","))
	translated := fmt.Sprintf("db.%s.updateMany(%s, %s)", name, filter, update)
	return translated, nil
}
