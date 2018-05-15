# Guide to Profile Go Programs
## The easiest way to profile a Go program is to use the [profile package](https://godoc.org/github.com/pkg/profile)

Add to the start of the main function:
```
defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
```

Build the Go program with:
```
$ go build <program>
```

Run pprof:
```
$ go tool pprof —text ./<compiled_program> <benchmark.pprof> 
```

To generate a PDF call graph run:
```
go tool pprof —pdf ./<compiled_program> <benchmark.pprof> > graph.pdf
```
