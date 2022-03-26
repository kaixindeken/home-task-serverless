### Goal 3

> Build function as image

Using Docker to compile and build docker image for runtime

> You can change pulsar url if needed at main function, in `runtime.go` file

To build the image, you can run the command below:

```bash
docker build -t runtime .
```

To start the container, you can create a consumer:

```bash
docker run --rm -v /path/to/dev:/dev runtime -c -t=test -s=reverse
```

With Topic `test` and consuming script `reverse`.

Or a producer:

```bash
docker run --rm -v /path/to/dev:/dev runtime -p -t=test
```

With Topic `test`.

>The input message was taken in /dev/stdin
>
>The output message was taken in /dev/stdout

Related image is now uploaded to docker hub, and it can be pulled with command below: 
```bash
docker pull kaixindeken/runtime
```

To run it in kubernetes, you can run the command below:
```bash
kubectl apply -f runtime-stateful.yaml 
```
