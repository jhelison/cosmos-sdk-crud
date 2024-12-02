# Cosmos-SDK - Students CRUD

This project is a Cosmos-SDK module to manage students. The module allows users to create, fetch, or delete students.

For detailed technical documentation, check:

- [`x/README.MD`](x/README.MD).

## Requirements

To run this project you will need:

- [Go v1.21+](https://tip.golang.org/doc/go1.21)

## Getting Started

### Installation

To build and install the binary:

```bash
make install
```

### Tests

To test the project and run unit/integration tests use:

```bash
make test
```

### Initialization

After installation, you can initialize a chain with:

```bash
make init
```

### Running a Simulation

You can also find a small simulation for the module at:

- [RPC](scripts/rps.sh)

To run it, you must first have a chain running with `make init`. After execute:

```bash
rpsd start --minimum-gas-prices 0rps --api.enable true
```

## Commands Overview

### CLI Commands

Here are some useful commands to interact with the module:

- **Query a student by ID**:

  ```bash
  rpsd query rps student [id]
  ```

  Example response:

  ```text
  student:
    age: "20"
    id: rps1u07d74l5t35384af5yxulx28fp9mamrm3rcx5r
    name: alice
  ```

- **Create a student**:

  ```bash
  rpsd tx rps create-student [name] [id] [age] --from [key]
  ```

- **Delete a student**:
  ```bash
  rpsd tx rps delete-student [id] --from [key]
  ```

### REST API Endpoints

- **Fetch a student**:
  ```http
  GET /challenge/rps/student/{id}
  ```
  Example response:
  ```json
  {
    "student": {
      "name": "alice",
      "id": "rps1qmuwmmz3twdrthjkd93w6ycu245na9caqygt02",
      "age": "20"
    }
  }
  ```

## Development Tools

### Linting

You can lint the project with:

```bash
make lint
```

### Protobuf Generation

To generate protobuf files:

```bash
make proto-all
```

---

For additional details on the module's messages, state, and events, see the documentation at [`x/README.MD`](x/README.MD).
