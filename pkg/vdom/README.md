# VDOM virtual dom interface and websockets streaming client

VDOM is a browser DOM abstraction layer derived from a simplified https://github.com/albrow/vdom
and supporting

* efficent DOM diff and patch
* components (a la React)
* registration of DOM event handlers
* Streaming of DOM diff and events via websocket

In streaming mode - DOM interaction can be implemented in pure (native compiled)
Go and steamed via a built in web server to a listening websocket client
