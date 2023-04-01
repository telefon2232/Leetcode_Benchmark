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
You can compile for any platform (riscv64, arm64, amd64). To do this, use:

```bash
./compile riscv64
```

### Profiling
After compile prifile command to create CPU-profile for analyze:
```bash
./profiling riscv64
```
 
## Plans
- [x] Create project and run tests
- [x] Create scripts for compile for any platforms
- [ ] Characterize by two metrics: By algorithm and by profile 
- [ ] Fix Printfs
- [ ] Fix some tests without correct benchmark-cycles
- [ ] Automatic profile analysis
