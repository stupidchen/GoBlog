package view

import (
	"net/http"
	"strconv"
	"blog/model"
	"fmt"
)

type ArticleHandler struct {
}

func (h *ArticleHandler) Get (r *http.Request) *ResponseData {
	s2 := getSubPath(r.URL.Path, 3)
	if s2 != nil && len(*s2) != 0 {
		id, err := strconv.ParseUint(*s2, 10, 64)
		if err != nil {
			return InitError(err.Error())
		}
		a := model.FindArticleById(uint(id))
		if a == nil {
			return InitError(fmt.Sprintf("Article %d does not exist.", id))
		}
		s3 := getSubPath(r.URL.Path, 4)
		if s3 != nil && *s3 == "comment" {
			c := model.FindCommentByArticleId(a.ID)
			return &ResponseData{
				Object: 	*c,
				Ok:			true,
				ModelType: 	"Comments",
				Message:	"Get comments of article success.",
			}
		}
		return &ResponseData{
			Object:    *a,
			Ok:        true,
			ModelType: "Article",
			Message:   "Get article success.",
		}

	} else {
		a := model.GetAllArticles()
		return &ResponseData{
			Object:    *a,
			Ok:        true,
			ModelType: "Articles",
			Message:   "Get articles success.",
		}
	}
}

func (h *ArticleHandler) Post (r *http.Request, body *RequestData) *ResponseData {
	u := loginCheck(r)
	if u == nil {
		return InitError("You have not login.")
	}
	s2 := getSubPath(r.URL.Path, 3)
	if s2 != nil && len(*s2) != 0 {
		id, err := strconv.ParseInt(*s2, 10, 64)
		if err != nil {
			return InitError(err.Error())
		}
		s3 := getSubPath(r.URL.Path, 3)
		if s3 != nil && *s3 == "comment" {
			body.Comment.Article = uint(id)
			body.Comment.Author = *u
			err = model.AddComment(&body.Comment)
			if err != nil {
				return InitError(fmt.Sprintf("Cannot create comment %d due to %s.", id, err.Error()))
			} else {
				return InitHint(fmt.Sprintf("Comment(id: %d) was created.", body.Comment.ID))
			}
		}
	} else {
		body.Article.Author = *u
		err := model.AddArticle(&body.Article)
		if err != nil {
			return InitError(fmt.Sprintf("Cannot create article due to %s.", err.Error()))
		} else {
			return InitHint(fmt.Sprintf("Article(id: %d) was created.", body.Article.ID))
		}
	}
	return InitError("Invalid parameter.")
}

func (h *ArticleHandler) Put (r *http.Request, body *RequestData) *ResponseData {
	u := loginCheck(r)
	if u == nil {
		return InitError("You have not login.")
	}
	s2 := getSubPath(r.URL.Path, 3)
	if s2 != nil && len(*s2) != 0 {
		id, err := strconv.ParseInt(*s2, 10, 64)
		if err != nil {
			return InitError(err.Error())
		}
		a := model.FindArticleById(body.Article.ID)
		if a == nil {
			return InitError(fmt.Sprintf("Article %d does not exist.", id))
		}
		if a.Author != *u {
			return InitError("Unauthorized.")
		}
		err = model.UpdateArticle(&body.Article)
		if err != nil {
			return InitError(fmt.Sprintf("Cannot update article %d due to %s.", id, err.Error()))
		}
	}
	return InitError("Invalid parameter.")
}

func (h *ArticleHandler) Delete (r *http.Request) *ResponseData {
	u := loginCheck(r)
	if u == nil {
		return InitError("You have not login.")
	}
	s2 := getSubPath(r.URL.Path, 3)
	if s2 != nil && len(*s2) != 0 {
		id, err := strconv.ParseInt(*s2, 10, 64)
		if err != nil {
			return InitError(err.Error())
		}
		a := model.FindArticleById(uint(id))
		if a == nil {
			return InitError(fmt.Sprintf("Article %d does not exist.", id))
		}
		if a.Author != *u {
			return InitError("Unauthorized.")
		}
		err = model.DeleteArticle(uint(id))
		if err != nil {
			return InitError(fmt.Sprintf("Cannot delete article(id: %d) due to %s.", id, err.Error()))
		}
		return InitHint(fmt.Sprintf("Article(id: %d) was deleted.", id))
	}
	return InitError("Invalid parameter.")
}

func init() {
	http.HandleFunc("/api/article/", JsonWrapper(&ArticleHandler{}))
}
