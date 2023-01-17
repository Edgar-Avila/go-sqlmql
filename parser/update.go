package parser

type Assignment struct {
	Name string  `parser:"@Ident '='"`
	Lit  Literal `parser:"@@"`
}

type UpdateStatement struct {
	Name    string       `parser:"Update @Ident"`
	Changes []Assignment `parser:"Set @@ (',' @@)*"`
	Where   *Where       `parser:"@@? ';'?"`
}

func (u *UpdateStatement) stmt() {}
