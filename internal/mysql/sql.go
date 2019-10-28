package mysql

type scanner interface {
	Scan(dest ...interface{}) error
}
