package translate

import (
	"fmt"

	"github.com/Edgar-Avila/go-sqlmql/parser"
)

func TranslateLiteral(literal parser.Literal) (string, error) {
    v := literal.GetVal()
    switch literal.(type) {
        case parser.Int, parser.Decimal, parser.Bool:
            return fmt.Sprintf("%v", v), nil
        case parser.String:
            return fmt.Sprintf("\"%v\"", v), nil
        default: 
            return "", fmt.Errorf("Invalid literal type %T", v)
    }
}
