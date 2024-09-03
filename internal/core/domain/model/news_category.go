package model

// reform:NewsCategories
//
//go:generate reform
type NewsCategory struct {
	NewsID     int64 `reform:"NewsId"`
	CategoryID int64 `reform:"CategoryId"`
}
