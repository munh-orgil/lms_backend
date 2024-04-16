package session

import (
	"github.com/craftzbay/go_grc/v2/converter"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const (
	IndexUserId = 0
	IndexRoles  = 1
	IndexOrgs   = 2
	IndexOthers = 3
)

type Token struct {
	jwt.RegisteredClaims
	Info [4]interface{} `json:"info"`
}

func GetTokenInfo(c *fiber.Ctx) *Token {
	cs := c.Locals("tokenInfo")
	info, ok := cs.(*Token)
	if !ok {
		return nil
	}
	return info
}

func (tokenInfo *Token) GetUserId() uint {
	if tokenInfo.Info[IndexUserId] != nil {
		val, _ := converter.InterfaceToUint(tokenInfo.Info[IndexUserId])
		return val
	}
	return 0
}

func (tokenInfo *Token) SetUserId(userId uint) {
	tokenInfo.Info[IndexUserId] = userId
}

func (tokenInfo *Token) GetRoles() (roles *[]string) {
	if tokenInfo.Info[IndexRoles] != nil {
		iarr := tokenInfo.Info[IndexRoles].([]interface{})
		sarr := make([]string, len(iarr))
		for i, v := range iarr {
			sarr[i] = v.(string)
		}

		roles = &sarr
	}
	return
}

func (tokenInfo *Token) SetRoles(roles string) {
	tokenInfo.Info[IndexRoles] = roles
}

func (tokenInfo *Token) GetOrganizationIds() (orgIds *[]uint) {
	if tokenInfo.Info[IndexOrgs] != nil {
		converter.MapToStruct(tokenInfo.Info[IndexOrgs], orgIds)
	}
	return
}

func (tokenInfo *Token) SetOrganizationIds(orgIds *[]uint) {
	tokenInfo.Info[IndexOrgs] = orgIds
}

func (tokenInfo *Token) SetOthers(others string) {
	tokenInfo.Info[IndexOthers] = others
}

func (tokenInfo *Token) GetOthers() string {
	if tokenInfo.Info[IndexOthers] != nil {
		return tokenInfo.Info[IndexOthers].(string)
	}
	return ""
}
