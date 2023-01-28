package translate

import (
	"fmt"

	"github.com/Edgar-Avila/go-sqlmql/parser"
)

func TranslateCreate(createStmt *parser.CreateStatement) (string, error) {
    return fmt.Sprintf("db.createCollection(\"%s\");", createStmt.Name), nil
}
