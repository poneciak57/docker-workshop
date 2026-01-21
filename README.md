# Docker workshops (road to containerization)

This document serves as a guide for the workshop.

## Book of Content

- [What is Docker?](#what-is-docker)
- [Installing Docker](#installing-docker)
- [Hello World with Docker](#hello-world-with-docker)
- [Valgrind even on macOS (with Docker)](#valgrind-even-on-macos-with-docker)
- [Dockerfile basics](#dockerfile-basics)
- [Docker Compose basics](#docker-compose-basics)
- [Examples](#examples)
  - [SimpleExample](#simpleexample)
  - [ComplicatedExample](#complicatedexample)
  - [ComposeExample](#composeexample)
- [Task: Dockerize a simple web application](#task-dockerize-a-simple-web-application)
- [Task: Dockerize your own project :)](#task-dockerize-your-own-project-)

---

# What is Docker?

Platform for developing, shipping, and running applications. It separates your applications from your infrastructure so you can deliver software quickly.

- **What is a Container?**
  - A standard unit of software that packages up code and all its dependencies.
  - Runs quickly and reliably from one computing environment to another.
  - Unlike VMs, containers share the host OS kernel (lightweight).

- **Why use Docker?**
  - **Consistency**: "It works on my machine" -> It works everywhere.
  - **Speed**: Fast deployment and scaling.
  - **Isolation**: Apps are isolated from each other and the host system.

- **Who uses Docker?**
  - Developers (local env setup, testing).
  - DevOps/SysAdmins (CI/CD, deployment).
  - Everyone who wants to run software without polluting their OS.

# Installing Docker

- **macOS/Windows**: Install Docker Desktop.
- **Linux**: Install Docker Engine.
- Verify: `docker --version`

# Hello World with Docker

Full flow of running a container manually.

1. **Pull the image** (Show explicit pulling):
   ```bash
   docker pull hello-world
   ```

2. **Check images list** (Verify it's there):
   ```bash
   docker images
   ```

3. **Run the container**:
   ```bash
   docker run hello-world
   ```

**What just happened?**
1. Docker client contacts the daemon.
2. Daemon pulls the "hello-world" image from Docker Hub (if not found locally).
3. Daemon creates a new container from that image which runs the executable that produces the output.
4. Daemon streams that output to the Docker client, which sends it to your terminal.

# Valgrind even on macOS (with Docker)

This is a very useful example of running Linux tools on non-Linux OS.
We can also mention **DevContainers** here, as they allow using a container as a full-featured development environment.

# Dockerfile basics

See [cheatsheet.md](cheatsheet.md) for syntax and commands.

# Docker Compose basics

See [cheatsheet.md](cheatsheet.md) for details.

# Examples

## [SimpleExample](SimpleExample/)
Just Nginx hosting HTML files.

## [ComplicatedExample](ComplicatedExample/)
Some server with Nginx behaving as a proxy.

## [ComposeExample](ComposeExample/)
The "ComplicatedExample" but showcasing how Docker Compose simplifies the process.

# Task: Dockerize a simple web application
See the [Task/](Task/) folder.
Your goal is to dockerize the provided application.

# Task: Dockerize your own project :)
Apply what you've learned to your own projects!
