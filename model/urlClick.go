package model

func CreateUrlClick(urlClick UrlClick) error {
	tx := db.Create(&urlClick)
	return tx.Error
}
