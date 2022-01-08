# silo

experimental document based access for boltdb

### Access path is for all patterns [CRUD]

```json
{
  "user": {
    "chats": {
      "refs": ["001", "002", "003"]
    }
  },

  "chats": {
    "001": {
      "title": "name",
      "messages": {
        "m1": {
          "id": "",
          "time": "now",
          "content": "contenthere"
        }
      }
    }
  }
}
```

fetch exactly one chat according to id.
--> one /users/chats/${id}

Access all chats
--> fetch /users/chats/${:limit, 5}

Delete one message
--> del /users/chats/:ref{id}/messages/${id}

:ref id would mean the path has to access a different bucket from the
default user bucket. 

${id} is a required field for all defined objects. if absent then the entire document
is accessed.