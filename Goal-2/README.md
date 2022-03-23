### Goal 2

> Build a BASH runtime

Use pulsar-client-go to achieve producer and consumer. With the usage of coroutine, the parallel running of producer and consumer become possible.

Also, the usage of `cli`, a golang package to get commands more standard. 

To compile it, you can run the command below:

```bash
go build runtime.go
```

When the compile of runtime is done, you can run the command like below:

```bash
./runtime -t test -s reverse -m please
```

The command above is to produce a "please" message to "test" topic and consumed with the function "reverse" and it can be seen in terminal. 

To learn more about `runtime`, please run the command below:

```bash
./runtime -h
```