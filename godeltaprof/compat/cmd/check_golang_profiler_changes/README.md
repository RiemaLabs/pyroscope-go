This command tracks latest commits in src/runtime/mprof.go and src/runtime/pprof
If there are new commits found it creates a PR to update [last_known_go_commits.json](last_known_go_commits.json) file

The idea is that godeltaprof was based on copy of go runtime externals, so if externals change then we may want to change
the godeltaprof as well. 

See [godeltaprof incorrectly scales mutex profile](https://github.com/RiemaLabs/pyroscope-go/issues/47)