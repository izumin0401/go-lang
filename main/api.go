package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
	Password string `json:"-"` // パスワードは出力しない
	AdminFlg int
}

type Users []User

func fetchUserJson() Users {
	return Users{
		User{UserId: "001", UserName: "ずみちゃん", Password: "qwertyuiop", AdminFlg: 1},
		User{UserId: "002", UserName: "ずみちゃん", Password: "qwertyuiop", AdminFlg: 0},
		User{UserId: "003", UserName: "ずみちゃん", Password: "qwertyuiop", AdminFlg: 0},
		User{UserId: "004", UserName: "ずみちゃん", Password: "qwertyuiop", AdminFlg: 0},
		User{UserId: "005", UserName: "ずみちゃん", Password: "qwertyuiop", AdminFlg: 0},
		User{UserId: "006", UserName: "ずみちゃん", Password: "qwertyuiop", AdminFlg: 0},
		User{UserId: "007", UserName: "ずみちゃん", Password: "qwertyuiop", AdminFlg: 0},
		User{UserId: "008", UserName: "ずみちゃん", Password: "qwertyuiop", AdminFlg: 0},
		User{UserId: "009", UserName: "ずみちゃん", Password: "qwertyuiop", AdminFlg: 0},
		User{UserId: "010", UserName: "ずみちゃん", Password: "qwertyuiop", AdminFlg: 0},
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ファックユー")
}

func userIndexHandler(w http.ResponseWriter, r *http.Request) {
	user := fetchUserJson()
	json, err := json.Marshal(user)

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Fprint(w, string(json))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/user", userIndexHandler)

	if err := http.ListenAndServe(":9999", nil); err != nil {
		fmt.Println(err)
	}
}
