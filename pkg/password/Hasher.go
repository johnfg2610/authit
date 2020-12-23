package password

//Hasher is a interface used to abstract the idea of taking a password string and converting it to a hash as well as validating that hash
//Each implementation of this is specific and can include additional construction variables to enable correct processing
type Hasher interface {
	//HashPassword takes a password hashes it using the chosen implementation and then returns a byte array containing the result(this is to account for some hashing algorithms using byte arrays, we can reduce allocations)
	HashPassword(password []byte) ([]byte, error)

	//ValidatePassword takes the users hash and password and compares them, will return nil if sucessful and a error describing the problem otherwise
	//It is generally not advised to log or return this error as it may undermine the confidence in the system(for instance may allow a attacker to verify a user has an account with the service)
	ValidatePassword(hash []byte, password []byte) error
}
