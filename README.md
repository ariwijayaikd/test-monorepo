# 🧰 Monorepo Backend Services – Go & Node.js

This repository demonstrates a simple **monorepo** setup for two backend services:

- 🌐 A service written in **Golang**
- ⚡ A service written in **Node.js (Express)**

These services are containerized and intended to be deployed to a **Kubernetes** cluster using CI/CD automation.

---

## 📁 Folder Structure

```
├── go/
│   ├── main.go
│   ├── main_test.go
│   ├── go.mod
│   └── go.Dockerfile
├── node/
│   ├── index.js
│   ├── index.test.js
│   ├── package.json
│   └── node.Dockerfile
├── k8s/                # Kubernetes manifests for deployment
│   ├── cluster-issuer.yaml
│   ├── deployment.yaml
│   ├── ingress.yaml
│   ├── namespace.yaml
│   └── service.yaml   
├── docker-compose.yml
└── .github/
    └── workflows/
        └── build-script.yaml
```


## 🚀 Test Overview

### ✅ 1. Sample Applications

Each service implements a simple "Hello World" API:

- **Go service** (port `8080`)
- **Node.js service** (port `3000`)

You can start both locally using Docker Compose:

```bash
docker-compose up --build
```

### ✅ 2. CI/CD Pipeline

A GitHub Actions workflow automates the following tasks:

- Run tests:
    - Run tests for each service
- Build Docker images:
    - Only run after tests have passed.
    - Only upload to the registry if the event is a push resulting from a merged pull request, not for merge requests themselves.
    - Each service has a unique tag to distinguish between the Node.js and Go apps.
- Push images to:
    - GitHub Container Registry (`ghcr.io`), this repository package registry
    - Dockerhub, in this case I use personal account [ariwijayaikd](https://hub.docker.com/u/ariwijayaikd)
- Deploy using Kubernetes manifests (located in `/k8s`)


### ✅ 3. Access Service

Service can be access at below url:

- Golang
    https://ariwijayaikd.me/monorepo/go

- Modejs
    https://ariwijayaikd.me/monorepo/node


### ✅ 4. Confirmation

To ensure the services are running with the latest image, you can access the [Go](https://ariwijayaikd.me/monorepo/go) service URL.