# Options

Options control what `protoc-gen-gorm` does. You set them in your `.proto` files, as regular  
[Protocol Buffer Options](https://developers.google.com/protocol-buffers/docs/proto3#options).

By default `protoc-gen-gorm` does nothing, you'll have to flag some of your messages to be models first, e.g. set [`model`](#model_1) to `true`.

## File Options

File options apply to all message types within the `.proto` file.

### model

Sets `model` for **all** messages in the file. [See `model` below](#model_1).

**Default:** `false`

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

option (gorm.file).model = true;
```

---

### validate

Sets `validate` for **all** messages in the file. [See `validate` below](#validate_1).

Implies `model = true` when set to `true`.

**Default:** `false`

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

option (gorm.file).validate = true;
```

---

### crud

Sets `crud` for **all** messages in the file. [See `crud` below](#crud_1).

Implies `model = true` when set to `true`.

**Default:** `false`

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

option (gorm.file).crud = true;
```

## Message Options

Message options control generation of model and supporting code for your message type.

### model

Marks a message as a model and have `protoc-gen-gorm` generate a Go struct for use with GORM v2.

The struct type name is the message name with "Model" appended.

**Default:** `false`

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;
}
```

Generates:

```go
package mypackage

type MyMessageModel struct {
	// ...
}
```

---

### validate

**TODO**

Implies `model = true` when set to `true`.

**Default:** `false`

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).validate = true;
}
```

---

### crud

TODO

Implies `model = true` when set to `true`.

**Default:** `false`

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).crud = true;
}
```

---

### table

Set the table name for models of this type.

**Default:** Unset, uses the [GORM default](https://gorm.io/docs/conventions.html#Pluralized-Table-Name).

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message) = {
    model: true,
    table: "mytable"
  };
}
```

The generated struct now implements [GORM's Tabler interface](https://pkg.go.dev/gorm.io/gorm/schema#Tabler):

```go
package mypackage

type MyMessageModel struct {
	// ...
}

func (m *MyMessageModel) TableName() string {
	return "mytable"
}
```

## Field Options

Field options refine how your generated model works with GORM through struct field tags and supporting code.

### column

Sets the [database column name](https://gorm.io/docs/conventions.html#Column-Name).

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string my_field = 1 [
    (gorm.field).column = "my_column"
  ];
}
```

Equivalent GORM struct field tag:

```go
package mypackage

type MyMessageModel struct {
	MyField string `gorm:"column:my_column"`
}
```

---

### not_null

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string my_field = 1 [
    (gorm.field).not_null = true
  ];
}
```

Equivalent GORM struct field tag:

```go
```

---

### default

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string my_field = 1 [
    (gorm.field).default = "a default value"
  ];
}
```

Equivalent GORM struct field tag:

```go
```

---

### unique

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string my_field = 1 [
    (gorm.field).unique = true
  ];
}
```

Equivalent GORM struct field tag:

```go
```

---

### primary_key

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string uuid = 1 [
    (gorm.field).primary_key = true
  ];
}
```

Equivalent GORM struct field tag:

```go
```

---

### index

**TODO**

#### default

**TODO**

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string my_field = 1 [
    (gorm.field).index = {default: true}
  ];
}
```

Equivalent GORM struct field tag:

```go
```

#### name

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string my_field = 1 [
    (gorm.field).index = {name: "my_index_name"}
  ];
}
```

Equivalent GORM struct field tag:

```go
```

---

### unique_index

TODO

#### default

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string my_field = 1 [
    (gorm.field).unique_index = {default: true}
  ];
}
```

Equivalent GORM struct field tag:

```go
```

#### name

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string my_field = 1 [
    (gorm.field).unique_index = {name: "my_index_name"}
  ];
}
```

Equivalent GORM struct field tag:

```go
```

---

### auto_create_time

Instructs GORM to [track creation time](https://gorm.io/docs/models.html#Creating-x2F-Updating-Time-x2F-Unix-Milli-x2F-Nano-Seconds-Tracking) in the flagged field.

**Example:**

```protobuf
syntax = "proto3";
import "google/protobuf/timestamp.proto";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  google.protobuf.Timestamp my_time = 1 [
    (gorm.field).auto_create_time = true
  ];
}
```

Equivalent GORM struct field tag:

```go
package mypackage

import "time"

type MyMessageModel struct {
	MyTime time.Time `gorm:"autoCreateTime"`
}
```

---

### auto_update_time

Instructs GORM to [track update time](https://gorm.io/docs/models.html#Creating-x2F-Updating-Time-x2F-Unix-Milli-x2F-Nano-Seconds-Tracking) in the flagged field.

**Example:**

```protobuf
syntax = "proto3";
import "google/protobuf/timestamp.proto";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  google.protobuf.Timestamp my_time = 1 [
    (gorm.field).auto_update_time = true
  ];
}
```

Equivalent GORM struct field tag:

```go
package mypackage

import "time"

type MyMessageModel struct {
	MyTime time.Time `gorm:"autoUpdateTime"`
}
```

---

### permissions

**TODO**

#### ignore

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string my_ignored_field = 1 [
    (gorm.field).ignore = true
  ];
}
```

Equivalent GORM struct field tag:

```go
```

#### deny

**TODO**

##### create

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string my_field = 1 [
    (gorm.field).deny = {create: true}
  ];
}
```

Equivalent GORM struct field tag:

```go
```

##### update

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string my_field = 1 [
    (gorm.field).deny = {update: true}
  ];
}
```

Equivalent GORM struct field tag:

```go
```

##### read

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  string my_field = 1 [
    (gorm.field).deny = {read: true}
  ];
}
```

Equivalent GORM struct field tag:

```go
```

---

### json

TODO

**Example:**

```protobuf
syntax = "proto3";
import "gorm/options.proto";
package mypackage;

message MyMessage {
  option (gorm.message).model = true;

  map<string, string> my_map = 1 [
    (gorm.field).json = true
  ];
}
```
