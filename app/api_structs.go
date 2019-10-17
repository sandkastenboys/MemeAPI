package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
/*
List of Expected Rest Input Data
 */

type getByID struct {
	ID uint32 		`json:id`
}

type getbyUUID struct {
	ID string
}

type getMemeSearch struct {
	tags []string
}

/*
List of standard return values
 */

type invalidRequest struct {
	error string
}

type sizeResponse struct {
	size uint32
}

type searchResponse struct {
	ids []uint32
}

/*
List of important internaly used data structures
 */

type databaseInforamtion struct {
	db * sql.DB
}


var dbInfo databaseInforamtion

type memeInfo struct{
	ID uint32			`json:id`
	link string			`json:link`
	postTime string
	creator uint32 //TODO: Maybe Change to UUID
}


