### Goal 3

> Build function as image

Using Docker to compile and build docker image for runtime

You may change broker if needed at the beginning in `runtime.go` file

To build the image, you can run the command below:

```bash
docker build -t runtime .
```

To start the container, you run it by:

> You can customize the mount dev folder and config.json file freely

```bash
docker run --rm -v /path/to/dev:/dev2 -v /path/to/config.json:/root/config.json kaixindeken/runtime -c=/root/config.json
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

> The input message was taken in /dev/stdin in container

Related image is now uploaded to docker hub, and it can be pulled with command below: 
```bash
docker pull kaixindeken/runtime
```

To rebuild it, please make sure to get your `broker` and input `filePath` correct

To run it in kubernetes, please make sure to get your params correct, you can run the command below:

```bash
kubectl apply -f runtime-stateful.yaml 
```

To deploy it, please make sure to get your params of mount and config.json in .yaml file correct