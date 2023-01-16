package parser

type DeleteStatement struct {
	Name  string `parser:"Delete From @Ident"`
	Where *Where `parser:"@@?"`
}

func (d *DeleteStatement) stmt() {}
