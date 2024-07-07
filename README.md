# Goqueue

I attempted to write a message queue in Rust, and I wasn't having a great time. I wanted to
see if I would have a better experience in Go where I can just pass pointers around and
disregard enforced memory safety. This system is VERY basic, and there are certainly bugs.

# TO-DO

The following are things I'd like to have done before this project is considered useful.

- [ ] Docker image
  - [X] Core functionality
  - [ ] Published to docker.io
  - [X] Test that SSL certificates are easily usable in docker
- [ ] Libraries
  - [ ] Go
    - [ ] Core functionality
    - [ ] Documentation
  - [ ] TypeScript
    - [ ] Core functionality
    - [ ] Documentation
  - [ ] Python
    - [ ] Core functionality
    - [ ] Documentation
  - [ ] Rust
    - [ ] Core functionality
    - [ ] Documentation
- [ ] Configuration
  - [ ] Toml file configuration
  - [X] CLI Arguments
  - [ ] Security
    - [X] Encryption
    - [ ] Simply token-based access
    - [ ] Admin account
    - [ ] Ability to define users with access to particular topics
  - [ ] Predefined topics with strict access
- [ ] Documentation
  - [ ] Web documentation
  - [ ] Code documentation (pkg.go.dev)

# Docker

I don't have the docker image published or anything, so you'll need to build the
container image yourself for now. You can use the following commands.

In any case, you'll need to clone down the project and clone into it.

```bash
docker build -t goqueue -f Dockerfile
```

To run, use the following:

```bash
docker run -p 4173:4173 goqueue
```

## Podman

I use podman, so I've included podman instructions.

```bash
podman build -t goqueue -f Dockerfile
```

```bash
podman run -p 4173:4173 goqueue
```

# Encryption

To enable encryption, you will need to provide a TLS certificate. For local testing,
you can generate your own with the following command:

```bash
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes
```

You can provide the paths to these files via the flags `--cert` and `--key`.

In docker, you can mount your certificate files and pass the configuration options as
follows:

```bash
docker run -v ./key.pem:/key.pem -v ./cert.pem:/cert.pem -p 4174:4173 \
goqueue --cert cert.pem --key key.pem
```

Or with podman:

```bash
podman run -v ./key.pem:/key.pem -v ./cert.pem:/cert.pem -p 4174:4173 \
goqueue --cert cert.pem --key key.pem
```
