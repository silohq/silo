# Thoughts while building

Here is the final doc definition that will represent client data

```json
{
  "tree": [
    {
      "parent": "root",
      "type": "document",
      "label": "user",
      "nodes": [
        {
          "parent": "user",
          "type": "flat-node",
          "kind": "string",
          "unique": false,
          "label": "name"
        },
        {
          "label": "contact",
          "parent": "user",
          "type": "nested-node",
          "kind": "object",
          "nodes": [
            {
              "parent": "user.contact",
              "type": "flat-node",
              "kind": "string",
              "unique": false,
              "label": "primary"
            },
            {
              "parent": "user.contact",
              "type": "flat-node",
              "kind": "string",
              "unique": false,
              "label": "secondary"
            }
          ]
        }
      ]
    }
  ]
}
```

will represent this

```json
{
  "user": {
    "name": "",
    "contact": {
      "primary": "",
      "secondary": ""
    }
  }
}
```

## create command

```
new /users/ {payload}
```

## update command

{:variable} -> indicates variable being passed in and where its going to be
placed

```
mod /users/{:id}/{:email} {payload} match {id}
```

```
mod /users/{:id}/{:email}/{:name} {payload}
```

#### payload below:

```json
{
  "id": "idhere",
  "email": "emailhere",
  "name": "has to match field name in command and doc structure as well"
}
```
