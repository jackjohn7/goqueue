# Goqueue

I attempted to write a message queue in Rust, and I wasn't having a great time. I wanted to
see if I would have a better experience in Go where I can just pass pointers around and
disregard enforced memory safety. This system is VERY basic, and there are certainly bugs.

# TO-DO

The following are things I'd like to have done before this project is considered useful.

- [ ] Docker image
  - [X] Core functionality
  - [ ] Published to docker.io
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
  - [ ] Security
    - [ ] Encryption
    - [ ] Admin account
    - [ ] Ability to define users with access to particular topics
  - [ ] Predefined topics with strict access
- [ ] Documentation
  - [ ] Web documentation
  - [ ] Code documentation (pkg.go.dev)

