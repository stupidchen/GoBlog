package view

import (
	"net/http"
	"strconv"
	"blog/model"
	"fmt"
)

type CommentHandler struct {
}

func (h *CommentHandler) Get (r *http.Request) *ResponseData {
	s2 := getSubPath(r.URL.Path, 2)
	if s2 != nil && len(*s2) != 0 {
		id, err := strconv.ParseUint(*s2, 10, 64)
		if err != nil {
			return InitError(err.Error())
		}
		a := model.FindCommentById(uint(id))
		if a == nil {
			return InitError(fmt.Sprintf("Comment %d does not exist.", id))
		}
		return &ResponseData{
			Object:    *a,
			Ok:        true,
			ModelType: "Comment",
			Message:   "Get comment success.",
		}

	} else {
		a := model.GetAllArticles()
		return &ResponseData{
			Object:    *a,
			Ok:        true,
			ModelType: "Comments",
			Message:   "Get comments success.",
		}
	}
	return nil
}

func (h *CommentHandler) Post (r *http.Request, body *RequestData) *ResponseData {
	s2 := getSubPath(r.URL.Path, 2)
	if s2 != nil && len(*s2) != 0 {
		id, err := strconv.ParseInt(*s2, 10, 64)
		if err != nil {
			return InitError(err.Error())
		}
		body.Article.ID = uint(id)
		err = model.AddComment(body.Comment)
		if err != nil {
			return InitError(fmt.Sprintf("Cannot create comment %d due to %s.", id, err.Error()))
		} else {
			return InitHint(fmt.Sprintf("Comment(id: %d) was created.", body.Comment.ID))
		}
	} else {
		err := model.AddComment(body.Comment)
		if err != nil {
			return InitError(fmt.Sprintf("Cannot create comment due to %s.", err.Error()))
		} else {
			return InitHint(fmt.Sprintf("Comment(id: %d) was created.", body.Comment.ID))
		}
	}
}

func (h *CommentHandler) Put (r *http.Request, body *RequestData) *ResponseData {
	s2 := getSubPath(r.URL.Path, 2)
	if s2 != nil && len(*s2) != 0 {
		id, err := strconv.ParseInt(*s2, 10, 64)
		if err != nil {
			return InitError(err.Error())
		}
		body.Comment.ID = uint(id)
		err = model.UpdateComment(body.Comment)
		if err != nil {
			return InitError(fmt.Sprintf("Cannot update comment %d due to %s.", id, err.Error()))
		}
	} else {
		err := model.UpdateComment(body.Comment)
		if err != nil {
			return InitError(fmt.Sprintf("Cannot update comment due to %s.", err.Error()))
		}
	}
	return InitHint(fmt.Sprintf("Comment(id: %d) was updated.", body.Comment.ID))
}

func (h *CommentHandler) Delete (r *http.Request) *ResponseData {
	s2 := getSubPath(r.URL.Path, 2)
	if s2 != nil && len(*s2) != 0 {
		id, err := strconv.ParseInt(*s2, 10, 64)
		if err != nil {
			return InitError(err.Error())
		}
		err = model.DeleteComment(uint(id))
		if err != nil {
			return InitError(fmt.Sprintf("Cannot delete comment(id: %d) due to %s.", id, err.Error()))
		}
		return InitHint(fmt.Sprintf("Comment(id: %d) was deleted.", id))
	} else {
		return InitError("Invalid parameter.")
	}
}

func init() {
	http.HandleFunc("/Comment/", JsonWrapper(&CommentHandler{}))
}
