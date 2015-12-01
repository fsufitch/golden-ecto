package db

import (
	"crypto/sha256"
	"crypto/rand"

	"github.com/yasvisu/gw2api"
)

type User struct {
	Username string
	ApiKey string
	Salt []byte
	PassHash []byte
}


func createSalt(length int) []byte {
	salt := make([]byte, length) 
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}
	return salt
}

func createPasswordHash(password string, salt []byte) []byte {
	fullToken := []byte{}
	fullToken = append(fullToken, salt...)
	fullToken = append(fullToken, []byte(password)...)
	hash := sha256.Sum256(fullToken)
	return hash[:]
}

func CreateUserTable() {
	transaction, err := DbConn.Begin()
	_, err = transaction.Exec(`
		CREATE TABLE IF NOT EXISTS users(
		username TEXT PRIMARY KEY NOT NULL,
		apikey TEXT NOT NULL,
		salt BLOB NOT NULL,
		passhash BLOB NOT NULL
		)`)

	if err != nil { panic(err) }
	transaction.Commit()
}

func GetUserByName(username string) (u *User) {
	var count int
	err := DbConn.QueryRow(`SELECT count(*) FROM users WHERE username=?`, username).Scan(&count)
	if err != nil { panic(err) }

	if count == 0 {
		u = nil
		return
	}
	
	u = &User{}
	err = DbConn.QueryRow(`SELECT username, apikey, salt, passhash FROM users WHERE username=?`, username).Scan(
		&u.Username, &u.ApiKey, &u.Salt, &u.PassHash)
	
	if err != nil { panic(err) }
	return
}

func UsernameFromApiLazy(apikey string) (username string, userExists bool) {
	var count int
	err := DbConn.QueryRow(`SELECT count(*) FROM users WHERE apikey=?`, apikey).Scan(&count)
	if err != nil { panic(err) }
	
	if count > 0 {
		err := DbConn.QueryRow(`SELECT username FROM users WHERE apikey=?`, apikey).Scan(&username)
		if err != nil { panic(err) }
		userExists = true
		return
	}

	userExists = false

	api, err := gw2api.NewAuthenticatedGW2Api(apikey)
	if err != nil { panic(err) }

	account, err := api.Account()
	if err != nil { panic(err) }

	username = account.Name
	return
}

/// User Object

func NewUser(username, apikey, password string) (u *User) {
	salt := createSalt(64)
	passHash := createPasswordHash(password, salt)

	u = &User{
		Username: username,
		ApiKey: apikey,
		Salt: salt,
		PassHash: passHash,
	}
	return
}

func (u *User) Store() {
	var count int
	err := DbConn.QueryRow(`SELECT count(*) FROM users WHERE username=?`, u.Username).Scan(&count)
	if err != nil { panic(err) }

	tx, _ := DbConn.Begin()
	if count == 0 { // New user, do insert
		_, err = tx.Exec(`INSERT INTO users (username, apikey, salt, passhash) VALUES (?,?,?,?)`,
			u.Username, u.ApiKey, u.Salt, u.PassHash)
	} else {
		_, err = tx.Exec(`UPDATE users SET apikey=?, salt=?, passhash=? WHERE username=?`,
			u.ApiKey, u.Salt, u.PassHash, u.Username)
	}
	tx.Commit()

	if err != nil { panic(err) }
}

