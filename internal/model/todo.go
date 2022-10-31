package model

type TodoList struct {
	Id    int    `json:"id"`
	Title string `json:"title" binding:"required"`
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}
