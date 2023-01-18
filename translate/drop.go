package translate

import (
	"fmt"

	"github.com/Edgar-Avila/go-sqlmql/parser"
)

func TranslateDrop(dropStmt *parser.DropStatement) (string, error) {
	translated := ""
	for _, name := range dropStmt.Tables {
		translated += fmt.Sprintf("db.%s.drop();\n", name)
	}
	return translated, nil
}
