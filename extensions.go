package gin

import (
	"strconv"
)

func (c *Context) ParamPtr(key string) *string {
	if val, ok := c.Params.Get(key); ok {
		return &val
	}
	return nil
}

func (c *Context) Int64Param(key string) *int64 {
	if val := c.ParamPtr(key); val != nil {
		if v, err := strconv.ParseInt(*val, 10, 64); err == nil {
			return &v
		}
	}
	return nil
}

func (c *Context) Id() *int64 {
	return c.Int64Param("id")
}
