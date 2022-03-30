### Home Task-Serverless

* [Goal 1](/Goal-1)
  * Structure 
    ```bash
    Goal-1
    ├── README.md
    └── functions
    ├── addEM.sh
    ├── length.sh
    └── reverce.sh
    ```
  * Usage
    ```bash
    sh addEM.sh message
    ```
    [Click here](/Goal-1/README.md) to reach the detail of usage

    
* [Goal 2](/Goal-2)
    * Structure
      ```bash
      Goal-2
      ├── README.md
      ├── functions
      │   ├── addEM.sh
      │   ├── length.sh
      │   └── reverse.sh
      ├── go.mod
      ├── go.sum
      ├── runtime
      └── runtime.go
      ```

    * Usage
      ```bash
      ./runtime_linux_amd64 -i=input -o=output -s=reverse
      ```
      [Click here](/Goal-2/README.md) to reach the detail of usage


* [Goal 3](/Goal-3)

  * Structure
    ```bash
     Goal-3
     ├── Dockerfile
     ├── README.md
     ├── functions
     │   ├── addEM.sh
     │   ├── length.sh
     │   └── reverse.sh
     ├── go.mod
     ├── go.sum
     ├── runtime-stateful.yaml
     └── runtime.go
     ```
  * Usage
    ```bash
    docker run --rm -v /path/to/dev:/dev2 kaixindeken/runtime -i=input -o=output -s=reverse
    ```
    ```bash
    kubectl apply -f runtime-stateful.yaml
    ```
    [Click here](/Goal-3/README.md) to reach the detail of usage


* [Goal 4](/Goal-4)

  ```bash
  Goal-4
  └── README.md
  ```