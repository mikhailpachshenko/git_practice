package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	m "router/model"
)

/* создание нового пользователя */
func (us *m.UsersStorage) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("Some error"))
		}
		defer r.Body.Close()

		var u m.User
		if err := json.Unmarshal(content, &u); err != nil {
			log.Println("Error", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error:" + err.Error()))
			return
		}

		if len(u.Name) < 2 {
			fmt.Println("Короткое имя")
		}

		/* Обрабатываем ошибки */

		if len(u.Name) < 2 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("SomeError"))
			return
		}

		/* покдлючаемся к бд */

	}

}
