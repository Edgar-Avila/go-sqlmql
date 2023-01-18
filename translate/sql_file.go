package translate

import (
	"fmt"

	"github.com/Edgar-Avila/go-sqlmql/parser"
)

func TranslateSqlFile(sqlFile *parser.SqlFile) (string, error) {
	for _, statement := range sqlFile.Statements {
		switch s := statement.(type) {
		case *parser.DropStatement:
			return TranslateDrop(s)
		default:
			return "", fmt.Errorf("Invalid statement type %T", s)
		}
	}
	return "", nil
}
