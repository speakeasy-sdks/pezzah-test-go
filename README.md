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

import (
	testapi "TestAPI"
	"TestAPI/pkg/models/operations"
	"TestAPI/pkg/models/shared"
	"context"
	"log"
)

func main() {
	s := testapi.New(
		testapi.WithSecurity(shared.Security{
			Password: "",
			Username: "",
		}),
	)

	ctx := context.Background()
	res, err := s.TestAPI.GetArtists(ctx, operations.GetArtistsRequest{})
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



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
<!-- End Pagination -->



<!-- Start Go Types -->

<!-- End Go Types -->



<!-- Start Error Handling -->
# Error Handling

Handling errors in your SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.


<!-- End Error Handling -->



<!-- Start Custom HTTP Client -->
# Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
