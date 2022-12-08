package gin

import "time"

func (c *Context) GetId() *int64 {
	if id := c.GetInt64Ptr("id"); id != nil && *id <= 0 {
		return nil
	} else {
		return id
	}
}

// GetString returns the value associated with the key as a string.
func (c *Context) GetStringPtr(key string) *string {
	if val, ok := c.Get(key); ok && val != nil {
		s, _ := val.(string)
		return &s
	}
	return nil
}

// GetBool returns the value associated with the key as a boolean.
func (c *Context) GetBoolPtr(key string) *bool {
	if val, ok := c.Get(key); ok && val != nil {
		b, _ := val.(bool)
		return &b
	}
	return nil
}

// GetInt returns the value associated with the key as an integer.
func (c *Context) GetIntPtr(key string) *int {
	if val, ok := c.Get(key); ok && val != nil {
		i, _ := val.(int)
		return &i
	}
	return nil
}

// GetInt64 returns the value associated with the key as an integer.
func (c *Context) GetInt64Ptr(key string) *int64 {
	if val, ok := c.Get(key); ok && val != nil {
		n, _ := val.(int64)
		return &n
	}
	return nil
}

// GetUint returns the value associated with the key as an unsigned integer.
func (c *Context) GetUintPtr(key string) *uint {
	if val, ok := c.Get(key); ok && val != nil {
		ui, _ := val.(uint)
		return &ui
	}
	return nil
}

// GetUint64 returns the value associated with the key as an unsigned integer.
func (c *Context) GetUint64Ptr(key string) *uint64 {
	if val, ok := c.Get(key); ok && val != nil {
		ui64, _ := val.(uint64)
		return &ui64
	}
	return nil
}

// GetFloat64 returns the value associated with the key as a float64.
func (c *Context) GetFloat64Ptr(key string) *float64 {
	if val, ok := c.Get(key); ok && val != nil {
		f64, _ := val.(float64)
		return &f64
	}
	return nil
}

// GetTime returns the value associated with the key as time.
func (c *Context) GetTimePtr(key string) *time.Time {
	if val, ok := c.Get(key); ok && val != nil {
		t, _ := val.(time.Time)
		return &t
	}
	return nil
}

// GetDuration returns the value associated with the key as a duration.
func (c *Context) GetDurationPtr(key string) *time.Duration {
	if val, ok := c.Get(key); ok && val != nil {
		d, _ := val.(time.Duration)
		return &d
	}
	return nil
}
