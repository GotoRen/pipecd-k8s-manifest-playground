# pipecd-k8s-manifest-playground

This repository is a Kubernetes manifest for testing PipeCD.

## PipeCD / Piped

- https://github.com/pipe-cd/pipecd

## Run Locally

### Control Plane

```shell
### Start running a Kubernetes cluster.
$ make kind-up

### Install the web dependencies module.
$ make update/web-deps

### Install Control Plane into the local cluster.
$ make run/pipecd

### Once all components are running up, use kubectl port-forward to expose the installed Control Plane on your localhost.
$ kubectl -n pipecd port-forward svc/pipecd 8080
```

### Data Plane

```shell
### Start running a piped.
$ make run/piped CONFIG_FILE=piped.yaml INSECURE=true
```

## Default

### PipeCD Web console

| Key      | Value        |
| :------- | :----------- |
| project  | quickstart   |
| username | hello-pipecd |
| password | hello-pipecd |

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
| Application          | // PipeCD アプリケーションマニフェスト（app.pipecd.yaml）を管理
| Command              | // PipeCD に対するデプロイ命令（コマンド）の実行履歴を管理
| Deployment           | // デプロイ履歴を管理
| DeploymentChain      |
| Event                |
| Piped                | // Piped マニフェスト（pipd.yaml）の情報を管理
| Project              |
+----------------------+
8 rows in set (0.00 sec)
```

## Project setting

| Key         | Value                   |
| :---------- | :---------------------- |
| Name        | pipecd-playground       |
| Description | Build and test locally. |

## PipeCD Web UI

### nginx

![nginx](https://github.com/GotoRen/pipecd-k8s-manifest-playground/assets/63791288/79e54f38-8ba6-4a46-b438-07307c7f0dc2)

### apache (httpd)

![apache](https://github.com/GotoRen/pipecd-k8s-manifest-playground/assets/63791288/56dbeabc-027c-4087-8c6d-6b08c94346a5)
