package repository

type BcryptRepository interface {
	ComparePassword([]byte, string) error
	GenerateHashPassword(string) (string, error)
}
