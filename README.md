# runtime-systems
## Project Description
In this project we are going to find 5 go programs in open source and turn them into a benchmark suite. Then we are going to run 2 inputs per program (one short running, one lunger running) and generate data for each. 
To look at the performance of these 5 programs, we want to reverse engineer [StackImpact](https://stackimpact.com/blog/profiling-go-applications-in-production/)'s server side dashboard, so that we can run it locally  and not use or login to their servers. We are not going to use all of the same features as StackImpact, but the main views such as CPU, Memory and Time usage will be implemented.

# Plan
The first two-three weeks will be used to look more into StackImpact, benchmarking, and finding our 5 programs. Then we are going to start benchmarking our programs and trying to visualize this by reverse engineering some of StackImpact's features.