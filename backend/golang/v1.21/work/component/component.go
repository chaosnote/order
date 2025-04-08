package component

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"

	"idv/chris/utils"
)

type Store interface {
	utils.Log

	SQLStore
	RedisStore
}

type store struct {
	utils.Log

	db *sql.DB
	rc *redis.Client
}

type Config struct {
	Debug    bool
	LogDir   string
	LogLevel int

	RASPath    string
	RASMax     int
	RASReplace bool

	DBAccount  string
	DBPassword string
	DBURL      string
	DBName     string
	DBDriver   string

	RedisURL string
	RedisDB  int
}

func DefaultConfig() Config {
	return Config{
		Debug:    true,
		LogDir:   "./dist/log",
		LogLevel: 0,

		RASPath:    "./dist/rsa/_pem",
		RASMax:     512,
		RASReplace: true,

		DBAccount:  "chris",
		DBPassword: "123456",
		DBURL:      "192.168.0.105:3306",
		DBName:     "simulate",
		DBDriver:   "mysql",

		RedisURL: "192.168.0.105:6379",
		RedisDB:  0,
	}
}

func New(c Config) Store {
	const msg = "componsnts.new"

	var _log utils.Log
	if c.Debug {
		_log = utils.GenComplex(c.LogDir, c.LogLevel, c.Debug)
	} else {
		_log = utils.GenFile(c.LogDir, c.LogLevel)
	}

	utils.RSAInit(c.RASPath, c.RASMax, c.RASReplace)

	var e error
	var _db *sql.DB
	// 例 : "user:password@tcp(ip)?parseTime=true/dbname"
	cmd := fmt.Sprintf(`%s:%s@tcp(%s)/%s?parseTime=true`, c.DBAccount, c.DBPassword, c.DBURL, c.DBName)
	_log.Logger().Debug(msg, zap.String("source", cmd))
	_db, e = sql.Open(c.DBDriver, cmd)
	if e != nil {
		panic(e)
	}
	e = _db.Ping()
	if e != nil {
		panic(e)
	}
	_db.SetConnMaxLifetime(5 * time.Minute)

	var _rc *redis.Client
	_rc = redis.NewClient(&redis.Options{
		Addr:     c.RedisURL,
		DB:       c.RedisDB, // use default DB
		Password: "",        // no password set
	})

	e = _rc.Ping().Err()
	if e != nil {
		panic(e)
	}

	return &store{
		Log: _log,

		db: _db,
		rc: _rc,
	}
}
