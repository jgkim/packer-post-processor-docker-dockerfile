package dockerfile

import (
	"bytes"

	"github.com/mitchellh/packer/builder/docker"
)

// MockDriver is a driver implementation that can be used for tests.
type MockDriver struct {
	*docker.MockDriver

	BuildImageCalled     bool
	BuildImageDockerfile *bytes.Buffer
	BuildImageErr        error
}

func (d *MockDriver) BuildImage(dockerfile *bytes.Buffer) (string, error) {
	d.BuildImageCalled = true
	d.BuildImageDockerfile = dockerfile
	return "1234567890abcdef", d.BuildImageErr
}
