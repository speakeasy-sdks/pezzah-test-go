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
	"context"
	"log"
	"TestAPI"
	"TestAPI/pkg/models/shared"
	"TestAPI/pkg/models/operations"
)

func main() {
    s := TestAPI.New(
        TestAPI.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.TestAPI.GetArtists(ctx, operations.GetArtistsRequest{
        Limit: TestAPI.Int64(715190),
        Offset: TestAPI.Int64(844266),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetArtists200ApplicationJSONObjects != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetArtistsRequest](../../models/operations/getartistsrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetArtistsResponse](../../models/operations/getartistsresponse.md), error**


## GetArtistsUsername

Obtain information about an artist from his or her unique username

### Example Usage

```go
package main

import(
	"context"
	"log"
	"TestAPI"
	"TestAPI/pkg/models/shared"
	"TestAPI/pkg/models/operations"
)

func main() {
    s := TestAPI.New(
        TestAPI.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.TestAPI.GetArtistsUsername(ctx, operations.GetArtistsUsernameRequest{
        Username: "Leda_Stiedemann",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetArtistsUsername200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetArtistsUsernameRequest](../../models/operations/getartistsusernamerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.GetArtistsUsernameResponse](../../models/operations/getartistsusernameresponse.md), error**


## PostArtists

Lets a user post a new artist

### Example Usage

```go
package main

import(
	"context"
	"log"
	"TestAPI"
	"TestAPI/pkg/models/shared"
	"TestAPI/pkg/models/operations"
)

func main() {
    s := TestAPI.New(
        TestAPI.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.TestAPI.PostArtists(ctx, operations.PostArtistsRequestBody{
        AlbumsRecorded: TestAPI.Int64(423655),
        ArtistGenre: TestAPI.String("error"),
        ArtistName: TestAPI.String("deserunt"),
        Username: "Gerda89",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PostArtistsRequestBody](../../models/operations/postartistsrequestbody.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.PostArtistsResponse](../../models/operations/postartistsresponse.md), error**

