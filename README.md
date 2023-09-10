# pipecd-k8s-manifest-playground

K8s manifest for testing PipeCD

## PipeCD / Piped config

- https://github.com/pipe-cd/pipecd

## Run Locally

### Control Plane

```shell
### Install Control Plane into the local cluster.
$ make run/pipecd

### Once all components are running up, use kubectl port-forward to expose the installed Control Plane on your localhost.
$ kubectl -n pipecd port-forward svc/pipecd 8080
```

### Data Plane

```shell
### Starting piped.
$ make run/piped CONFIG_FILE=piped.yaml INSECURE=true
```

## Default

### PipeCD Web console

| Key      | Value        |
| :------- | :----------- |
| project  | quickstart   |
| username | hello-pipecd |
| password | hello-pipecd |

### Project settig

| Key         | Value                   |
| :---------- | :---------------------- |
| Name        | pipecd-playground       |
| Description | Build and test locally. |

### MySQL

| Key                   | Value      |
| :-------------------- | :--------- |
| `MYSQL_ROOT_USER`     | root       |
| `MYSQL_ROOT_PASSWORD` | test       |
| `MYSQL_DATABASE`      | quickstart |

```shell
mysql> show tables;
+----------------------+
| Tables_in_quickstart |
+----------------------+
| APIKey               |
| Application          |
| Command              |
| Deployment           |
| DeploymentChain      |
| Event                |
| Piped                | // piped-config マニフェストの情報を格納
| Project              |
+----------------------+
8 rows in set (0.00 sec)
```
