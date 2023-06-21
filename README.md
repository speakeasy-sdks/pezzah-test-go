# TestAPI

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/pezzah-test-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"TestAPI"
	"TestAPI/pkg/models/operations"
)

func main() {
    s := testapi.New(
        testapi.WithSecurity(shared.Security{
            Password: "",
            Username: "",
        }),
    )

    ctx := context.Background()
    res, err := s.GetArtists(ctx, operations.GetArtistsRequest{
        Limit: testapi.Int64(548814),
        Offset: testapi.Int64(592845),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetArtists200ApplicationJSONObjects != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations

### [TestAPI SDK](docs/sdks/testapi/README.md)

* [GetArtists](docs/sdks/testapi/README.md#getartists) - Returns a list of artists
* [GetArtistsUsername](docs/sdks/testapi/README.md#getartistsusername) - Obtain information about an artist from his or her unique username
* [PostArtists](docs/sdks/testapi/README.md#postartists) - Lets a user post a new artist
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
