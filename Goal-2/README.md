### Goal 2

> Build a BASH runtime

Use channel to go through, from reading file, producing to input topic, to consuming message, and producing to output topic.

Also, the usage of `cli`, a golang package to get commands achieve with more choices.

You can use the existed executable file if you platform matched:

```bash
# for linux amd64 platform
./runtime_linux_amd64 -c=/path/to/config.json

# for macos intel platform
./runtime_macos_intel -c=/path/to/config.json
```

The structure of config.json can refer to config.json.example:

| Parameters      | Detail                                          |
|-----------------|-------------------------------------------------|
| broker          | pulsar url                                      |
| input_file_path | path of message input source                    |
| FunctionRoot    | folder of functions                             |
| input_topic     | topic get input message                         |
| output_topic    | topic get output message                        |
| script          | function in functions folder to consume message |


To compile it, you can run:

```bash
CGO_ENABLED=0 go build -a -o runtime *.go
```

When compile is done, you run it:

```bash
./runtime -c=config.json
```

> The input message was taken in /dev/stdin

To learn more about `runtime`, please run the command below:

```bash
./runtime -h
```