### Goal 3

> Build function as image

Using Docker to compile and build docker image for runtime

> You can change broker if needed at the beginning in `runtime.go` file

To build the image, you can run the command below:

```bash
docker build -t runtime .
```

To start the container, you run it by:

> because `/dev` almost exist in every linux container,
> so it will be great to mount your path of dev to `/dev2` in container 

```bash
docker run --rm -v /path/to/dev:/dev2 kaixindeken/runtime -i=input -o=output -s=reverse
```

With input topic `input`, output topic `output` and consuming script `reverse`.

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
