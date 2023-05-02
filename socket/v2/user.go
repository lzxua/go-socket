package main

import (
	"net"
)

type User struct {
	conn     net.Conn
	UserName string
}

func NewUser(c net.Conn, name string) *User {
	return &User{
		conn:     c,
		UserName: name,
	}
}

func (user *User) SendMsg(msg string) {
	user.conn.Write([]byte(msg))
}

func (user *User) Online() {
	usermap.lock.Lock()
	defer usermap.lock.Unlock()
	usermap.userMaps[user.UserName] = user
}

func (user *User) Offline() {
	user.conn.Write([]byte("欢迎下次登陆\n"))
	usermap.lock.Lock()
	defer usermap.lock.Unlock()
	delete(usermap.userMaps, user.conn.RemoteAddr().String())
	user.conn.Close()
	return
}
