package dsn

import (
	"fmt"
	"net/url"
)

type Dsn struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Database  string `json:"database"`
	Charset   string `json:"charset"`
	ParseTime bool   `json:"parse_time"`
}

type DsnOption func(dsn *Dsn)

func DsnWithHost(host string) DsnOption {
	return func(dsn *Dsn) {
		dsn.Host = host
	}
}

func DsnWithPort(port int) DsnOption {
	return func(dsn *Dsn) {
		dsn.Port = port
	}
}

func DsnWithUsername(username string) DsnOption {
	return func(dsn *Dsn) {
		dsn.Username = username
	}
}

func DsnWithPassword(password string) DsnOption {
	return func(dsn *Dsn) {
		dsn.Password = password
	}
}

func DsnWithDatabase(database string) DsnOption {
	return func(dsn *Dsn) {
		dsn.Database = database
	}
}

func DsnWithCharset(charset string) DsnOption {
	return func(dsn *Dsn) {
		dsn.Charset = charset
	}
}

func DsnWithParseTime(parseTime bool) DsnOption {
	return func(dsn *Dsn) {
		dsn.ParseTime = parseTime
	}
}

func NewDsn(opts ...DsnOption) *Dsn {
	dsn := &Dsn{
		Host:      "127.0.0.1",
		Port:      3306,
		Username:  "default",
		Password:  "default",
		Database:  "test",
		Charset:   "utf8mb4",
		ParseTime: true,
	}

	for _, op := range opts {
		op(dsn)
	}

	return dsn
}

func (dsn *Dsn) ToString() string {
	s := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dsn.Username, dsn.Password, dsn.Host, dsn.Password, dsn.Database)

	u := url.URL{}
	values := url.Values{}
	if dsn.Charset != "" {
		values.Add("charset", dsn.Charset)
	}

	if dsn.ParseTime {
		values.Add("parseTime", "True")
	}

	u.RawQuery = values.Encode()
	return fmt.Sprintf("%s%s", s, u.String())
}
