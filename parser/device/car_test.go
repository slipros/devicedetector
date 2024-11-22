package device

import (
	"path"
	"testing"

	. "github.com/slipros/devicedetector/parser"
	"gotest.tools/assert"
)

func TestCarParse(t *testing.T) {
	ps := NewCar(path.Join(dir, FixtureFileCar))
	var list []*DeviceFixture
	err := ReadYamlFile(`fixtures/car_browser.yml`, &list)
	if err != nil {
		t.Error(err)
	}

	for _, item := range list {
		ua := item.UserAgent
		r := ps.Parse(ua)
		test := item.GetDeviceMatchResult()
		assert.DeepEqual(t, test, r)
	}
}
