# envYaml (Yaml with environment value loader)

[![codecov](https://codecov.io/github/yuseferi/envyaml/graph/badge.svg?token=0DUS258IUD)](https://codecov.io/github/yuseferi/envyaml)
[![Check & Build](https://github.com/yuseferi/envyaml/actions/workflows/ci.yml/badge.svg)](https://github.com/yuseferi/envyaml/actions/workflows/ci.yml)
[![License: AGPL v3](https://img.shields.io/badge/License-AGPL_v3-blue.svg)](https://www.gnu.org/licenses/agpl-3.0)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/yuseferi/envyaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/yuseferi/envyaml)](https://goreportcard.com/report/github.com/yuseferi/envyaml)

<p align="center">
<img src="https://github.com/user-attachments/assets/b6b5bbc6-f9d7-4d2f-b5c8-e86ce0e0fd9b" width="250" />
</p>

Storing application configuration in YAML files offers a clean and straightforward solution, but it's crucial to avoid exposing sensitive data like passwords and API keys. Environment variables provide a secure way to store secrets, preventing them from leaking into your codebase.

The **envyaml** package bridges this gap, allowing you to seamlessly integrate environment variables into your YAML configuration. No more plain-text secrets or complex workarounds! Simply reference environment variables within your YAML file using placeholders, and envyaml will handle the rest, securely substituting the values during parsing.

With **envyaml**, you can:

- Keep your configuration clean and organized in YAML.
- Protect sensitive data by storing it in environment variables.
- Enjoy a simple and intuitive integration process.
- No more compromising between convenience and security. **envyaml** empowers you to manage your application configuration effectively while keeping your secrets safe.

## Installation

To install envyaml, use `go get`:

```
go get github.com/yuseferi/envyaml@latest
```

## Usage

### Example 1: Error on required env variable

```go
type TestConfig struct {
    Host     string `yaml:"host" env:"TEST_HOST"`
    Port     int    `yaml:"port" env:"TEST_PORT"`
    Password string `yaml:"password" env:"TEST_PASSWORD,required"`
}

// Load the configuration
var cfg TestConfig
// assume your configs are in `config.yml` file and this is its content:
//host: localhost
//port: 3606
//password: ${TEST_PASSWORD}
err := envyaml.LoadConfig("config.yml", &cfg)
if err != nil {
    log.Fatalln(err)
}
fmt.Println(cfg)
```

Error `failed to parse environment variables: env: required environment variable "TEST_PASSWORD" is not set` is expected because this variable has not been set. 

### Example 2: When env is defined

```go
type TestConfig struct {
    Host     string `yaml:"host" env:"TEST_HOST"`
    Port     int    `yaml:"port" env:"TEST_PORT"`
    Password string `yaml:"password" env:"TEST_PASSWORD,required"`
}

// Load the configuration
var cfg TestConfig
// assume your configs are in `config.yml` file and this is its content:
//host: localhost
//port: 3606
//password: ${TEST_PASSWORD}
_ = os.Setenv("TEST_PASSWORD", "envyaml_pass")
err := envyaml.LoadConfig("config.yml", &cfg)
if err != nil {
log.Fatalln(err)
log.Fatalln(err)

    log.Fatalln(err)

}
fmt.Println(cfg)
```

Expected output:
```
{localhost 3606 envyaml_pass}
```

## Development

This project uses [Task](https://taskfile.dev) for managing development tasks. Make sure you have Task installed on your system.

### Available Tasks

- `task build`: Build the project
- `task test`: Run tests
- `task test-coverage`: Run tests with coverage and generate a coverage report
- `task clean`: Clean up generated files
- `task all`: Run all tasks (build, test, and coverage)

To run a task, use the `task` command followed by the task name. For example:

```
task build
```

### Running Tests

To run tests:

```
task test
```

To run tests with coverage:

```
task test-coverage
```

This will generate a coverage report in HTML format (`coverage.html`).

## Contributing

We strongly believe in open-source ❤️😊. Please feel free to contribute by raising issues and submitting pull requests to make envYaml even better!

## License

Released under the [GNU GENERAL PUBLIC LICENSE](LICENSE).


