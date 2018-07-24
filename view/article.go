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
	s2 := getSubPath(r.URL.Path, 2)
	if s2 != nil && len(*s2) != 0 {
		id, err := strconv.ParseUint(*s2, 10, 64)
		if err != nil {
			return InitError(err.Error())
		}
		a := model.FindArticleById(uint(id))
		if a == nil {
			return InitError(fmt.Sprintf("Article %d does not exist.", id))
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
	s2 := getSubPath(r.URL.Path, 2)
	if s2 != nil && len(*s2) != 0 {
		id, err := strconv.ParseInt(*s2, 10, 64)
		if err != nil {
			return InitError(err.Error())
		}
		body.Article.ID = uint(id)
		err = model.AddArticle(body.Article)
		if err != nil {
			return InitError(fmt.Sprintf("Cannot update article %d due to %s.", id, err.Error()))
		} else {
			return InitHint(fmt.Sprintf("Article(id: %d) is updated.", body.Article.ID))
		}
	} else {
		err := model.AddArticle(body.Article)
		if err != nil {
			return InitError(fmt.Sprintf("Cannot create article due to %s.", err.Error()))
		} else {
			return InitHint(fmt.Sprintf("Article(id: %d) is created.", body.Article.ID))
		}
	}
}

func (h *ArticleHandler) Put (r *http.Request, body *RequestData) *ResponseData {
	s2 := getSubPath(r.URL.Path, 2)
	if s2 != nil && len(*s2) != 0 {
		id, err := strconv.ParseInt(*s2, 10, 64)
		if err != nil {
			return InitError(err.Error())
		}
		body.Article.ID = uint(id)
		err = model.UpdateArticle(body.Article)
		if err != nil {
			return InitError(fmt.Sprintf("Cannot update article %d due to %s.", id, err.Error()))
		}
	} else {
		err := model.UpdateArticle(body.Article)
		if err != nil {
			return InitError(fmt.Sprintf("Cannot update article due to %s.", err.Error()))
		}
	}
	return nil
}

func (h *ArticleHandler) Delete (r *http.Request) *ResponseData {
	s2 := getSubPath(r.URL.Path, 2)
	if s2 != nil && len(*s2) != 0 {
		id, err := strconv.ParseInt(*s2, 10, 64)
		if err != nil {
			return InitError(err.Error())
		}
		err = model.DeleteArticle(uint(id))
		if err != nil {
			return InitError(fmt.Sprintf("Cannot delete article(id: %d) due to %s.", id, err.Error()))
		}
		return InitHint(fmt.Sprintf("Article(id: %d) was deleted.", id))
	} else {
		return InitError("Invalid parameter.")
	}
}

func init() {
	http.HandleFunc("/Article/", JsonWrapper(&ArticleHandler{}))
}
