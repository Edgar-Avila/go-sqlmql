package translate

import (
	"fmt"
	"testing"

	"github.com/Edgar-Avila/go-sqlmql/parser"
)

func HelpTranslate(t *testing.T, text string) string {
	t.Helper()
	parser, err := parser.NewParser()
	if err != nil {
		t.Fatal(err)
	}
	parsed, err := parser.ParseString("", text)
	if err != nil {
		t.Fatal(err)
	}
	translated, err := TranslateSqlFile(parsed)
	if err != nil {
		t.Fatal(err)
	}
	return translated
}

func TestDrop(t *testing.T) {
	translated := HelpTranslate(t, `DROP TABLE students, users;`)
	want := "db.students.drop();\ndb.users.drop();\n"
	fmt.Println(translated)
	if translated != want {
		t.Fatal("Translated and wanted do not match")
	}
}
