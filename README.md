# ETCD-DEMO 🔍

本项目是一个使用 Golang 编写的示例程序，旨在演示如何使用 etcd 的 watch 与 lease 机制进行服务发现与服务注册。在运行本程序之前，您需要先启动一个 etcd 集群。

## 安装依赖

本示例程序依赖以下软件，请先安装它们：
- Golang
- ETCD

## 启动 etcd 集群

在运行本程序之前，您需要启动一个 etcd 集群。您可以参考 [etcd 官方文档](https://etcd.io/docs/v3.5/op-guide/clustering/) 来启动一个 etcd 集群。

## 运行示例程序

在启动了 etcd 集群之后，您可以按照以下步骤来运行本示例程序：

1. 进入`etcd-demo/cmd/discovery`目录，运行以下命令：
    ```sh
    go run main.go
    ```
2. 打开另一个终端窗口，进入`etcd-demo/cmd/registry`目录，运行以下命令：
   ```sh
   go run main.go
   ```
3. 现在，您可以在 `etcd-demo/cmd/discovery` 窗口中看到服务发现的相关信息，而在 `etcd-demo/cmd/registry` 窗口中可以看到服务注册的相关信息

## 注意事项
- 在运行示例程序之前，请确保已经启动了 etcd 集群。
- 本示例程序仅用于演示。
