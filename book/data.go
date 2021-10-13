package book

import "go.mongodb.org/mongo-driver/bson/primitive"

var	NilStudent = Student{}

type Book struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Title   string             `bson:"title,omitempty"`
	Autor   string             `bson:"autor,omitempty"`
	Editor  string             `bson:"editor,omitempty"`
	Year    int                `bson:"release_year,omitempty"`
	LimitDate string		   `bson:"limit_date,omitempty"`
	Student Student            `bson:"student,omitempty"`
}

type Student struct {
	Name  string `bson:"name,omitempty"`
	Class string `bson:"class,omitempty"`
	SRA   RA     `bson:"sra,omitempty"`
}

type RA struct {
	Code  int    `bson:"code,omitempty"`
	Digit string `bson:"digit,omitempty"`
	UF    string `bson:"uf,omitempty"`
}

func RegisterNewBook(title string, autor string, editor string, year int) *Book {
	return &Book{
		Title: title,
		Autor: autor,
		Editor: editor,
		Year: year,
		Student: NilStudent,
	}
}

func (b *Book) AlterateStudentOfBook(stu Student, ld string) {
	b.Student = stu
	b.LimitDate = ld
}

func (b *Book) AlterateLimitDateOfBook(ld string) {
	b.LimitDate = ld
}

func (b *Book) RemoveBook() {
	
}
