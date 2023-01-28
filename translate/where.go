package translate

import (
	"fmt"
	"strings"

	"github.com/Edgar-Avila/go-sqlmql/parser"
)

// TranslateFactor translates a single boolean factor
func TranslateFactor(factor parser.BoolFactor, surround bool) (string, error) {
	translated := ""
	if surround {
		translated += "{"
	}
	if factor.Not {
		translated += "$not:{"
	}
	translated += factor.Cmp.Col
	translated += ":{"
	if factor.Cmp.Op == "=" {
		translated += "$eq:"
	} else if factor.Cmp.Op == "<" {
		translated += "$lt:"
	} else if factor.Cmp.Op == ">" {
		translated += "$gt:"
	} else if factor.Cmp.Op == "<=" {
		translated += "$lte:"
	} else if factor.Cmp.Op == ">=" {
		translated += "$gte:"
	}
    val, err := TranslateLiteral(factor.Cmp.Lit)
	if err != nil {
		return "", err
	}
	translated += string(val)
	translated += "}"
	if factor.Not {
		translated += "}"
	}
	if surround {
		translated += "}"
	}
	return translated, nil
}

// TranslateTerm translates a single boolean term
func TranslateTerm(term parser.BoolTerm, surround bool) (string, error) {
	translated := ""
	multipleFactors := len(term.Factors) > 1
	if surround {
		translated += "{"
	}
	if multipleFactors {
		translated += "$and:["
	}
	andArr := make([]string, 0)
	for _, factor := range term.Factors {
		factorStr, err := TranslateFactor(factor, multipleFactors)
		if err != nil {
			return "", err
		}
		andArr = append(andArr, factorStr)
	}
	translated += strings.Join(andArr, ",")
	if multipleFactors {
		translated += "]"
	}
	if surround {
		translated += "}"
	}
	return translated, nil
}

// TranslateWhere translates a where statement,
// take into account that it considers OR to have 
// a higher priority than AND 
func TranslateWhere(where *parser.Where) (string, error) {
	translated := ""
	findArr := make([]string, 0)
	if where != nil {
		multipleTerms := len(where.Terms) > 1
		obj := ""
		if multipleTerms {
			obj += "$or:["
		}
		orArr := make([]string, 0)
		for _, term := range where.Terms {
			termStr, err := TranslateTerm(term, multipleTerms)
			if err != nil {
				return "", err
			}
			orArr = append(orArr, termStr)
		}
		obj += strings.Join(orArr, ",")
		if multipleTerms {
			obj += "]"
		}
		findArr = append(findArr, obj)
	}
	translated = fmt.Sprintf("{%s}", strings.Join(findArr, ", "))

	return translated, nil
}
