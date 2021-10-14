# Flutterの試験用のGraphQLサーバ

ログイン機能とToDo機能を備えた簡単なGraphQLサーバです。


## ビルド方法

```bash
$ make
2021/10/15 03:22:06 connect to http://localhost:8080/ for GraphQL playground
```


## 使い方

1. ユーザ作成

```:クエリ
mutation createUser {
  createUser(input: {name: "taro", password: "secret"}) {
    userName
    jwt
  }
}
```
↓
```:レスポンス
{
  "data": {
    "createUser": {
      "userName": "taro",
      "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ1c2VyIiwiZXhwIjoxNjY1NzcyMDc0LCJqdGkiOiJ0YXJvIiwiaWF0IjoxNjM0MjM2MDc0LCJpc3MiOiJmbHV0dGVyLWV4YW0iLCJuYmYiOjE2MzQyMzYwNzR9.iRonPi7nzf9Xk5uFKIQOrRUPwfhhBBzaE5RqvGy8iLQ"
    }
  }
}
```

2. ログイン

```:クエリ
mutation createUser {
  authenticate(input: {name: "taro", password: "secret"}) {
    userName
    jwt
  }
}
```
↓
```:レスポンス
  "data": {
    "authenticate": {
      "userName": "taro",
      "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ1c2VyIiwiZXhwIjoxNjY1NzcyMjYwLCJqdGkiOiJ0YXJvIiwiaWF0IjoxNjM0MjM2MjYwLCJpc3MiOiJmbHV0dGVyLWV4YW0iLCJuYmYiOjE2MzQyMzYyNjB9.aRHpK1R1QZJCYwmfm1WwMkhXb51TpEIx9RzU9AADAfQ"
    }
  }
}
```

3. ToDoの追加

```:クエリ
mutation {
  createTodo(input: {userName: "taro", text: "B"}) {
    todoID
    text
  }
}
```

```:ヘッダー
{
  "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ1c2VyIiwiZXhwIjoxNjY1NzY5ODk1LCJqdGkiOiJ0YXJvIiwiaWF0IjoxNjM0MjMzODk1LCJpc3MiOiJmbHV0dGVyLWV4YW0iLCJuYmYiOjE2MzQyMzM4OTV9.TnOn5Cktrq6QgfuwuD5jjV65ZIAMsyg8Q1GI-1bXogo"
}
```
↓
```:レスポンス
{
  "data": {
    "createTodo": {
      "todoID": "6ab37bf6-5c54-4e43-9c6a-ef7ac6344a91",
      "text": "B"
    }
  }
}
```

3. ToDoの確認

```:クエリ
query {
  todos {
    todoID
    text
    done
  }
}
```

```:ヘッダー
{
  "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ1c2VyIiwiZXhwIjoxNjY1NzY5ODk1LCJqdGkiOiJ0YXJvIiwiaWF0IjoxNjM0MjMzODk1LCJpc3MiOiJmbHV0dGVyLWV4YW0iLCJuYmYiOjE2MzQyMzM4OTV9.TnOn5Cktrq6QgfuwuD5jjV65ZIAMsyg8Q1GI-1bXogo"
}
```
↓
```:レスポンス
{
  "data": {
    "todos": [
      {
        "todoID": "6ab37bf6-5c54-4e43-9c6a-ef7ac6344a91",
        "text": "B",
        "done": false
      }
    ]
  }
}
```

3. ToDoの完了

```:クエリ
mutation doneTodo {
  doneTodo(input: {todoID: "6c94a51b-1eeb-4e88-a40f-0c0e8ee977a3"}) {
    todoID
  }
}
```

```:ヘッダー
{
  "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ1c2VyIiwiZXhwIjoxNjY1NzY5ODk1LCJqdGkiOiJ0YXJvIiwiaWF0IjoxNjM0MjMzODk1LCJpc3MiOiJmbHV0dGVyLWV4YW0iLCJuYmYiOjE2MzQyMzM4OTV9.TnOn5Cktrq6QgfuwuD5jjV65ZIAMsyg8Q1GI-1bXogo"
}
```
↓
```:レスポンス
{
  "data": {
    "doneTodo": {
      "todoID": "6ab37bf6-5c54-4e43-9c6a-ef7ac6344a91"
    }
  }
}
```

