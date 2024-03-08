<!-- Start SDK Example Usage [usage] -->
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
<!-- End SDK Example Usage [usage] -->