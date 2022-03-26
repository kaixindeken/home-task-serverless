### Goal 2

> Build a BASH runtime

Use pulsar-client-go to achieve producer and consumer. With the usage of coroutine, the parallel running of producer and consumer become possible.

Also, the usage of `cli`, a golang package to get commands more standard. 

> You can change pulsar url if needed at main function, in `runtime.go` file

To compile it, you can run the command below:

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o runtime *.go
```

When the compile of runtime is done, you can create a consumer:

```bash
./runtime -c -t=test -s=reverse
```

or a producer:

```bash
./runtime -p -t=test
```

The command above is to produce a "please" message to "test" topic and consumed with the function "reverse" and it can be seen in terminal. 

To learn more about `runtime`, please run the command below:

```bash
./runtime -h
```