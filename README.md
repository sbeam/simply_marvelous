# A simply marvelous toy HTTP server in Go with Docker

This shows how to build a simple go HTTP server inside a Docker container,
which can then be used to deploy Google Compute Engine or AWS's Fargate. Useful
as a small-footprint example for testing.

## Build the image

The `-t` option gives docker way to refer to your image via 'repository:tag',
which is mostly an arbitrary string. The `--platform` option tells Docker what
type of CPU we plan to run this image on. It's important for most deployments
that it matches the architecture of the underlying hardware.

```
$ docker build -t sbeam/simply_marvelous:1.0-amd64 --platform linux/amd64 .
...
=> exporting to image                                                                           1.7s
=> exporting layers                                                                             1.5s
=> writing image sha256:396afd8d10aeb7c71a1bbdf18b83c576ef79805768daf3114c5f633cfc4c7ed7        0.0s
=> naming to docker.io/sbeam/simply_marvelous:1.0-amd64                                         0.0s
```

## Check the images is now known to your local docker daemon

`docker images` lists all the images currently installed on your system.

```
$ docker images | head
REPOSITORY                                                   TAG                         IMAGE ID       CREATED          SIZE
sbeam/simply_marvelous                                       1.0-amd64                   396afd8d10ae   35 seconds ago   308MB
...
```

## Run the image, for testing and fun

`docker run` tells docker to fetch the given image (if not already installed),
create a runnable container from it, and then exectute whatever command was
given in the `CMD` instruction in the Dockerfile _inside_ the container.

  * `--rm` tells Docker to remove the container after it exits. This prevents a lot of unused containers from building up on your system.
  * `--it` tells Docker to start a terminal (TTY) in the running container, and make it interactive, so you can interact with the command running inside the container via your command line on the host.
  * `-p 9999:80` binds the port on your local host to the container's port 80

Note I get a warning about the platform because I'm running this on Apple silicon. (wait, how does Docker translate x86 instructions to ARM? Magic!)

```
$ docker run --rm -it -p 9999:80 sbeam/simply_marvelous:1.0-amd64
WARNING: The requested image's platform (linux/amd64) does not match the detected host platform (linux/arm64/v8) and no specific platform was requested
2021/07/07 18:44:56 serving on port 80
```

in another terminal window, use `curl`, a common command line utility that will make web requests:

```
$ curl http://localhost:9999/hi
Simply Marvelous, Darling!
$ curl http://localhost:9999/echo/isanyonethere
isanyonethere
$ curl http://localhost:9999/more
1
$ curl http://localhost:9999/more
2
```

You should see corresponding log output in the window where you started the server.

## Deploying to AWS Fargate with Terraform

## Deploying to Google Compute Engine with Terraform

