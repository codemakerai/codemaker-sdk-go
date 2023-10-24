# codemaker-sdk-go

Official SDK for CodeMaker AI APIs.

# Getting Started

In order to start using the SDK.

1. Sign up for the Early Access Program at https://codemaker.ai.
2. Receive the Early Access invitation email.
3. Add CodeMaker SDK to your project.

```bash
$ go get github.com/codemakerai/codemaker-sdk-go
```

4. Integrate the SDK with your code.

```go
c := client.NewClient(&client.Config{
    ApiKey: apiKey,
})
```

# License

MIT License