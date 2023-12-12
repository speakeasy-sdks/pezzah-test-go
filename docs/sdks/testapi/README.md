# TestAPI SDK


## Overview

Simple API: A simple API to illustrate OpenAPI concepts

### Available Operations

* [GetArtists](#getartists) - Returns a list of artists
* [GetArtistsUsername](#getartistsusername) - Obtain information about an artist from his or her unique username
* [PostArtists](#postartists) - Lets a user post a new artist

## GetArtists

Returns a list of artists

### Example Usage

```go
package main

import(
	"TestAPI/pkg/models/shared"
	testapi "TestAPI"
	"context"
	"TestAPI/pkg/models/operations"
	"log"
)

func main() {
    s := testapi.New(
        testapi.WithSecurity(shared.Security{
            Password: "<YOUR_PASSWORD_HERE>",
            Username: "<YOUR_USERNAME_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.GetArtists(ctx, operations.GetArtistsRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetArtistsRequest](../../pkg/models/operations/getartistsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.GetArtistsResponse](../../pkg/models/operations/getartistsresponse.md), error**
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.GetArtistsResponseBody | 400                              | application/json                 |
| sdkerrors.SDKError               | 400-600                          | */*                              |

## GetArtistsUsername

Obtain information about an artist from his or her unique username

### Example Usage

```go
package main

import(
	"TestAPI/pkg/models/shared"
	testapi "TestAPI"
	"context"
	"TestAPI/pkg/models/operations"
	"log"
)

func main() {
    s := testapi.New(
        testapi.WithSecurity(shared.Security{
            Password: "<YOUR_PASSWORD_HERE>",
            Username: "<YOUR_USERNAME_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.GetArtistsUsername(ctx, operations.GetArtistsUsernameRequest{
        Username: "Marilyne.Hackett49",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetArtistsUsernameRequest](../../pkg/models/operations/getartistsusernamerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.GetArtistsUsernameResponse](../../pkg/models/operations/getartistsusernameresponse.md), error**
| Error Object                             | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| sdkerrors.GetArtistsUsernameResponseBody | 400                                      | application/json                         |
| sdkerrors.SDKError                       | 400-600                                  | */*                                      |

## PostArtists

Lets a user post a new artist

### Example Usage

```go
package main

import(
	"TestAPI/pkg/models/shared"
	testapi "TestAPI"
	"context"
	"TestAPI/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := testapi.New(
        testapi.WithSecurity(shared.Security{
            Password: "<YOUR_PASSWORD_HERE>",
            Username: "<YOUR_USERNAME_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.PostArtists(ctx, operations.PostArtistsRequestBody{
        Username: "Cordie64",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PostArtistsRequestBody](../../pkg/models/operations/postartistsrequestbody.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.PostArtistsResponse](../../pkg/models/operations/postartistsresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.PostArtistsResponseBody | 400                               | application/json                  |
| sdkerrors.SDKError                | 400-600                           | */*                               |
