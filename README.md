# silo

experimental document based access for boltdb

## USAGE

```
go get github.com/silohq/silo
```

### Importing silo

Silo uses boltdb as the underlying key value store. It also expects a [variable].json
file that dictates the structure of the root document. Sample can be found in notes/.
DBPath is the file used by bolt engine.

```go
pkg, err := silo.New(&silo.Config{
    DocPath: "definition.json",
    DBPath:  "test.db",
})
if err != nil {
    log.Printf("could not create silo %s", err)
}
```

## Features

### create

Features are currently very limited.

- create is currently only implemented at the root level.
- creating json object with missing fields leads to a lot of issues

basic creation is aimed to work as below:

```go
pkg.Create("user", map[string]interface{}{
    "user": map[string]interface{}{
        "name": "last",
        "contact": map[string]interface{}{
            "primary":   "primary",
            "secondary": "secondary",
        },
    },
})
```

### find

find is implemented at any level but is currently very slow for
larger stores

[m] defines the fields you would prefer to access . such that if you
passed in a json object as below. only the buckets containing the below values
would be queried

```json
{
  "user": {
    "name": "",
    "contact": {
      "primary": ""
    }
  }
}
```

issues

- currently only working with []byte type when fetching values

```go
m := map[string]interface{}{
    "user": map[string]interface{}{
        "name": "",
        "contact": map[string]interface{}{
            "primary":   "",
            "secondary": "",
        },
    },
}

pkg.Find("user.name", "last", m)

```

Final query language will interact with db as below:

```
create /user {...payload}

find /user/name "last"
```
