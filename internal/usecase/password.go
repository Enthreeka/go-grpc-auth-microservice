package usecase

import (
	"crypto/subtle"
	"encoding/hex"
	"github.com/NASandGAP/auth-microservice/internal/apperror"
	"github.com/NASandGAP/auth-microservice/internal/entity"
	"github.com/NASandGAP/auth-microservice/pkg/validation"
	"golang.org/x/crypto/argon2"
)

// Argon Struct stores parameters for hash function
type Argon struct {
	salt    []byte
	version int
	threads uint8
	time    uint32
	memory  uint32
	keyLen  uint32
}

// TODO Create global salt
// Salt settings for hash generation. Salt = user_id + ...
func (a *Argon) setSalt(salt string) {
	userSalt := []byte(salt)
	a.salt = userSalt
}

// Set parameters for function IDKey
func (a *Argon) setParamsArgon() {
	a.time = 1
	a.memory = 64 * 1024
	a.threads = 4
	a.keyLen = 32
}

// This function generating hash in format argon2. Version = 0x13. Also,
// validation here performs for email and password. Validation here need for
// screening out in case false on early stage.
func generateHashFromPassword(user *entity.User, a *Argon) (string, error) {
	if !validation.IsValidEmail(user.Email) && !validation.IsValidPassword(user.Password) {
		return "", apperror.ErrDataNotValid
	}

	a.setSalt(user.ID)
	a.setParamsArgon()

	hashPasswordByte := argon2.IDKey([]byte(user.Password), a.salt, a.time, a.memory, a.threads, a.keyLen)
	hashPasswordString := hex.EncodeToString(hashPasswordByte)

	return hashPasswordString, nil
}

// Verify serves for compare the entered password and the password in the database
func verifyPassword(hashPassword string, user *entity.User, a *Argon) (bool, error) {
	a.setSalt(user.ID)
	a.setParamsArgon()

	newHashByte := argon2.IDKey([]byte(user.Password), a.salt, a.time, a.memory, a.threads, a.keyLen)

	newHashString := hex.EncodeToString(newHashByte)

	if subtle.ConstantTimeCompare([]byte(hashPassword), []byte(newHashString)) == 0 {
		return false, apperror.ErrHashPasswordsNotEqual
	}

	return true, nil
}
