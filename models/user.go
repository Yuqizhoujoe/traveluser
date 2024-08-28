package models

type User struct {
	Rooms []string `firestore:"rooms"`
	Posts []string `firestore:"posts"`
}
