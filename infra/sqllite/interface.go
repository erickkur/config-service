package sqllite

type SqlLiteInfraInterface interface {
	Connect()
	Close()
}
