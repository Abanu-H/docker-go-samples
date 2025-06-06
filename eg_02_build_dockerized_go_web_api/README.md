# Docker Basics – Go Web API (Multi-Stage Build)
This repository contains a simple Go API to demonstrate how to build and run containerized applications using Docker.
The main goal is to understand how Docker works, with a focus on creating efficient container images using multi-stage builds.
In this example, we build the Go application into a binary, making it suitable for deployment across different isolated environments.

## What this project does
- Provides a minimal Go API with one health check endpoint (`/healtz`)  
- Uses a Docker multi-stage build to create a small, efficient image  
- Runs the app inside a lightweight Alpine Linux container  

## How to use

1. **Clone ther repository**
```bash
git clone https://github.com/Abanu-H/docker-go-samples.git
```
2. **Navigate to the exercise directory**
```bash
cd docker-go-samples/eg_02_build_dockerized_go_web_api
```
3. **Build the Docker image**
```bash
docker build -t build_dockerized_go_web_api .
```
4. **Run the container**
```bash
docker run -p 3000:3000 build_dockerized_go_web_api
```
Server starts at http://localhost:3000.

5. **Test the app**

```bash
curl http://localhost:3000/healtz
```

We could see:
App is Up and Running Successfully



