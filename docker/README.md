## Dockerisation

This used to be deployed in a simple docker container, but with the removal of `docker-compose` it is no longer worth the hassle.

To use docker again, simply:
 - Move the included files back to the root.
 - Move the ports to the `8` prefixed version as mapped in the `Dockerfile`
 - Build it with docker
 - Run the container