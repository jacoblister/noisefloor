// +build js

package assets

// Assets returns a nil filesystem for the javascript target
var Assets = func() interface{} {
	return nil
}
