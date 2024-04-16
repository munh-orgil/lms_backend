package log

import (
	"encoding/json"
	"fmt"
	"lms_backend/database"
	"lms_backend/utils"
	"strconv"
	"strings"
	"time"

	"github.com/craftzbay/go_grc/v2/client"
	"github.com/craftzbay/go_grc/v2/data"
	"github.com/gofiber/fiber/v2"
	"gorm.io/datatypes"
)

type Log struct {
	Id             uint           `json:"id" gorm:"primaryKey"`
	Path           string         `json:"path" gorm:"type:varchar(1000)"`
	Method         string         `json:"method" gorm:"type:varchar(10)"`
	HttpStatusCode int            `json:"http_status_code"`
	Duration       float64        `json:"duration"`
	ReqBody        datatypes.JSON `json:"-"`
	ResBody        datatypes.JSON `json:"-"`
	CreatedAt      data.LocalTime `json:"created_at" gorm:"autoCreateTime"`
	CreatedBy      uint           `json:"-"`
}

func RunMigrations() {
	db := database.DBconn
	db.AutoMigrate(Log{})
}

func (s *Log) Save() error {
	db := database.DBconn
	return db.Save(&s).Error
}

func FiberLogSaver(c *fiber.Ctx, logString []byte) {
	log := new(Log)
	log.Path = string(c.Request().RequestURI())
	if utils.Contains(strings.Split(log.Path, "/"), "file", func(a, b string) bool { return a == b }) {
		return
	}
	startTimeStr := c.Get("X-Request-Start-Time")
	startTime, _ := strconv.ParseInt(startTimeStr, 0, 0)
	log.HttpStatusCode = c.Response().StatusCode()
	log.Method = string(c.Request().Header.Method())
	log.Duration = float64(time.Now().UnixMicro()-startTime) / 1000000
	reqBody := c.Request().Body()
	if len(reqBody) < 100*1024 {
		log.ReqBody = c.Response().Body()
	}
	resBody := c.Response().Body()
	if len(resBody) < 100*1024 {
		log.ResBody = c.Response().Body()
	}
	log.Save()
}

func RequestLogSaver(path, method string, duration float64, req, res interface{}, err *client.RequestError) {
	var (
		reqBody []byte
		resBody []byte
	)
	reqBody, _ = json.Marshal(req)
	log := Log{
		Method:   method,
		Path:     path,
		Duration: duration,
	}
	if err != nil {
		resBody = datatypes.JSON(fmt.Sprintf("{\"message\": \"%s\"}", err.Error()))
		log.HttpStatusCode = err.StatusCode
	} else {
		resBody, _ = json.Marshal(res)
		log.HttpStatusCode = 200
	}
	log.ReqBody = reqBody
	log.ResBody = resBody
	go log.Save()
}
