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
	"github.com/sdn0303/gcputil/storage"
	"github.com/sdn0303/gcputil/errorreporting"

)

var (
	errReporting *errorreporting.ErrorReporting
	gcs          *storage.Storage
)

func init() {
	ctx := context.Background()
	errReporting = errorreporting.New(ctx, "{ProjectID}", "{ServiceName}")
	gcs = storage.New("{BucketName}", ctx)
}

func main() {
	
	data, err := ioutil.ReadFile("{yourFile}")
	errReporting.SendError(err)
	
	err := gcs.Put("{Prefix}", "{ContentType}", data)
	errReporting.SendError(err)
}
```