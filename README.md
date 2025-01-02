# h3-go

This is a reimplementation of the [H3](https://github.com/uber/h3) spatial indexing library in pure Go (i.e. without using CGO to wrap the existing C++ library).

## Design goals

1. **Pure Go**: The library should be usable in any Go environment without requiring any additional dependencies. Notably, CGO should not be required.
2. **Performance**: The library should be fast enough to be usable in production applications.
3. **Idiomatic Go API**: The library should have an API that feels natural to Go developers.

## Status

The library is in the early stages of development. We are implementing the core functionality of the H3 library, but not all features are available yet. The API is also subject to change.

Some core pieces of functionality are implemented:

- [x] Basic H3 index/cell Go types
- [x] Conversion between lat/lon and H3 indexes
- [x] Grid Disk algorithm

Other important features are not yet implemented:
- [ ] Clean up public API
- [ ] Grid Ring algorithm
- [ ] Performance optimizations, including microbenchmarking

## Usage

### Convert a lat/lon to an H3 index:

```go
package main

import (
	"github.com/ziprecruiter/h3-go/pkg/h3"
)

func main() {
	lat := 37.775938728915946
	lon := -122.41795063018799
	resolution := 9
	
	ll := h3.NewLatLng(lat, lon)
	h3Index, err := h3.NewCellFromLatLng(ll, resolution)
	if err != nil {
		panic(err)
	}

	println(h3Index)
}
```