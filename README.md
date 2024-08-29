<div align="center">
	<h1>Toolkit</h1>
    <p>Go commonly used development kits</p>
	<img src="assets/logo.png" alt="logo" width="400px">
</div>

## Modules

All modules can be found in the [pkg](./pkg/) directory.

### 1. Command

The [**command**](./pkg/command/) module is used to execute `useage` and `help` commands, making it very useful. It is based on the `github.com/spf13/cobra` package.

### 2. Config

The [**config**](./pkg/config/) module is used to load configuration from files. It has been tested with `yaml`, `json`, and `toml` files.

### 3. Conver

The [**conver**](./pkg/conver/) module is used to convert data types, such as converting `string` to `int`, `string` to `bytes`, etc.

### 4. Httptool

The [**httptool**](./pkg/httptool/) module is used to handle standard response data structures.

### 5 HttpServer

The [**httpserver**](./pkg/httpserver/) module is a tiny http server that can serve Kubernetes Pod. It is written in Go and uses the standard library.
