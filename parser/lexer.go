package parser

import (
	"github.com/alecthomas/participle/v2/lexer"
)

func newLexer() (lexer.Definition, error) {
	return lexer.NewSimple([]lexer.SimpleRule{
		// ******************************************************************
		// *********                  Keywords                      *********
		// ******************************************************************
        {
            Name: `Create`,
            Pattern: `(?i)CREATE`,
        },
        {
            Name: `Table`,
            Pattern: `(?i)TABLE`,
        },
        {
            Name: `Primary`,
            Pattern: `(?i)PRIMARY`,
        },
        {
            Name: `Key`,
            Pattern: `(?i)KEY`,
        },
        {
            Name: `Not`,
            Pattern: `(?i)NOT`,
        },
        {
            Name: `Null`,
            Pattern: `(?i)NULL`,
        },
        {
            Name: `Unique`,
            Pattern: `(?i)UNIQUE`,
        },
		// ******************************************************************
		// *********                 Data Types                     *********
		// ******************************************************************
		{
			Name:    `TextType`,
			Pattern: `(?i)CHAR|VARCHAR|TEXT|TINYTEXT|MEDIUMTEXT|LONGTEXT`,
		},
		{
			Name:    `IntType`,
			Pattern: `(?i)TINYINT|SMALLINT|MEDIUMINT|INT|INTEGER|BIGINT`,
		},
		{
			Name:    `DecimalType`,
			Pattern: `(?i)FLOAT|DOUBLE|DECIMAL|DEC`,
		},
		{
			Name:    `BoolType`,
			Pattern: `(?i)BOOLEAN|BOOL`,
		},
		{
			Name:    `TimeType`,
			Pattern: `(?i)DATE|DATETIME|TIMESTAMP|TIME|YEAR`,
		},

		// ******************************************************************
		// *********               Literal Values                   *********
		// ******************************************************************
		{
			Name:    `Decimal`,
			Pattern: `[-+]?\d*\.\d+([eE][-+]?\d+)?`,
		},
		{
			Name:    `Int`,
			Pattern: `[-+]?\d+([eE][-+]?\d+)?`,
		},
		{
			Name:    `String`,
			Pattern: `'[^']*'|"[^"]*"`,
		},
		{
			Name:    `Bool`,
			Pattern: `(?i)TRUE|FALSE`,
		},

		// ******************************************************************
		// *********                   Others                       *********
		// ******************************************************************
		{
			Name:    `Ident`,
			Pattern: `[a-zA-Z_][a-zA-Z0-9_]*`,
		},
		{
			Name:    `Semicolon`,
			Pattern: `;`,
		},
		{
			Name:    `Operator`,
			Pattern: `<>|!=|<=|>=|[-+*/%,.()=<>]`,
		},
		{
			Name:    "whitespace",
			Pattern: `\s+`,
		},
	})
}
