package mongo

type decoder interface {
	Decode(val interface{}) error
}
