# runtime-systems &middot; [![Maintainability](https://api.codeclimate.com/v1/badges/ca4d743c426d04643da2/maintainability)](https://codeclimate.com/github/estensen/runtime-systems/maintainability)

## Getting started
The project requires go, docker and docker-compose.

## Run
Start the app by running:
```
$ docker-compose up -d
```

## Project Description
In this project we are going to find 5 go programs in open source and turn them into a benchmark suite. Then we are going to run 2 inputs per program (one short running, one lunger running) and generate data for each. 
To look at the performance of these 5 programs, we want to reverse engineer [StackImpact](https://stackimpact.com/blog/profiling-go-applications-in-production/)'s server side dashboard, so that we can run it locally  and not use or login to their servers. We are not going to use all of the same features as StackImpact, but the main views such as CPU, Memory and Time usage will be implemented.

## Plan
The first two-three weeks will be used to look more into StackImpact, benchmarking, and finding our 5 programs. Then we are going to start benchmarking our programs and trying to visualize this by reverse engineering some of StackImpact's features.

## Profiling do's and don't's
Before you profile, you must have a stable environment to get repeatable results.
* The machine must be idle - don't profile on shared hardware or browse the web while waiting for a long benchmark to run.
* Watch out for power savings and thermal scaling.
* Avoid virtual machines and shared cloud hosting; they have too much noise for consistent measurements.
* Have a before and after sample and run them multiple times to get consistent results.
* Run one profile at a time, so you measure the program, and not yourself.

## Profile a function
Use the testing package

## Profile whole programs
```
import "github.com/pkg/profile"

func main() {
    defer profile.Start().Stop()
    ...
}
```

## pprof
If your program runs a webserver you can enable debugging over http.
```
go tool pprof /path/to/binary /path/to/profile
```
### CPU
When CPU profiling is enabled, the runtime will interrupt itself every 10ms and record the stack trace of the currently running goroutines.
The more times a function appears in the profile, the more time that code path is taking as a percentage of the total runtime.
### Memory
Memory profiling records the stack trace when a heap allocation is made.
Memory profiling like CPU is sample based. By default memory profiling samples 1 in every 1000 allocation.
Because memory profiling is sample based and tracks allocations, not use, it's difficult to determine the app's overall memory usage.