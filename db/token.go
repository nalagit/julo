package db

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"golang.org/x/crypto/bcrypt"
)

func registerUser(username string, password string) GenericResponse {
	db, err := OpenConnection()
	defer db.Close()

	qString := "insert into user(username, password) values (?, ?)"

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	_, err = db.Exec(qString, username, hashedPassword)

	if err != nil {
		panic(err.Error())
	}

	var uid string
	qString = "select userid from user where username=?"
	result := ExecuteSQL(qString, username)
	for result.Next() {
		err = result.Scan(&uid)
		if err != nil {
			panic(err.Error())
		}
	}

	qString = "insert into wallet(userid,balance,status) values(?,?,?)"
	_, err = db.Exec(qString, uid, 0, "disabled")
	if err != nil {
		panic(err.Error())
	}

	gr := GenericResponse{
		Status:  "Success",
		Message: "User Registration Success",
	}
	return gr

}

func generateToken(username string, password string) (UsersToken, error) {
	db, err := OpenConnection()
	defer db.Close()
	qString := "select username, password from user where username = ?"

	result := ExecuteSQL(qString, username)
	var usr Users
	var getpassword string
	for result.Next() {
		err = result.Scan(&usr.Username, &usr.Password)
		if err != nil {
			panic(err.Error)
		}
		getpassword = usr.Password
	}

	err = bcrypt.CompareHashAndPassword([]byte(getpassword), []byte(password))

	if err != nil {
		panic(err.Error())
	}

	qString = "update user set token=?, generated_at=?, expired_at=? where username=?"
	randomToken := make([]byte, 32)

	_, err = rand.Read(randomToken)

	if err != nil {
		panic(err.Error())
	}

	authToken := base64.URLEncoding.EncodeToString(randomToken)

	const timeLayout = "2006-01-02 15:04:05"

	dt := time.Now()
	expirtyTime := time.Now().Add(time.Minute * 5)

	generatedAt := dt.Format(timeLayout)
	expiresAt := expirtyTime.Format(timeLayout)

	_, err = db.Exec(qString, authToken, generatedAt, expiresAt, username)

	if err != nil {
		panic(err.Error())
	}

	tokenDetails := UsersToken{
		Username:    username,
		Token:       authToken,
		GeneratedAt: generatedAt,
		ExpiredAt:   expiresAt,
	}

	return tokenDetails, nil
}

func validateToken(authToken string) (UsersToken, error) {
	db, err := OpenConnection()
	defer db.Close()

	qString := `select 
                username,
                generated_at,
                expired_at                         
            from user
            where token = ?`

	result := ExecuteSQL(qString, authToken)
	var ut UsersToken
	for result.Next() {
		err = result.Scan(&ut.Username, &ut.GeneratedAt, &ut.ExpiredAt)
		if err != nil {
			panic(err.Error())
		}
	}

	const timeLayout = "2006-01-02 15:04:05"
	const timeLayout2 = "2006-01-02T15:04:05Z"
	expiryTime, _ := time.Parse(timeLayout2, ut.ExpiredAt)
	currentTime, _ := time.Parse(timeLayout, time.Now().Format(timeLayout))

	if expiryTime.Before(currentTime) {
		return UsersToken{}, errors.New("The token is expired.")
	}

	userDetails := UsersToken{
		Username:    ut.Username,
		Token:       ut.Token,
		GeneratedAt: ut.GeneratedAt,
		ExpiredAt:   ut.ExpiredAt,
	}

	return userDetails, nil

}
