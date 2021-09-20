# Load Schema

When Dgraph service has been started, just execute the following command to create the schemas and generate a GraphQL API automatically in the route 'localhost:8080/graphql'

> curl -X POST localhost:8080/admin/schema --data-binary '@schema.graphql'

## Notes

The database can be modified in a few ways:

1. Using the Go client for Dgraph (recommended)
2. Using the GraphQL API (recommended)
3. Using the HTTP API
4. Using the Dgraph's frontend client (Ratel). Good for testing or manually making changes to the database.

The first two are recommended because they're easier and have great APIs. Although the third option is not necessarily a bad idea.

## Simple database use case example

1. Run Dgraph with docker

    > docker run -it -p 6080:6080 -p 8080:8080 -p 9080:9080 -p 8000:8000 -v ~/dgraph/:/dgraph dgraph/standalone:v20.03.0

2. Go to the Ratel GUI visiting the following URL

    > http://localhost:8000/?latest

3. Go to the "Console" tab (located at the left side)

4. Create the database by using [the file "schema.graphql"](../models/schema.graphql) using the following command

    > curl -X POST localhost:8080/admin/schema --data-binary '@schema.graphql'

5. Confirm that you have a success response like the following

    > {"data":{"code":"Success","message":"Done"}}

6. To populate schemas with data. Execute the following mutation from the "Console" tab of the Ratel client:

> **Note:** Please verify you select the "subtab" called "Mutate" when you're in the "Console" tab.

```
{
   "set":[
      {
         "Code.name":"hello-world.py",
         "Code.code":"print('hello world')",
         "Code.astTree": "A JSON serialized string AST representation"
      },
      {
         "Code.name":"hello-people.py",
         "Code.code":"print('hello people')",
         "Code.astTree": "A JSON serialized string AST representation 2.0"
      }
   ]
}
```

7. To query all data, execute the following query (from the "Query subtab")

> **Note:** You could change the name "data" for whatever name you find better if you want.

```
{
     data( func: has(Code.name)){
         uid
         Code.name
         Code.code
         Code.astTree
     }
}
```

8. The response from that query should look like this one:

```
{
  "data": {
    "data": [
      {
        "uid": "0x2",
        "Code.name": "hello-world.py",
        "Code.code": "print('hello world')",
        "Code.astTree": "A JSON serialized string AST representation"
      },
      {
        "uid": "0x4",
        "Code.name": "hello-people.py",
        "Code.code": "print('hello people')",
        "Code.astTree": "A JSON serialized string AST representation 2.0"
      }
    ]
  },
  "extensions": {
    "server_latency": {
      "parsing_ns": 68111,
      "processing_ns": 37725458,
      "encoding_ns": 24507,
      "assign_timestamp_ns": 645438,
      "total_ns": 38543462
    },
    "txn": {
      "start_ts": 32
    },
    "metrics": {
      "num_uids": {
        "AST.tree": 2,
        "Code.code": 2,
        "Code.name": 2,
        "ast": 2,
        "uid": 2
      }
    }
  }
}
```
