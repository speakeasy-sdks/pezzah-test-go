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