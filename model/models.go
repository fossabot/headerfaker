package model

// Test The question struct
type Test struct {
	TestTitle       string `json:"testtitle"`
	TestDegree      int    `json:"testdegree"`
	TestDescription string `json:"testdescription"`
	TestUrl         string `json:"testurl"`
	TestTag         string `json:"testtag"`
}

// Check The check request struct
type Check struct {
	CheckTag  string `json:"checktag"`
	CheckFlag string `json:"checkflag"`
}
