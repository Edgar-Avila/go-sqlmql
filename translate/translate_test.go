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
}]);`
	fmt.Println(translated)
	if translated != want {
		t.Fatal("Translated and wanted do not match")
	}
}


func TestSelect(t *testing.T) {
    text := 
`
SELECT name FROM students 
WHERE score >= 70 AND graduated = TRUE OR name = "Amadeus"
ORDER BY score DESC
LIMIT 3;
`
    want := "db.students.find({$or:[{$and:[{score:{$gte:70}},{graduated:{$eq:true}}]},{name:{$eq:\"Amadeus\"}}]}, {_id: 0,name: 1}).sort({score: -1}).limit(3);"
    translated := HelpTranslate(t, text)
    fmt.Println(translated)
	if translated != want {
		t.Fatal("Translated and wanted do not match")
	}
}

func TestUpdate(t *testing.T) {
    text := 
`
UPDATE students
SET score = 11
WHERE score = 9 OR score = 10 OR name = "Amadeus"
`
    want := "db.students.updateMany({$or:[{score:{$eq:9}},{score:{$eq:10}},{name:{$eq:\"Amadeus\"}}]}, {$set: {score: 11}})"
    translated := HelpTranslate(t, text)
    fmt.Println(translated)
	if translated != want {
		t.Fatal("Translated and wanted do not match")
	}
}

func TestCreate(t *testing.T) {
    text := 
`
CREATE TABLE students(
    id INT,
    name VARCHAR(100)
)
`
    want := "db.createCollection(\"students\");"
    translated := HelpTranslate(t, text)
    fmt.Println(translated)
	if translated != want {
		t.Fatal("Translated and wanted do not match")
	}
}
