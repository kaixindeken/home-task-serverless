### Goal 2

> Build a BASH runtime

Use pulsar-client-go to achieve producer and consumer.

Also, the usage of `cli`, a golang package to get commands achieve with more choices. 

> You can change pulsar url if needed at main function, in `runtime.go` file

To compile it for linux, you can run the command below:

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o runtime *.go
```

When compile is done, you can create a consumer:

```bash
./runtime -c -t=test -s=reverse
```
With Topic `test` and consuming script `reverse`.

Or a producer :

```bash
./runtime -p -t=test
```

With Topic `test`.

>The input message was taken in /dev/stdin
>
>The output message was taken in /dev/stdout

To learn more about `runtime`, please run the command below:

```bash
./runtime -h
```