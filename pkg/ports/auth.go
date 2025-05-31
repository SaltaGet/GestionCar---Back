package ports

type AuhtService interface {
	AuthLogin(username string , password string) (token string, err error)
}

type AuthRepository interface {
	AuthLogin(username string, password string) (token string, err error)
}
