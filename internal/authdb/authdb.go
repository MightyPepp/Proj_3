package authdb

import (
	"golang.org/x/crypto/bcrypt"
)

var userPasswords = map[string][]byte {
	"joe": 	[]byte("$2a$12$XCGBCJJ7N2WBrMGDxsLPcO59dYEPBdsmVHof9//CdFhRuJOxM4UYi"), 	//Actual password is "12345678"
	"mary": []byte("$2a$12$wvJAxEiMqolhDurSRRTa..V6diYHs1lxuBBs2gSy92GLQFxt3ibkS"), 	//Actuel password is "A1s135246"
}

func VerifyUserPass(username, password string) bool {
	wantPass, hasUser := userPasswords[username]
	if !hasUser {
		return false
	}
	if cmperr := bcrypt.CompareHashAndPassword(wantPass, []byte(password)); cmperr == nil {
		return true
	}
	return false
}