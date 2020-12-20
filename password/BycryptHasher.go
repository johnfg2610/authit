package password

import (
	"golang.org/x/crypto/bcrypt"
)

//BcryptHasher is the default password hasher implemtation using the adapative hashing algorithm bcrypt
type BcryptHasher struct {
	Cost int
}

//NewBcryptHasher creates a new default instance of the bcrypt password hasher using the default cost, this can be changed later and the algorithm should handle this.
func NewBcryptHasher() BcryptHasher {
	return BcryptHasher{
		Cost: bcrypt.DefaultCost,
	}
}

func (bcryptHasher BcryptHasher) HashPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcryptHasher.Cost)
}

func (bcryptHasher BcryptHasher) ValidatePassword(hash []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, []byte(password))
}
