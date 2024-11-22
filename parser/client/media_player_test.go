package client

import (
	"path"
	"testing"

	. "github.com/slipros/devicedetector/parser"
	"gotest.tools/assert"
)

func TestMediaPlayerParse(t *testing.T) {
	ps := NewMediaPlayer(path.Join(dir, FixtureFileMediaPlayer))
	var list []*ClientFixture
	err := ReadYamlFile(`fixtures/mediaplayer.yml`, &list)
	if err != nil {
		t.Error(err)
	}

	for _, item := range list {
		ua := item.UserAgent
		r := ps.Parse(ua)
		assert.DeepEqual(t, item.ClientMatchResult, r)
	}
}
