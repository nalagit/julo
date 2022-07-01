package db

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

func EnableWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authToken := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]

	_, err := validateToken(authToken)

	if err == nil {
		db, err := OpenConnection()
		defer db.Close()

		var wlt WalletStatus
		json.NewDecoder(r.Body).Decode(&wlt)

		qString := "select status, userid, balance " +
			"from wallet " +
			"where walletid=?"
		var status string
		var user string
		var balance int
		result := ExecuteSQL(qString, &wlt.Id)
		for result.Next() {
			err = result.Scan(&status, &user, &balance)
			if err != nil {
				gr := GenericResponse{}
				gr.Status = "Failed"
				gr.Message = err.Error()
				json.NewEncoder(w).Encode(gr)
				panic(err.Error())
			}
		}

		if status != "enabled" {
			const timeLayout = "2006-01-02 15:04:05"
			currentTime, _ := time.Parse(timeLayout, time.Now().Format(timeLayout))

			qString = "update wallet set status=?, enabled_at=? where walletid=?"
			fmt.Println("Update Wallet Status")
			_, err = db.Exec(qString, "enabled", currentTime, &wlt.Id)
			if err != nil {
				gr := GenericResponse{}
				gr.Status = "Failed"
				gr.Message = err.Error()
				json.NewEncoder(w).Encode(gr)
				panic(err.Error())
			}
			wr := Wallet{
				Id:        wlt.Id,
				OwnedBy:   user,
				Status:    "enabled",
				EnabledAt: currentTime.String(),
				Balance:   balance,
			}
			json.NewEncoder(w).Encode(wr)
		} else {
			gr := GenericResponse{}
			gr.Status = "Failed"
			gr.Message = "Wallet already enabled"
			json.NewEncoder(w).Encode(gr)
		}

	} else {
		gr := GenericResponse{
			Status:  "Failed",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(gr)
	}

}

func DisableWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authToken := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]

	_, err := validateToken(authToken)

	if err == nil {
		db, err := OpenConnection()
		defer db.Close()

		var wlt WalletStatus
		json.NewDecoder(r.Body).Decode(&wlt)

		qString := "select status, userid, balance " +
			"from wallet " +
			"where walletid=?"
		var status string
		var user string
		var balance int
		result := ExecuteSQL(qString, &wlt.Id)
		for result.Next() {
			err = result.Scan(&status, &user, &balance)
			if err != nil {
				gr := GenericResponse{}
				gr.Status = "Failed"
				gr.Message = err.Error()
				json.NewEncoder(w).Encode(gr)
				panic(err.Error())
			}
		}

		if status != "disabled" {
			const timeLayout = "2006-01-02 15:04:05"
			currentTime, _ := time.Parse(timeLayout, time.Now().Format(timeLayout))

			qString = "update wallet set status=?, disabled_at=? where walletid=?"
			fmt.Println("Update Wallet Status")
			_, err = db.Exec(qString, "disabled", currentTime, &wlt.Id)
			if err != nil {
				gr := GenericResponse{}
				gr.Status = "Failed"
				gr.Message = err.Error()
				json.NewEncoder(w).Encode(gr)
				panic(err.Error())
			}
			wr := Wallet{
				Id:        wlt.Id,
				OwnedBy:   user,
				Status:    "disabled",
				EnabledAt: currentTime.String(),
				Balance:   balance,
			}
			json.NewEncoder(w).Encode(wr)
		} else {
			gr := GenericResponse{}
			gr.Status = "Failed"
			gr.Message = "Wallet already enabled"
			json.NewEncoder(w).Encode(gr)
		}

	} else {
		gr := GenericResponse{
			Status:  "Failed",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(gr)
	}

}

func GetWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authToken := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]

	_, err := validateToken(authToken)

	if err == nil {
		db, err := OpenConnection()
		defer db.Close()

		var wlt WalletStatus
		json.NewDecoder(r.Body).Decode(&wlt)

		qString := "select userid, status, enabled_at, balance " +
			"from wallet " +
			"where walletid=?"
		var status string
		var user string
		var balance int
		var enabledat time.Time
		result := ExecuteSQL(qString, &wlt.Id)
		for result.Next() {
			err = result.Scan(&user, &status, &enabledat, &balance)
			if err != nil {
				gr := GenericResponse{}
				gr.Status = "Failed"
				gr.Message = err.Error()
				json.NewEncoder(w).Encode(gr)
				panic(err.Error())
			}
		}
		if status != "disabled" {
			wr := Wallet{
				Id:        wlt.Id,
				OwnedBy:   user,
				Status:    status,
				EnabledAt: enabledat.String(),
				Balance:   balance,
			}
			json.NewEncoder(w).Encode(wr)
		} else {
			gr := GenericResponse{
				Status:  "Failed",
				Message: "Account Disabled",
			}
			json.NewEncoder(w).Encode(gr)
		}

	} else {
		gr := GenericResponse{
			Status:  "Failed",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(gr)
	}
}

func DepositWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authToken := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]

	_, err := validateToken(authToken)

	if err == nil {
		db, err := OpenConnection()
		defer db.Close()

		var wlt DepositRequest
		json.NewDecoder(r.Body).Decode(&wlt)

		qString := "select status, balance " +
			"from wallet " +
			"where walletid=?"
		var status string
		var balance int
		result := ExecuteSQL(qString, &wlt.Id)
		for result.Next() {
			err = result.Scan(&status, &balance)
			if err != nil {
				gr := GenericResponse{}
				gr.Status = "Failed"
				gr.Message = err.Error()
				json.NewEncoder(w).Encode(gr)
				panic(err.Error())
			}
		}

		if status != "disabled" {
			const timeLayout = "2006-01-02 15:04:05"
			currentTime, _ := time.Parse(timeLayout, time.Now().Format(timeLayout))

			updatedBal := balance + wlt.Amount
			qString = "update wallet set balance=? where walletid=?"
			fmt.Println("Update Wallet Balance")
			_, err = db.Exec(qString, updatedBal, &wlt.Id)
			if err != nil {
				gr := GenericResponse{}
				gr.Status = "Failed"
				gr.Message = err.Error()
				json.NewEncoder(w).Encode(gr)
				panic(err.Error())
			}
			wr := Deposit{
				Id:          wlt.Id,
				DepositBy:   uuid.NewString(),
				Status:      "enabled",
				DepositAt:   currentTime.String(),
				Amount:      wlt.Amount,
				ReferenceId: uuid.NewString(),
			}
			json.NewEncoder(w).Encode(wr)
		} else {
			gr := GenericResponse{}
			gr.Status = "Failed"
			gr.Message = "Wallet already enabled"
			json.NewEncoder(w).Encode(gr)
		}

	} else {
		gr := GenericResponse{
			Status:  "Failed",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(gr)
	}

}

func WithdrawWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authToken := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]

	_, err := validateToken(authToken)

	if err == nil {
		db, err := OpenConnection()
		defer db.Close()

		var wlt DepositRequest
		json.NewDecoder(r.Body).Decode(&wlt)

		qString := "select status, balance " +
			"from wallet " +
			"where walletid=?"
		var status string
		var balance int
		result := ExecuteSQL(qString, &wlt.Id)
		for result.Next() {
			err = result.Scan(&status, &balance)
			if err != nil {
				gr := GenericResponse{}
				gr.Status = "Failed"
				gr.Message = err.Error()
				json.NewEncoder(w).Encode(gr)
				panic(err.Error())
			}
		}

		if status != "disabled" {
			const timeLayout = "2006-01-02 15:04:05"
			currentTime, _ := time.Parse(timeLayout, time.Now().Format(timeLayout))

			updatedBal := balance - wlt.Amount
			if updatedBal < 0 {
				gr := GenericResponse{}
				gr.Status = "Failed"
				gr.Message = "Wallet balance insufficient"
				json.NewEncoder(w).Encode(gr)
			} else {
				qString = "update wallet set balance=? where walletid=?"
				fmt.Println("Update Wallet Balance")
				_, err = db.Exec(qString, updatedBal, &wlt.Id)
				if err != nil {
					gr := GenericResponse{}
					gr.Status = "Failed"
					gr.Message = err.Error()
					json.NewEncoder(w).Encode(gr)
					panic(err.Error())
				}
				wr := Deposit{
					Id:          wlt.Id,
					DepositBy:   uuid.NewString(),
					Status:      "enabled",
					DepositAt:   currentTime.String(),
					Amount:      wlt.Amount,
					ReferenceId: uuid.NewString(),
				}
				json.NewEncoder(w).Encode(wr)
			}

		} else {
			gr := GenericResponse{}
			gr.Status = "Failed"
			gr.Message = "Wallet already enabled"
			json.NewEncoder(w).Encode(gr)
		}

	} else {
		gr := GenericResponse{
			Status:  "Failed",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(gr)
	}

}
