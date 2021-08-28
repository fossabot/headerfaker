package model

// AppConfig The config struct
type AppConfig struct {
	Port         int  `ini:"port"`
	Release      bool `ini:"release"`
	*MySQLConfig `ini:"mysql"`
}

// Check The check request struct
type Check struct {
	CheckTag  string `json:"checktag"`
	CheckFlag string `json:"checkflag"`
}

// MySQLConfig The DB config struct
type MySQLConfig struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	DB       string `ini:"db"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
}

// Record The DB record table struct
type Record struct {
	ID   int    `gorm:"column:id"`
	UUID string `gorm:"column:uuid"`
	TUID string `gorm:"column:tuid"`
}

// Test The question struct
type Test struct {
	TestTitle       string `json:"testtitle"`
	TestDegree      int    `json:"testdegree"`
	TestDescription string `json:"testdescription"`
	TestUrl         string `json:"testurl"`
	TestTag         string `json:"testtag"`
	TestID          string `json:"testid"`
}

// User The DB user table struct
type User struct {
	ID       int
	UUID     string
	Name     string
	Password string
}
