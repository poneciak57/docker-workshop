# Basic Docker Commands Cheatsheet

## Container Management
| Command | Description |
|---------|-------------|
| `docker run [OPTIONS] IMAGE [COMMAND] [ARG...]` | Run a command in a new container |
| `docker stop [OPTIONS] CONTAINER [CONTAINER...]` | Stop one or more running containers |
| `docker rm [OPTIONS] CONTAINER [CONTAINER...]` | Remove one or more containers |
| `docker ps [OPTIONS]` | List running containers |
| `docker ps -a` | List all containers (running and stopped) |
| `docker logs [OPTIONS] CONTAINER` | Fetch the logs of a container |
| `docker exec [OPTIONS] CONTAINER COMMAND [ARG...]` | Run a command in a running container |
| `docker inspect [OPTIONS] NAME\|ID [NAME\|ID...]` | Return low-level information on Docker objects |

## Image Management
| Command | Description |
|---------|-------------|
| `docker build [OPTIONS] PATH \| URL \| - ` | Build an image from a Dockerfile |
| `docker pull [OPTIONS] NAME[:TAG\|@DIGEST]` | Pull an image or a repository from a registry |
| `docker push [OPTIONS] NAME[:TAG]` | Push an image or a repository to a registry |
| `docker images [OPTIONS]` | List images |
| `docker rmi [OPTIONS] IMAGE [IMAGE...]` | Remove one or more images |

## Volume Management
| Command | Description |
|---------|-------------|
| `docker volume ls` | List all volumes |
| `docker volume create [OPTIONS] [VOLUME]` | Create a volume |
| `docker volume rm VOLUME [VOLUME...]` | Remove one or more volumes |
| `docker volume inspect VOLUME [VOLUME...]` | Display detailed information on one or more volumes |

## Network Management
| Command | Description |
|---------|-------------|
| `docker network ls` | List all networks |
| `docker network create [OPTIONS] NETWORK` | Create a network |
| `docker network rm NETWORK [NETWORK...]` | Remove one or more networks |
| `docker network inspect NETWORK [NETWORK...]` | Display detailed information on one or more networks

## Compose Commands
| Command | Description |
|---------|-------------|
| `docker compose up [OPTIONS]` | Create and start containers defined in a docker-compose.yml file |
| `docker compose down [OPTIONS]` | Stop and remove containers, networks, images, and volumes defined in a docker-compose.yml file |
| `docker compose ps` | List containers defined in a docker-compose.yml file |
| `docker compose logs [OPTIONS] [SERVICE...]` | View output from containers defined in a docker-compose.yml file |
| `docker compose build [OPTIONS] [SERVICE...]` | Build or rebuild services defined in a docker-compose.yml file |

# Most usefull commands

## Build and start containers in detached mode, removing orphaned containers with a specific compose file
```bash
docker compose -f <filename> up -d --build --remove-orphans
```

## Connect to a running container's shell
```bash
docker exec -it <container_name_or_id> /bin/bash
```

## View real-time logs of a specific container
```bash
docker logs -f <container_name_or_id>
```

## List all Docker containers (running and stopped)
```bash
docker ps -a
```

## Run a container with port mapping and volume mounting
```bash
docker run -d -p <host_port>:<container_port> -v <host_directory>:<container_directory> <image_name>
```

## Run a container connecting to its shell
```bash
docker run -it <image_name> /bin/bash
```

# Dockerfile Basic Structure

```Dockerfile
# Use an official base image
FROM <base_image>:<tag>

# Set environment variables
ENV <key>=<value>

# Install dependencies
RUN apt-get update && apt-get install -y <package_name>

# Copy application files
COPY <source_path> <destination_path>

# Set the working directory
WORKDIR <working_directory>

# Expose ports
EXPOSE <port_number>

# Specify volumes
VOLUME ["/data"]

# Define the command to run the application
CMD ["<executable>", "<arg1>", "<arg2>"]

# Define the entry point for the container
ENTRYPOINT ["<entrypoint_executable>", "<arg1>", "<arg2>"]

# Diference between CMD and ENTRYPOINT
# CMD sets default command and/or parameters, which can be overridden when running the container.
# ENTRYPOINT sets the command and parameters that will always be executed when the container starts.

```


# Layering in Docker Images
- Each instruction in a Dockerfile creates a new layer in the image.
- Layers are cached, so if a layer hasn't changed, Docker can reuse it during builds, speeding up the process.
- Minimizing the number of layers and combining related commands can lead to smaller and more efficient images.
- Use multi-stage builds to reduce the final image size by separating build-time dependencies from runtime dependencies.

```Dockerfile
# Example of multi-stage build
FROM golang:1.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o myapp

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/myapp .
CMD ["./myapp"]
```