package ports

type DbPort interface {
	CloseDbConnection()
	AddToHistory(result int, operation string) error
}
