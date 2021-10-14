package book

import "go.mongodb.org/mongo-driver/bson/primitive"

var	NilStudent = Student{}

type Book struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Title   string             `bson:"title,omitempty"`
	Autor   string             `bson:"autor,omitempty"`
	Editor  string             `bson:"editor,omitempty"`
	Year    string             `bson:"release_year,omitempty"`
	LimitDate string		   `bson:"limit_date,omitempty"`
	Student Student            `bson:"student,omitempty"`
}

type Student struct {
	Name  string `bson:"name,omitempty"`
	Class string `bson:"class,omitempty"`
	SRA   RA     `bson:"sra,omitempty"`
}

type RA struct {
	Code  string    `bson:"code,omitempty"`
	Digit string `bson:"digit,omitempty"`
	UF    string `bson:"uf,omitempty"`
}

func BookFromForm(form map[string]string) *Book {
	return &Book{
		Title: form["title"],
		Autor: form["autor"],
		Editor: form["editor"],
		Year: form["pubYear"],
		LimitDate: form["dtLimit"],
		Student: Student{
			Name: form["sName"],
			Class: form["sClass"],
			SRA: RA{
				Code: form["raCod"],
				Digit: form["raDig"],
				UF: form["raUf"],
			},
		},
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
