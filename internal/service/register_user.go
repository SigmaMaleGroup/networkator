package service

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"math/big"

	errs "github.com/SigmaMaleGroup/networkator/internal/customErrors"
	"github.com/SigmaMaleGroup/networkator/internal/models"
	"github.com/SigmaMaleGroup/networkator/internal/tokens"
)

func (s service) RegisterUser(ctx context.Context, credits *models.RegisterCredentials) (string, error) {
	exists, err := s.storage.CheckDuplicateUser(ctx, credits.Email)
	if err != nil {
		return "", err
	}

	if exists {
		return "", errs.ErrCredentialsInUse
	}

	passwordSalt, err := RandSymbols(10)
	if err != nil {
		return "", err
	}

	passwordHash := mdHash(credits.Password, passwordSalt)

	userID, err := s.storage.CreateUser(ctx, credits.Email, passwordHash, passwordSalt, credits.IsRecruiter)
	if err != nil {
		return "", err
	}

	if userID <= 0 {
		return "", errors.New("returned userID 0")
	}

	return tokens.GenerateJWT(userID, credits.Email, models.GetUserType(credits.IsRecruiter))
}

// mdHash hashes password with salt
func mdHash(password, passwordSalt string) string {
	hasher := md5.New()
	io.WriteString(hasher, password+passwordSalt)

	return hex.EncodeToString(hasher.Sum(nil))
}

// RandSymbols returns n-char string of random characters
func RandSymbols(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}
