# silo

experimental document based access for boltdb

create is currently only implemented at the root level.

```go
pkg.Create("/user", map[string]interface{}{
    "user": map[string]interface{}{
        "name": "last",
        "contact": map[string]interface{}{
            "primary":   "primary",
            "secondary": "secondary",
        },
    },
})
```

find is implemented at any level

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