package accountHandler

import (
	"encoding/json"
	"net/http"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"without.framework/domain/account"
	accountRepository "without.framework/infrastructure/account"
)

type requestBody struct {
	UserId string `json:"userId"`
}

func AddAccount(res http.ResponseWriter, req *http.Request) {
	if req.Body == nil {
		http.Error(res, "Please send a request body", 400)
		return
	}

	var body requestBody
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	// Create account
	account := account.New(body.UserId)

	// Save account
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/golang")

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	repo := accountRepository.New(db)

	err2 := repo.Save(account)
	if err2 != nil {
		http.Error(res, err2.Error(), http.StatusInternalServerError)
		return
	}
}
