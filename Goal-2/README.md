### Goal 2

> Build a BASH runtime

Use pulsar-client-go to achieve producer and consumer.

Also, the usage of `cli`, a golang package to get commands achieve with more choices.

You can use the existed executable file if you platform matched:

```bash
# for linux amd64 platform
./runtime_linux_amd64 -i=input -o=output -s=reverse

# for macos intel platform
./runtime_macos_intel -i=input -o=output -s=reverse
```
With input topic `input`, output topic `output` and consuming script `reverse`.

To compile it, please make sure to get your `broker` and input `filePath` correct:

```bash
go build -a -o runtime *.go
```

When compile is done, you run it:

```bash
# if linux amd64 platform
./runtime -i=input -o=output -s=reverse
```

> The input message was taken in /dev/stdin

To learn more about `runtime`, please run the command below:

```bash
./runtime -h
```