package client

import (
	"path"
	"testing"

	. "github.com/slipros/devicedetector/parser"
	"gotest.tools/assert"
)

func TestMediaAppParse(t *testing.T) {
	ps := NewMobileApp(path.Join(dir, FixtureFileMobileApp))
	var list []*ClientFixture
	err := ReadYamlFile(`fixtures/mobile_app.yml`, &list)
	if err != nil {
		t.Error(err)
	}

	for _, item := range list {
		ua := item.UserAgent
		r := ps.Parse(ua)
		assert.DeepEqual(t, item.ClientMatchResult, r)
	}
}
