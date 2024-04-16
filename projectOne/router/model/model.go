package model

type User struct {
	ID      int64   `json:"id"`
	Name    string  `json:"name"`
	Age     int64   `json:"age"`
	Friends []int64 `json:"friends"`
}

type UsersStorage struct {
	Store map[string]*User
}

func NewUsersStorager() *UsersStorage {
	return &UsersStorage{
		Store: make(map[string]*User),
	}
}

type MakeFriends struct {
	ID1 int64 `json:"source_id"`
	ID2 int64 `json:"target_id"`
}

type RemoveUser struct {
	ID int64 `json:"target_id"`
}

type NewAge struct {
	ID  int64 `json:"target_id"`
	Age int64 `json:"name"`
}
