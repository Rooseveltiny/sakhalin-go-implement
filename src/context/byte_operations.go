package context

import "sakhalin/color"

func (c *Context) AppendSingleByte(b byte)      { c.messagePool.AppendSingleByte(b) }
func (c *Context) AppendFloat64(f float64)      { c.messagePool.AppendFloat64(f) }
func (c *Context) AppendUint32(ui uint32)       { c.messagePool.AppendUint32(ui) }
func (c *Context) AppendBoolean(b bool)         { c.messagePool.AppendBoolean(b) }
func (c *Context) AppendBytes(bytes []byte)     { c.messagePool.AppendBytes(bytes) }
func (c *Context) AppendString(s string)        { c.messagePool.AppendString(s) }
func (c *Context) AppendColor(col color.IColor) { c.messagePool.AppendColor(col) }
func (c *Context) RetrieveSingleByte() byte     { return c.messagePool.RetrieveSingleByte() }
func (c *Context) RetrieveUint32() uint32       { return c.messagePool.RetrieveUint32() }
func (c *Context) RetrieveUint64() uint64       { return c.messagePool.RetrieveUint64() }
func (c *Context) RetrieveFloat64() float64     { return c.messagePool.RetrieveFloat64() }
func (c *Context) RetrieveBool() bool           { return c.messagePool.RetrieveBool() }
func (c *Context) RetrieveString() string       { return c.messagePool.RetrieveString() }
func (c *Context) Clear()                       { c.messagePool.Clear() }
