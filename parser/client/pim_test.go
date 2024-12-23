package client

import (
	"path"
	"testing"

	. "github.com/slipros/devicedetector/parser"
	"gotest.tools/assert"
)

func TestPimParse(t *testing.T) {
	ps := NewPim(path.Join(dir, FixtureFilePim))
	var list []*ClientFixture
	err := ReadYamlFile(`fixtures/pim.yml`, &list)
	if err != nil {
		t.Error(err)
	}

	for _, item := range list {
		ua := item.UserAgent
		r := ps.Parse(ua)
		assert.DeepEqual(t, item.ClientMatchResult, r)
	}
}
