package parser

type InsertedRow struct {
	Values []Literal `parser:"'(' @@ (',' @@)* ')'"`
}

type InsertStatement struct {
	Into string        `parser:"Insert Into @Ident"`
	Cols []string      `parser:"('(' @Ident (',' @Ident)* ')')?"`
	Rows []InsertedRow `parser:"Values @@ (',' @@)* ';'?"`
}

func (i *InsertStatement) stmt() {}
