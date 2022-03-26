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

or a producer:

```bash
docker run --rm -v /path/to/dev:/dev runtime -p -t=test
```


> You also can customize your input and output file path.

Related image is now uploaded to docker hub, and it can be pulled with command below 
```bash
docker pull kaixindeken/runtime
```

To run it in kubernetes, you can run the command below:
```bash
kubectl apply -f runtime-stateful.yaml 
```
