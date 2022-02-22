package comment

import (
	"strconv"

	con "github.com/galigaribaldi/PersonalBooks/pkg/controller"
	models "github.com/galigaribaldi/PersonalBooks/pkg/models"
)

func init() {
	con.ConnectDatabase()
	con.DB.AutoMigrate()
	con.DB.AutoMigrate(&models.CommentAuthor{})
	con.DB.AutoMigrate(&models.CalificationAuthor{})
}

func CreateCommentAuthor(commentAuthor models.CommentAuthor) {
	con.DB.Create(&commentAuthor)
}

func SelectAllCommentAuthor() []models.CommentAuthor {
	var comments []models.CommentAuthor
	con.DB.Find(&comments)
	return comments
}

func SelectByIdCommentAuthor(ids int) []models.CommentAuthor {
	idString := strconv.Itoa(ids)
	var newCommentAuthor []models.CommentAuthor
	if result := con.DB.Where("id = ?", idString).Find(&newCommentAuthor); result.Error != nil {
		return newCommentAuthor
	}
	return newCommentAuthor
}

func DeleteCommentAuthor(ids int) int {
	var newCommentAuthor models.CommentAuthor
	if result := con.DB.Delete(&newCommentAuthor, ids); result.Error != nil {
		return 0
	}
	return ids
}

func UpdateCommentAuthor(commentAuthor models.CommentAuthor) {
	con.DB.Model(&commentAuthor).Updates(models.CommentAuthor{
		Comment:             commentAuthor.Comment,
		Personal_Experience: commentAuthor.Personal_Experience,
		AuthorID:            commentAuthor.AuthorID,
	})
}
