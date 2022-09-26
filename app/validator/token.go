package validator

import (
	"log"
	"tpk-backend/app/pkg"
)

func CheckStatusToken(t string) bool {
	token, err := pkg.FindToken(t)
	if err != nil {
		log.Println(err)
		return false
	}
	if token.Status != `A` {
		return false
	}
	return true
}
