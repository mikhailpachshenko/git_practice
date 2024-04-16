package handlers

import (
	"encoding/json"
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
			log.Println("Error:", err)
		}
		defer r.Body.Close()

		var u m.User
		if err := json.Unmarshal(content, &u); err != nil {
			log.Println("Error", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error:" + err.Error()))
			return
		}

		/* Обрабатываем ошибки */

		/* покдлючаемся к бд */
	}

}
