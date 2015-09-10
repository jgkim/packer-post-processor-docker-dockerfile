package dockerfile

import (
	"bytes"

	"github.com/mitchellh/packer/builder/docker"
)

// Driver is the interface that has to be implemented to communicate with
// Docker. The Driver interface also allows the steps to be tested since
// a mock driver can be shimmed in.
type Driver interface {
	docker.Driver

	// Build an image with the given Dockerfile and returns the ID for that image,
	// along with a potential error.
	BuildImage(dockerfile *bytes.Buffer) (string, error)
}
