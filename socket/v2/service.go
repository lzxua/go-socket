package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
)

type UserMapStruct struct {
	userMaps sync.Map
	lock     sync.Mutex
}

var usermap *UserMapStruct

func ServiceStart() {
	usermap = &UserMapStruct{
		userMaps: sync.Map{},
		lock:     sync.Mutex{},
	}
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("run error")
	}
	count := 0
	for {
		c, _ := l.Accept()
		c.Write([]byte("上线成功\n"))
		user := NewUser(c, "user"+strconv.Itoa(count))
		count++
		go getDate(user)
	}
}

// 获取传过来的数据
func getDate(user *User) {
	for {
		buf := make([]byte, 4096)
		n, _ := user.conn.Read(buf)
		if user.conn == nil {
			break
		}
		if _, ok := usermap.userMaps.Load(user.conn.RemoteAddr().String()); !ok {
			user.Online()
		}
		if n == 0 {
			continue
		}
		ChoiceTree(user, buf, n)
	}
}

// ChoiceTree 选择树
func ChoiceTree(user *User, buff []byte, n int) {
	inStr := strings.TrimSpace(string(buff[0:n]))
	split := strings.Split(inStr, " ")
	switch split[0] {
	case "quit":
		fmt.Printf("下线成功")
		user.Offline()
	case "from":
		PrivateChat(user, split[1], split[2])
	default:
		GroupChat(user, inStr)
	}
	return
}

// GroupChat 群聊
func GroupChat(user *User, msg string) {
	per := AddSuffix(user.UserName)
	usermap.userMaps.Range(func(key, value any) bool {
		value.(*User).conn.Write([]byte(per + msg + "\n"))
		return true
	})
}

// PrivateChat 私聊
func PrivateChat(user *User, from, msg string) {
	if _, ok := usermap.userMaps.Load(from); !ok {
		user.conn.Write([]byte("无此用户或用户未上限\n"))
		return
	}
	per := AddSuffix(user.UserName + "[private]")
	value, ok := usermap.userMaps.Load(from)
	if ok {
		value.(*User).conn.Write([]byte(per + msg))
	}
	return
}

func AddSuffix(str string) string {
	return str + ":"
}
