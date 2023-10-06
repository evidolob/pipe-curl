# pipe-curl
Window CLU util to make HTTP requests through Windows named pipe

# Usage

To use this, just call `pipe-curl.exe -p <path_to_pipe> http://localhost/<your_api_path>`

# Build

To build run:
```shell
  go build -mod vendor
```
> Note: This is Windows only tool, build possible only on windows