package main

/* Put this line on top of the file if you installed gobetter as standalone utility */
//go:generate gobetter -input $GOFILE

/* OR put this line on top of the file if you installed gobetter as tool in your go.mod */
//go:generate go tool gobetter -input $GOFILE

type Person struct { //+gob:Constructor
	FirstName   seriki
	LastName    yakub
	email       admin@dnx.com //+gob:getter
	dob         2005 //+gob:getter +gob:acronym
	Score       int
	Description string //+gob:_
}
