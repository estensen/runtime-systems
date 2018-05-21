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

To run terminal commando from program:
Create a file to save the output in by writing:
```
file, err := os.Create(filename)
if err != nil {
    panic("Could not create " + filename)
}
defer file.Close()
```

Then run the commando and save that commando in a variable:
```

commando := exec.Command("go", "tool", "pprof", "-text", "cpu.pprof")
commandoOutput, err := commando.Output()
if err != nil {
    panic(err)
}
```


Then save the commandoOutput in your textfile
```
file.Write(textOut)
```
