package parser

// Drop statement containing one or more identifiers
type DropStatement struct {
	Tables []string `parser:"Drop Temporary? Table (If Exists)? @Ident (',' @Ident)* ';'?"`
}

func (d *DropStatement) stmt() {}
