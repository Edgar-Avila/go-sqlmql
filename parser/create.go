package parser

// The type of a column, it has a name and parameters
type ColumnType struct {
	Name   string `parser:"@(TextType|IntType|BoolType|DecimalType|TimeType)"`
	Params []int  `parser:"('(' @Int (',' @Int)* ')')?"`
}

// Column declaration, with a name, type and modifiers
type ColumnDeclaration struct {
	Name      string     `parser:"@Ident"`
	Type      ColumnType `parser:"@@"`
	Modifiers string     `parser:"@(Primary Key|Not Null|Unique)*"`
}

// Create statement, wiht a name and the columns needed
type CreateStatement struct {
	Name string              `parser:"Create Table @Ident"`
	Cols []ColumnDeclaration `parser:"'(' @@ (',' @@)* ')'"`
}

func (c *CreateStatement) stmt() {}
