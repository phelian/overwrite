# Go overwrite

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/phelian/overwrite) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Currently, overwrite requires Go version 1.13 or greater.

This repository contains a library that enable caller to setup two structs of the same type and copy only tagged values between then.

Possible use cases are:

- Dynamically set up a struct based on input files
- Overloading secret configuration in environments where secrets are stored outside repo, escpecially useful for docker and kubernetes environments.

## Usage

```go
import "github.com/phelian/overwrite"

type Config struct {
    Username string
    Password string `overwrite:"true"`
    Driver string `overwrite:"true,omitempty"`
}

func main() {
    cfg := &Config{
        Username: "bernie",
        Driver: "postgres",
    }

    secretCfg := Config {
        Password: "feelthe",
    }

    if err := overwrite.Do(cfg, secretCfg); err != nil {
        panic(err)
    }
}
```

The result will be that `cfg` be a pointer to a Config struct with contents

```
{
    Username: "bernie",
    Password: "feelthe",
    Driver: "postgres",
}
```

### Do

Do copies tagged fields of arguments <src> into <dst>

dst needs to be a pointer to same type of struct that src is
src needs to be passed as value and not a pointer

Do traverses the value src recursively. If an encountered field is tagged to overwrite it tries to copy the value of that field into the dst counterpart field.

#### Errors

| Error             | Description                                                                                    |
| ----------------- | ---------------------------------------------------------------------------------------------- |
| ErrSrcNil         | Returned when input src is nil                                                                 |
| ErrDstNil         | Returned when input dst is nil                                                                 |
| ErrDstNotPtr      | Returned when input dst is not pointer to struct                                               |
| ErrSrcNotStruct   | Returned when input src is not struct type                                                     |
| ErrNotSameType    | Returned when the input dst and src are not pointers or and static copy of same type of struct |
| ErrCannotSetField | Returned when a tag have been set on a field that cannot be set, for instance not exported     |
| ErrTagValueWrong  | Wrong input in tag, see Tags section for possible values allowed                               |

### Tags

```go
`overwrite:"<true/false>[,omitempty]"`
```

Possible arguments for tag are boolean "true" or "false" with possibility to add ",omitempty" in the end.
| Tag Value | Description |
| --- | --- |
| `true` | Overwrite value from src struct into dst struct |
| `false` | In effect redundant and same as not setting tag |
| `,omitempty` | Do not overwrite if source value is empty |

## Compatable types

There types can be overwritten, others will be silently ignored

- int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64
- string
- boolean
- float32, float64
  Also
- arrays, slices, maps and structs containing compatable types
