# negroni-requestid
Negroni middleware to add a randomly generated request ID to the request context
and optionally add it as an X-Request-ID header.

## Usage

__Default Options__

By default, `negroni-requestid` will generate a UUID-based ID, store it in the
http.Request's Context and write the ID to the X-Request-ID header on the
response.
```go
n := negroni.New()
n.Use(requestid.NewMiddleware())
```

__Custom ID Generator__

If you'd like to use a different way to generate request ID's other than a
standard UUID, it is possible to provide a new function to do it:
```go
n := negroni.New()

reqmw := requestid.NewMiddleware()
reqmw.GenerateID = func() string {
    return "7"
}
```

__Disable X- Header Output__

If you'd like to disable the X-Request-ID header being added to responses or
you'd like to change the header being written, you are able to set that on the
middleware.  If you'd like to disable the header completely, just supply an empty
string:
```go
n := negroni.New()

reqmw := requestid.NewMiddleware()
reqmw.XHeader = "X-My-ID"
// or to disable it:
reqmw.XHeader = ""
```

__Retrieve ID From Context__

To retrieve the request ID from an http.Request's Context, you can use the
`FromContext` method in the package:
```go
func handleReq(w http.ResponseWriter, req *http.Request) {
    reqID, err := requestid.FromContext(req.Context())
    if err != nil {
        fmt.Printf("Error getting request ID: %s\n", err)
    } else {
        fmt.Printf("Got request ID: %s\n", reqID)
    }
}

```