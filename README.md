# Packer Docker build post-processor

[![Build Status](https://travis-ci.org/jgkim/packer-post-processor-docker-dockerfile.svg)](https://travis-ci.org/jgkim/packer-post-processor-docker-dockerfile)

This is a [Packer](http://packer.io/) post-processor plugin which allows setting Docker metadata on an artifact from the [Docker builder](https://packer.io/docs/builders/docker.html) that was committed.

Normally, Docker images built using Packer cannot include user, environment variables and other metadata that is available in Dockerfiles.

This plugin will automatically create a temporary Dockerfile and run `docker build` in an annonymous context. Most Dockerfile instructions are supported as json parameters, note that `RUN`, `ADD`, `COPY`, and `ONBUILD` are not supported because packer provisioners should be used for their functionality.

## Usage

In your packer template, configure the post processor:

    {
      ...
      "post-processors": [{
        "type": "docker-dockerfile",
        "maintainer": "James G. Kim <jgkim@jayg.org>",
        "cmd": ["-v"],
        "label": {
          "version": "1.0"
        },
        "expose": [8080],
        "env": {
          "NAME": "James G. Kim"
        },
        "entrypoint": ["/bin/bash"],
        "user": "jgkim",
        "volume": ["/home/jgkim"]
      }]
      ...
    }

Values can include user variables and other Packer functions as documented on the [Packer manual](https://packer.io/docs/templates/user-variables.html).

`cmd` and `entrypoint` can have either array or string values, this mirrors the Dockerfile format and functionality; See the [Dockerfile reference](http://docs.docker.com/reference/builder/) for details.

Please note that if you are using the `docker-tag` post processor to tag the resulting artifact of this post processor then you must put both post processor on the same chain:

    {
      ...
      "post-processors": [
        [{
            "type": "docker-dockerfile",
            "maintainer": "James G. Kim <jgkim@jayg.org>",
            "volume": ["/var/db", "/var/log"]
        }, {
            "type": "docker-tag",
            "repository": "jgkim/image",
            "tag": "latest"
        }]
      ]
      ...
    }

## Building

Install the necessary dependencies by running `go get -d ./...` and then just type `go test ./...`. This will compile some more dependencies and then run the tests. If this exits with exit status 0, then everything is working!

    $ go get -d ./...
    ...
    $ go test ./...
    ...

To compile the Packer plugin, run `go build`.

    $ go build

Put the binary `packer-post-processor-docker-dockerfile` into the `bin` directory.

## Acknowledgement

This plugin has been rewritten from scratch, but still heavily based on [Avishai Ish-Shalom](https://github.com/avishai-ish-shalom)'s [original work](https://github.com/avishai-ish-shalom/packer-post-processor-docker-dockerfile).

## License

This plugin is released under the Apache License, Version 2.0.

## Support

Please file an issue on the github repository if you think anything isn't working properly or an improvement is required.

This plugin has been tested with the development version (0.8) of Packer.
