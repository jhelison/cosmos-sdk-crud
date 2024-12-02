# `/x/rps`

## Concept

RPC is a Cosmos-SDK for managing students.
Users can create, fetch or delete students.

# State

This module has a single state:

- Students

## Students

This state stores the information for all the students. The structure of each student is as follow:

- `Name`: The name of the student
- `ID`: The id of the student (the creator address)
- `Age`: The age of the student

The proto definitions are as follow:

```proto
// GenesisState defines the rps module's genesis State
message GenesisState {
  // Student defines the state in the module
  repeated Student students = 1 [ (gogoproto.nullable) = false ];
}

// Student defines the name, age and id object
message Student {
  // name is the name of the student
  string name = 1;
  // id is a unique address for each student (the owner address)
  string id = 2 [ (cosmos_proto.scalar) = "cosmos.AddressString" ];
  // age is the age of the student
  uint64 age = 3;
}
```

# Messages

This module has two messages:

- CreateStudent
- DeleteStudent

## CreateStudent

This message creates a new student.

Here's the proto definition of the message:

```proto
// Msg defines the Msg service.
service Msg {
  option (cosmos.msg.v1.service) = true;

  // CreateStudent create a new student on the module
  rpc CreateStudent(MsgCreateStudent) returns (MsgCreateStudentResponse);
}

// MsgCreateStudent is used when creating a new student
message MsgCreateStudent {
  option (cosmos.msg.v1.signer) = "id";

  // name is the name of the student
  string name = 1;
  // id is a unique address for each student
  string id = 2 [ (cosmos_proto.scalar) = "cosmos.AddressString" ];
  // age is the age of the student
  uint64 age = 3;
}

// MsgCreateStudentResponse is the response for the create student
message MsgCreateStudentResponse {}
```

This message will fail if:

- The student with the ID already exists

Upon successful processing:

- A new student is created

## DeleteStudent

This message delete student.

Here's the proto definition of the message:

```proto
// Msg defines the Msg service.
service Msg {
  option (cosmos.msg.v1.service) = true;

  // DeleteStudent remove a student from the module
  rpc DeleteStudent(MsgDeleteStudent) returns (MsgDeleteStudentResponse);
}

// MsgDeleteStudent is used when deleting a student
message MsgDeleteStudent {
  option (cosmos.msg.v1.signer) = "id";

  // id for the student who will be deleted
  string id = 1 [ (cosmos_proto.scalar) = "cosmos.AddressString" ];
}

// MsgDeleteStudentResponse is the response for the delete student
message MsgDeleteStudentResponse {}
```

This message will fail if:

- No student with the ID is found

Upon successful processing:

- The student with that ID is deleted

# Query

This module has a single query:

- GetStudent

## GetStudent

Returns a single student by it's ID.

### CLI

Usage:

```bash
rpsd query rps student [id]
```

Response:

```text
student:
  age: "20"
  id: rps1u07d74l5t35384af5yxulx28fp9mamrm3rcx5r
  name: alice
```

### Rest

At:

- `/challenge/rps/student/{id}`

Response:

```text
{
  "student": {
    "name": "alice",
    "id": "rps1qmuwmmz3twdrthjkd93w6ycu245na9caqygt02",
    "age": "20"
  }
}
```

# Begin-block

This module doesn't participate on the begin-block

# End-block

This module doesn't participate on the end-block

# Events

The claim module emits the following events:

## EventNewStudent

This is emitted on a new student creation.

Here's its proto definition:

```proto
// EventNewStudent is the event that is emitted on new student creation
message EventNewStudent {
  // name is the name of the student
  string name = 1;
  // id is the id of the student
  string id = 2 [ (cosmos_proto.scalar) = "cosmos.AddressString" ];
  // age is the age of the student
  uint64 age = 3;
}
```

## EventDeleteStudent

This is emitted on a student removal.

Here's its proto definition:

```proto
// EventDeleteStudent is the event tht is emitted on student removal
message EventDeleteStudent {
  // id is the id of the student
  string id = 1 [ (cosmos_proto.scalar) = "cosmos.AddressString" ];
}
```
