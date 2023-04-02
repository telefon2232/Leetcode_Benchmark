# Leetcode Benchmark
This is the first benchmark based on Leetcode tasks. Used to test CPU efficiency.

## Usage

### Install
If you don't have Golang installed on your system, you must install it:

```bash
apt install golang-go
```

### Run
For quick test you can use:

```bash
./gotest.sh
```

### Compile
You can compile for any platform (**riscv64, arm64, amd64**). To do this, use:

```bash
./compile.sh riscv64
```

### Profiling
After compile prifile command to create CPU-profile for analyze (**Carefully!** Each test run 10 second and we will get 10*700=7000second run).
```bash
./profiling.sh riscv64
```
If you want check only 1 test, you may use in binary dir:
```bash
./${FILE} -test.bench .  -test.benchtime 10s -test.cpuprofile ${FILE}.out
go tool pprof -text ${FILE}.out
```

 
## Plans
- [x] Create project and run tests
- [x] Create scripts for compile for any platforms
- [ ] Characterize by two metrics: By algorithm and by profile 
- [ ] Fix Printfs
- [ ] Can change inputs
- [ ] Fix some tests without correct benchmark-cycles
- [ ] Automatic profile analysis
