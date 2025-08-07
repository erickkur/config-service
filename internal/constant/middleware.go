package constant

type VerificationType int

const (
	AppTokenValue VerificationType = 1
)

var VerificationTypeConstants = struct {
	AppToken VerificationType
}{
	AppToken: AppTokenValue,
}
