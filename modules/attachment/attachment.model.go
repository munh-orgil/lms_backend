package attachment

import (
	"lms_backend/database"

	"github.com/craftzbay/go_grc/v2/data"
	"github.com/gofiber/fiber/v2"
)

type Attachment struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"type:varchar(255)"`
	FileName  string         `json:"file_name" gorm:"type:varchar(255)"`
	FileType  string         `json:"file_type" gorm:"type:varchar(20)"`
	CreatedAt data.LocalTime `json:"created_at" gorm:"autoCreateTime"`
}

func RunMigrations() {
	database.DBconn.AutoMigrate(Attachment{})
}

func AttachmentList(c *fiber.Ctx) (res *data.Pagination[Attachment], err error) {
	var totalRows int64
	attachments := make([]Attachment, 0)

	db := database.DBconn
	tx := db.Model(Attachment{})
	tx.Count(&totalRows)

	p := data.Paginate[Attachment](c, totalRows)

	if err = tx.Offset(p.Offset).Limit(p.PageSize).Find(&attachments).Error; err != nil {
		return
	}
	p.Items = attachments

	res = p

	return
}

func (a *Attachment) Create() error {
	db := database.DBconn
	return db.Create(&a).Error
}

func (a *Attachment) Update() error {
	db := database.DBconn
	return db.Updates(&a).Error
}

func (a *Attachment) Delete() error {
	db := database.DBconn
	return db.Delete(&a).Error
}
