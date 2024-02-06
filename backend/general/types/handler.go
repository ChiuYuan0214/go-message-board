package types

import "net/http"

type Hander func(req *http.Request) (res interface{}, statusCode int)
