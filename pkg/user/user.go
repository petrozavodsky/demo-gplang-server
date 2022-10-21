package user

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends,omitempty"`
}

// MakeUser - создает пользователя
func MakeUser() *User {
	return &User{}
}

// GetId - получает имя
func (u *User) GetId() int {
	return u.Id
}

// GetName - получает имя
func (u *User) GetName() string {
	return u.Name
}

// GetAge - получает возраст
func (u *User) GetAge() int {
	return u.Age
}

// SetId - устанавливает id
func (u *User) SetId(id int) {
	u.Id = id
}

// SetName - устанавливает имя
func (u *User) SetName(name string) {
	u.Name = name
}

// SetAge - устанавливает возраст
func (u *User) SetAge(age int) {
	u.Age = age
}

// GetFriends - получает друзей
func (u *User) GetFriends() []int {
	return u.Friends
}

// SetFriends - задает (всех) друзей
func (u *User) SetFriends(friends []int) []int {
	u.Friends = friends

	return u.Friends
}

// AddFriend - добавляет друга
func (u *User) AddFriend(id int) {

	isFriends := u.isFriends(id)
	if !isFriends {
		u.Friends = append(u.Friends, id)
	}
}

// DeleteFriend - удаляет друга
func (u *User) DeleteFriend(id int) {

	for i := 0; i < len(u.Friends); i++ {
		if u.Friends[i] == id {
			if len(u.Friends) != 0 && i < len(u.Friends)-1 { // protect from panic
				u.Friends = append(u.Friends[:i], u.Friends[i+1:]...)
				return
			}
			u.Friends = u.Friends[:len(u.Friends)-1]
			return
		}
	}
}

// isFriends - проверяет (по id) является пользователь ли другом
func (u *User) isFriends(id int) bool {

	for _, v := range u.Friends {
		if v == id {
			return true
		}
	}
	return false
}
