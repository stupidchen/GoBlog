package view

import (
	"net/http"
	"fmt"
	"blog/model"
)

type UserHandler struct {
}

func (h *UserHandler) Get(r *http.Request) *ResponseData {
	username := getSubPath(r.URL.Path, 3)
	if username == nil {
		return InitError("User name is missing!")
	}
	return InitHint(fmt.Sprintf("Getting User %s", *username))
}

func (h *UserHandler) Post(r *http.Request, body *RequestData) *ResponseData {
	s2 := getSubPath(r.URL.Path, 3)
	if s2 != nil {
		switch *s2 {
		case "login":
			if loginCheck(r) != nil {
				return InitError("You have already login.")
			}
			u := model.User{
				Username: body.User.Username,
				Password: generateBase64OfSha1(body.User.Password),
			}
			tu := model.FindUser(&u)
			if tu == nil {
				return InitError("Invalid username or password.")
			}
			token := *addToken(tu.ID)
			return InitHint(fmt.Sprintf("User %s login success. Token is %s", body.User.Username, token))
			break
		case "logout":
			if loginCheck(r) == nil {
				return InitError("You have not login.")
			} else {
				clearToken(r)
				return InitHint("Logout success.")
			}
			break
		default:
			body.User.Password = generateBase64OfSha1(body.User.Password)
			err := model.AddUser(&body.User)
			if err != nil {
				return InitError(fmt.Sprintf("Create user failed due to %s", err.Error()))
			} else {
				return InitHint(fmt.Sprintf("Create user %s success.", body.User.Username))
			}
		}
	}
	return InitError("Invalid parameter.")
}

func (h *UserHandler) Put(r *http.Request, body *RequestData) *ResponseData {
	return InitHint(fmt.Sprintf("Unsupported method: PUT."))
}

func (h *UserHandler) Delete(r *http.Request) *ResponseData {
	return InitHint(fmt.Sprintf("Unsupported method: DELETE."))
}

func init() {
	http.HandleFunc("/api/user/", JsonWrapper(&UserHandler{}))
}