package translate

import (
	"fmt"
	"strings"

	"github.com/Edgar-Avila/go-sqlmql/parser"
)

// TranslateInsert translates an insert statements returning a
// string with the translated output
func TranslateInsert(insertStmt *parser.InsertStatement) (string, error) {
	translated := fmt.Sprintf("db.%s.insertMany([", insertStmt.Into)
	docs := make([]string, len(insertStmt.Rows))
	for i, row := range insertStmt.Rows {
		var doc string
		var err error
		if insertStmt.Cols == nil {
			keys := make([]string, len(row.Values))
			for i := range row.Values {
				keys[i] = "<field>"
			}
			doc, err = makeDoc(keys, row.Values, 0, TranslateLiteral)
		} else {
			doc, err = makeDoc(insertStmt.Cols, row.Values, 0, TranslateLiteral)
		}
		if err != nil {
			return "", err
		}
		docs[i] = doc
	}
	translated += strings.Join(docs, ",\n")
	translated += "]"

	return translated, nil
}
