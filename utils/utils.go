package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"lms_backend/database"
	"lms_backend/global"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"unicode"

	"github.com/craftzbay/go_grc/v2/converter"
	"github.com/craftzbay/go_grc/v2/data"
	"gorm.io/datatypes"
)

func GetTableName(name ...interface{}) string {
	var tableName string
	var alias string
	for idx, v := range name {
		switch idx {
		case 0:
			tableName = v.(string)
		case 1:
			alias = v.(string)
		}
	}
	tName := global.Conf.DBSchema + "." + global.Conf.DBTablePrefix + "_" + tableName
	if alias != "" {
		tName += " as " + alias
	}
	return tName
}

func DateFilter(req *map[string]string) (string, string) {
	filters := *req
	// dates from {start_date} to {end_date} ---> 2023-07-05 - 2023-07-06
	startStr := filters["start_date"]
	endStr := filters["end_date"]
	// last {duration} days	---> {duration} == 7 ? 2023-06-29 - 2023-07-05
	durationStr := filters["duration"]
	// days in {year}-{month} ---> {year} == 2023 & {month} == 7 ? 2023-07-01 - 2023-07-31
	yearStr := filters["year"]
	monthStr := filters["month"]

	// no filter if none of the above is given
	if startStr == "" && endStr == "" && durationStr == "" && yearStr == "" && monthStr == "" {
		return "", ""
	}
	// conversions
	startTime := converter.DateStringToTime(startStr)
	endTime := converter.DateStringToTime(endStr)
	duration := converter.StringToInt(durationStr)
	year := converter.StringToInt(yearStr)
	month := converter.StringToInt(monthStr)

	if year == 0 {
		year = time.Now().Year()
	}
	if month == 0 {
		month = int(time.Now().Month())
	}
	if duration == 0 {
		if startTime.IsZero() && endTime.IsZero() {
			startTime = time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Now().Location())
			endTime = time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.Now().Location())
		} else {
			if startTime.IsZero() {
				startTime = time.Now()
			}
			if endTime.IsZero() {
				endTime = time.Now()
			}
		}
	} else {
		endTime = time.Now()
		startTime = endTime.Add(-time.Hour * 24 * time.Duration(duration))
	}

	startTimeStr := data.LocalDate(startTime).String() + " 00:00:00"
	endTimeStr := data.LocalDate(endTime).String() + " 23:59:59.999999"

	return startTimeStr, endTimeStr
}

func StrToIntArray(str string) []uint {
	res := make([]uint, 0)
	strJson := datatypes.JSON(str)
	intArray := make([]uint, 0)

	if err := json.Unmarshal(strJson, &intArray); err != nil {
		return res
	}
	res = RemoveDuplicates(intArray, func(a, b uint) bool { return a == b })

	return res
}

func MapToSlice[T any](m *map[interface{}]T) []T {
	req := *m
	res := make([]T, 0)
	for _, val := range req {
		res = append(res, val)
	}
	return res
}

func RemoveDuplicates[T any](list []T, cmp func(a, b T) bool) []T {
	res := make([]T, 0)

	for _, elem := range list {
		if !Contains(res, elem, cmp) {
			res = append(res, elem)
		}
	}

	return res
}

func Contains[T any](list []T, val T, cmp func(a T, b T) bool) bool {
	for _, x := range list {
		if cmp(x, val) {
			return true
		}
	}
	return false
}

func DayStartTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func DayEndTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 1e9-1, t.Location())
}

func CheckExists(table string, columns []string, values []interface{}) bool {
	var exists bool
	if len(columns) != len(values) {
		return false
	}
	db := database.DBconn
	tx := db.Table(GetTableName(table)).Select("count(*) > 0")

	for i := range columns {
		tx.Where(columns[i]+" = ?", values[i])
	}
	if err := tx.Take(&exists).Error; err != nil {
		return false
	}
	return exists
}

func RoundFloat(v float64) float64 {
	val := math.Floor(v*100) / 100
	return val
}

func InterfaceToStruct[T any](m any) (T, error) {
	var responseStruct T
	jsonString, err := json.Marshal(m)
	if err != nil {
		return responseStruct, err
	}
	err = json.Unmarshal(jsonString, &responseStruct)
	return responseStruct, err
}

func ToString(i interface{}) string {
	switch s := i.(type) {
	case string:
		return s
	case bool:
		return strconv.FormatBool(s)
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32)
	case int:
		return strconv.Itoa(s)
	case int64:
		return strconv.FormatInt(s, 10)
	case int32:
		return strconv.Itoa(int(s))
	case int16:
		return strconv.FormatInt(int64(s), 10)
	case int8:
		return strconv.FormatInt(int64(s), 10)
	case uint:
		return strconv.FormatUint(uint64(s), 10)
	case uint64:
		return strconv.FormatUint(uint64(s), 10)
	case uint32:
		return strconv.FormatUint(uint64(s), 10)
	case uint16:
		return strconv.FormatUint(uint64(s), 10)
	case uint8:
		return strconv.FormatUint(uint64(s), 10)
	case []byte:
		return string(s)
	case nil:
		return ""
	case fmt.Stringer:
		return s.String()
	case error:
		return s.Error()
	default:
		return ""
	}
}

func GetRandomNumber(len int) uint {
	min := int(math.Pow10(len - 1))
	max := int(math.Pow10(len) - 1)
	rand.NewSource(time.Now().UnixNano())
	return uint(rand.Intn(max-min) + min)
}

func CalculateAge(birthDate data.LocalDate) float64 {
	return time.Since(time.Time(birthDate)).Hours() / 24 / 365
}

func ReadBase64Image(base64Img string) ([]byte, string) {
	imageData, err := base64.StdEncoding.DecodeString(base64Img)
	if err != nil {
		return nil, ""
	}
	mimeType := http.DetectContentType(imageData)
	var formatString string
	switch mimeType {
	case "image/jpeg":
		formatString = "jpeg"
	case "image/png":
		formatString = "png"
	case "image/jpg":
		formatString = "jpg"
	}
	return imageData, formatString
}

func IsNumber(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func Find[T any](item T, list []T, cmp func(a, b T) bool) *T {
	for _, elem := range list {
		if cmp(item, elem) {
			return &elem
		}
	}
	return nil
}
