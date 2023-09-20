<!-- Start SDK Example Usage -->


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
    res, err := s.GetArtists(ctx, operations.GetArtistsRequest{
        Limit: TestAPI.Int64(548814),
        Offset: TestAPI.Int64(592845),
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