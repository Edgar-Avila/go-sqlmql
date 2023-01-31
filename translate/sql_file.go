package translate

import (
	"fmt"
	"strings"

	"github.com/Edgar-Avila/go-sqlmql/parser"
)


func TranslateStatement(statement parser.Statement) (string, error) {
    switch s := statement.(type) {
    case *parser.CreateStatement:
        return TranslateCreate(s)
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

func TranslateSqlFile(sqlFile *parser.SqlFile) (string, error) {
    translated := make([]string, 0)
	for _, statement := range sqlFile.Statements {
        t, err := TranslateStatement(statement)
        if err != nil {
            return "", err
        }
        translated = append(translated, t)
	}
	return strings.Join(translated, "\n"), nil
}
