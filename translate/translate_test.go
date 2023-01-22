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

func TestInsert(t *testing.T) {
    text := 
    `
		INSERT INTO students
        (id, name, lastname, graduated, score)
        VALUES
        (1, 'John', 'Doe', TRUE, 10.0),
        (2, 'Jane', 'Doe', FALSE, 10.0);
    `
	translated := HelpTranslate(t, text)
	want := 
`db.students.insertMany([{
  id: 1,
  name: "John",
  lastname: "Doe",
  graduated: true,
  score: 10
},
{
  id: 2,
  name: "Jane",
  lastname: "Doe",
  graduated: false,
  score: 10
}]`
	fmt.Println(translated)
	if translated != want {
		t.Fatal("Translated and wanted do not match")
	}
}


