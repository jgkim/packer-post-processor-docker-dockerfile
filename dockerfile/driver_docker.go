package dockerfile

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"

	"github.com/mitchellh/packer/builder/docker"
)

type DockerDriver struct {
	*docker.DockerDriver
}

func (d *DockerDriver) BuildImage(dockerfile *bytes.Buffer) (string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("docker", "build", "--force-rm=true", "--no-cache=true", "--quiet=true", "-")
	cmd.Stdin = dockerfile
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		return "", err
	}

	if err := cmd.Wait(); err != nil {
		err = fmt.Errorf("Error building image: %s\nStderr: %s",
			err, stderr.String())
		return "", err
	}

	image_id_regexp := regexp.MustCompile("Successfully built ([a-f0-9]+)")
	matches := image_id_regexp.FindStringSubmatch(stdout.String())
	if matches == nil {
		err := fmt.Errorf("Could not parse `docker build` output: %s",
			stdout.String())

		return "", err
	}
	id := matches[len(matches)-1]

	return id, nil
}
