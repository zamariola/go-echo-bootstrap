package internal

import "log"

//ServerStatus defines a webserver state
type ServerStatus struct {
	Status string `json:"status,omitempty"`
}

//User defines a instance of an user entity
type User struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

//NewUser creates a new User instance
func NewUser(name, email string) *User {
	return &User{name, email}
}

//Save persists it into database
func (u *User) Save() error {

	db, err := DB()
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO users(name,email) VALUES($1,$2)")
	if err != nil {
		log.Fatal(err)
	}

	if _, err = stmt.Exec(u.Name, u.Email); err != nil {
		return err
	}

	return nil
}

//FindUser searchs within the database for a specific user
func FindUser(name string) (*User, error) {

	db, err := DB()
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare("SELECT name, email FROM users WHERE name = $1")
	if err != nil {
		log.Fatal(err)
	}

	u := &User{}

	sRow := stmt.QueryRow(name)
	if err = sRow.Scan(&u.Name, &u.Email); err != nil {
		return nil, err
	}
	return u, err
}
