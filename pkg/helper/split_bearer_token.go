package helper

import (
	"simple-attendance/pkg/tokenize"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetIDUserFromToken(c *fiber.Ctx) string {
	header := c.GetReqHeaders()
	reqToken := header["Authorization"]
	token := SplitBearer(reqToken)
	idUser, _ := tokenize.GetIdUserFromToken(token)

	return idUser
}

func SplitBearer(token string) string {
	splitToken := strings.Split(token, "Bearer ")
	token = splitToken[1]
	return token
}
