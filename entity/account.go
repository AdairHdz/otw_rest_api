package entity

const (
	_ = iota
	SERVICE_PROVIDER
	SERVICE_REQUESTER
)

type Account struct {
	EntityUUID
	UserID string
	EmailAddress string `gorm:"unique"`
	Password string
	RecoveryCode string
	VerificationCode string
	UserType int
	Verified bool
}