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
        case *parser.InsertStatement:
            return TranslateInsert(s)
        case *parser.SelectStatement:
            return TranslateSelect(s)
        case *parser.UpdateStatement:
            return TranslateUpdate(s)
		default:
			return "", fmt.Errorf("Invalid statement type %T", s)
		}
	}
	return "", nil
}
