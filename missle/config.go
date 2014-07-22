package missle

var DSN string = "root:abc123@tcp(localhost:3306)/mr?charset=utf8"

const (
	MC = "localhost:11211"
)

func init() {
}

func SetDSN(dsn string) {
	DSN = dsn
}
