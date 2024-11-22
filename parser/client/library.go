package client

import (
	"path"
)

const ParserNameLibrary = `library`
const FixtureFileLibrary = `libraries.yml`

func init() {
	RegClientParser(ParserNameLibrary,
		func(dir string) ClientParser {
			return NewLibrary(path.Join(dir, FixtureFileLibrary))
		})
}

func NewLibrary(fileName string) *Library {
	c := &Library{}
	c.ParserName = ParserNameLibrary
	if err := c.Load(fileName); err != nil {
		return nil
	}
	return c
}

// Client parser for tool & software detection
type Library struct {
	ClientParserAbstract
}
