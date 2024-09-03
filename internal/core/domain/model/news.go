package model

// reform:News
//
//go:generate reform
type News struct {
	ID         int64  `reform:"Id,pk"` // primary key
	Title      string `reform:"Title"`
	Content    string `reform:"Content"`
	Categories []int64
}
