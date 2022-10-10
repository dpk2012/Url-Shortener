package model

func CreateUrlTag(urlTag UrlTag) error {
	tx := db.Create(&urlTag)
	return tx.Error
}
