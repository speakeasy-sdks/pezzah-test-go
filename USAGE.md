<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	testapi "TestAPI"
	"TestAPI/pkg/models/shared"
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
        Limit: testapi.Int64(919877),
        Offset: testapi.Int64(481153),
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