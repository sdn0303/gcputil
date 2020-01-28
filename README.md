# gcputil
This is utils package for google cloud platform.

### Supported
#### Google Cloud Storage
| actions |    |    |
|:-------:|:--:|:--:|
| Put     | ●  |    |
| Read    | ●  |    |

#### Stackdriver Error Reporting
| actions |    |    |
|:-------:|:--:|:--:|
| Send    | ●  |    |

Other services will be added sequentially.

### Examples
```go
package main

import (
	"context"
    "io/ioutil"

	"github.com/sdn0303/gcputil/storage"
	"github.com/sdn0303/gcputil/errorreporting"
)

func main() {
	ctx := context.Background()
	errReporting := errorreporting.New(ctx, "{ProjectID}", "{ServiceName}")

	data, err := ioutil.ReadFile("{yourFile}")
	if err != nil {
		errReporting.SendError(err)
	}

	gcs := storage.New(ctx, "{BucketName}")
	if err := gcs.Put(ctx, "{Prefix}", "{ContentType}", data); err != nil {
		errReporting.SendError(err)
	}
	
}
```