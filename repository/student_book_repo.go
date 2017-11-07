package repository

import (
	"database/sql"

	"github.com/supunz/go-crud/db"
)

//StudentBookRepo - StudentBook repository
type StudentBookRepo struct {
	Database *sql.DB
	Name     string
}

//GetStudentBookRepository -sdsf
func GetStudentBookRepository() (StudentBookRepo, error) {
	db, err := db.GetDatabase()
	return StudentBookRepo{Database: db, Name: "student_book"}, err
}

//Select - Select books from db
func (repo StudentBookRepo) Select() ([]interface{}, error) {
	//todo - implement here
	return nil, nil
}

//Insert - Insert books to db
func (repo StudentBookRepo) Insert(doc interface{}) error {
	//todo - implement here
	return nil
}

//Update - Update books
func (repo StudentBookRepo) Update(doc interface{}) error {
	//todo - implement here
	return nil
}

//Remove - Delete books from db
func (repo StudentBookRepo) Remove(doc interface{}) error {
	//todo - implement here
	return nil
}

//Close - close database
func (repo StudentBookRepo) Close() {

}