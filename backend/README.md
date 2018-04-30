```
$ docker build ./

Get the image id for the build
$ docker images

$ docker run -it -p 8080:8080 -v $(pwd):/go/src/github.com/estensen/runtime-systems/backend [image id]
```

To build for production run:
```
$ docker build ./ --build-arg app_env=production
$ docker run -i -t -p 8080:8080 [image id]
```
