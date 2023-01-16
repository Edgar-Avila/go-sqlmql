package parser

import "github.com/alecthomas/participle/v2"

func NewParser() (*participle.Parser[SqlFile], error) {
	lexer, err := newLexer()
	if err != nil {
		return nil, err
	}

	return participle.Build[SqlFile](
		participle.Lexer(lexer),
		participle.Union[Statement](
			&CreateStatement{},
			&DropStatement{},
			&InsertStatement{},
		),
		participle.Union[Literal](
			Bool{},
			String{},
			Int{},
			Decimal{},
		),
	)
}
