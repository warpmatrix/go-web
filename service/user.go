package service

import (
	"net/url"
)

type user struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}

var uList []user

func isValid(form url.Values) bool {
	fname := form["fname"][0]
	lname := form["lname"][0]
	return len(fname) > 0 && len(lname) > 0
}

func parseUser(form url.Values) user {
	user := user{
		Fname: form["fname"][0],
		Lname: form["lname"][0],
	}
	return user
}
