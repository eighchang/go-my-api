package models

import "time"

var (
	Comment1 = Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "test comment1",
		CreatedAt: time.Now(),
	}

	Comment2 = Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "test comment2",
		CreatedAt: time.Now(),
	}
)

var (
	Aritce1 = Article{
		ID:          1,
		Title:       "first article",
		Contents:    "This is the test article.",
		UserName:    "eiji",
		NiceNum:     1,
		CommentList: []Comment{Comment1, Comment2},
		CreatedAt:   time.Now(),
	}

	Aritce2 = Article{
		ID:        1,
		Title:     "seconds article",
		Contents:  "This is the test article.",
		UserName:  "yasuko",
		NiceNum:   2,
		CreatedAt: time.Now(),
	}
)
