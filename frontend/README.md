# Frontend

## Interface
Since a part of this project was to make a local visualization of profiling, we have done this by making a frontend using React. The webpage consists of two parts, CPU and Memory, which is the two most important parts for profiling. Currently we have only implemented CPU backend.

### CPU
When looking at the CPU interface, the programs available for profiling is listed. Currently this is sort and fibonacci.

#### Program
When clicking on one of the programs, you have three options for visualising the profiling. 
* PNG graph:
which shows duration, number of samples done thoughout the profiling, and duration/percentage of CPU usage by each function or part of the program
* CPU Graph:
which shows CPU usage in percentage while the program was run. The samples are taken every 50ms. 
* Top 10 functions:
which shows the 10 hot spots of the program. 

### MEMORY
Currently not implemented


