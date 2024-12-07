package lib

import "io"

type Solver func(io.Reader) (any, any)
