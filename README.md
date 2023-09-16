# Sectran
Sectran provides web-based access control and access management, supporting H5 access for SSH, Telnet (RDP support coming soon), along with local tool proxies. It also offers user operation recording and detailed event auditing.

[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](https://makeapullrequest.com)

## Usage

If you want to use an H5 terminal in your Golang code, you should specify `protocol` with "web socket" to start Sectran. Here's a simple example of how to start a WebSocket server in Golang to enable H5 terminal connections:

```shell
./sectran ssh -p="websocket" 
```

If not, you should specify it as TCP.

```shell
./sectran ssh -p="tcp" 
```

Now we haven't implemented any authentication or other business logic on the proxy server side. specify the target server you want to connect with `-p` 

```shell
./sectran ssh -p="websocket" --username"foo" --password="bar" "-t=192.168.31.100:22
```

