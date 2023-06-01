package controller

import (
	"net/http"

	"github.com/gocql/gocql"
)

type User struct {
	ID             gocql.UUID `json:"id"`
	Username       string     `json:"username"`
	RollNumber     string     `json:"roll_number"`
	Email          string     `json:"email"`
	GithubUsername string     `json:"github_username"`
	ProfilePic     string     `json:"profile_pic"`
	Skills         []string   `json:"skills"`
	Interests      []string   `json:"interests"`
	Achievements   []string   `json:"achievements"`
	Role           string     `json:"role"`
}

func Getuser(w http.ResponseWriter, r *http.Request) {
	// Perform necessary operations
	// Write response
	w.Write([]byte("Hello, this is the GetUser handler!"))
}
func Getprojects(w http.ResponseWriter, r *http.Request) {
	// Perform necessary operations
	// Write response
	w.Write([]byte("Hello, this is the Getprojects handler!"))
}

func Creatprofile(w http.ResponseWriter, r *http.Request) {

}
