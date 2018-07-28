package view

import (
	"net/http"
	"strings"
	"blog/global"
	"github.com/google/uuid"
	"math/rand"
	"strconv"
)

func getRequestBody(r *http.Request) *string {
	l := r.ContentLength
	buf := make([]byte, l)
	rl, _ := r.Body.Read(buf)
	if rl == 0 {
		return nil
	}
	s := string(buf)
	return &s
}

func getSubPath(path string, index int) *string {
	t := strings.Split(path, "/")
	if len(t) <= index {
		return nil
	} else {
		return &t[index]
	}
}

func loginCheck(r *http.Request) *uint {
	var ret *uint
	token := r.Header.Get("Goblog-Token")
	tokens := *global.GLOBAL.Token
	if id, valid := tokens[token]; valid {
		ret = &id
	}

	return ret
}

func clearToken(r *http.Request) {
	token := r.Header.Get("Goblog-Token")
	tokens := *global.GLOBAL.Token
	if _, valid := tokens[token]; valid {
		delete(tokens, token)
	}
}

func addToken(u uint) *string {
	var token string
	tokenBytes, err := uuid.NewRandom()
	if err != nil {
		global.GLOBAL.Logger.Printf("Generate uuid error: %s. Use simple token.", err.Error())
		token = strconv.FormatInt(rand.Int63(), 16)
	} else {
		token = string(tokenBytes[:])
	}
	tokens := *global.GLOBAL.Token
	tokens[token] = u
	return &token
}
