# golang-graphql-simple

If you want to run this project, open terminal and navigate to the project folder.
Then run `docker-compose up` 

You can send your requests on `localhost:8080`

Create request example:

```
mutation {
    create(title: "something", content: "something") {
        id,
        title,
        content
    }
}
```

Get request example:

```
query {
    node(id: 1) {
        id,
        title,
        content
    }
}
```

[GraphQl docs](https://graphql.org/learn/)
