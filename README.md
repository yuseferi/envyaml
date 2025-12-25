<div align="center">

# 🔐 envYaml

### Seamlessly merge YAML configuration with environment variables

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![codecov](https://img.shields.io/codecov/c/github/yuseferi/envyaml?style=for-the-badge&logo=codecov&logoColor=white)](https://codecov.io/github/yuseferi/envyaml)
[![CI](https://img.shields.io/github/actions/workflow/status/yuseferi/envyaml/ci.yml?style=for-the-badge&logo=github&label=CI)](https://github.com/yuseferi/envyaml/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/yuseferi/envyaml?style=for-the-badge)](https://goreportcard.com/report/github.com/yuseferi/envyaml)
[![License](https://img.shields.io/badge/License-GPL_v3-blue?style=for-the-badge)](LICENSE)
[![Release](https://img.shields.io/github/v/release/yuseferi/envyaml?style=for-the-badge&logo=github)](https://github.com/yuseferi/envyaml/releases)

<p>
<img src="https://github.com/user-attachments/assets/b6b5bbc6-f9d7-4d2f-b5c8-e86ce0e0fd9b" width="280" alt="envYaml Logo" />
</p>

**Keep your configuration clean. Keep your secrets safe.**

[Installation](#-installation) •
[Quick Start](#-quick-start) •
[Features](#-features) •
[Examples](#-examples) •
[Contributing](#-contributing)

---

</div>

## 🎯 The Problem

You love YAML for configuration—it's clean, readable, and organized. But what about sensitive data like API keys, database passwords, and tokens? Hardcoding them is a security nightmare. 😱

## ✨ The Solution

**envYaml** bridges the gap between clean YAML configuration and secure environment variable management. Reference environment variables directly in your YAML files, and envYaml handles the rest!

```yaml
# config.yml - Clean and secure! 🔒
database:
  host: localhost
  port: 5432
  password: ${DB_PASSWORD}  # Loaded from environment

api:
  key: ${API_KEY}           # Never committed to git
  secret: ${API_SECRET}     # Always secure
```

## 🚀 Installation

```bash
go get github.com/yuseferi/envyaml@latest
```

**Requirements:** Go 1.25+

## ⚡ Quick Start

**1. Create your YAML configuration:**

```yaml
# config.yml
host: localhost
port: 3606
password: ${DB_PASSWORD}
```

**2. Define your config struct:**

```go
type Config struct {
    Host     string `yaml:"host" env:"HOST"`
    Port     int    `yaml:"port" env:"PORT"`
    Password string `yaml:"password" env:"DB_PASSWORD,required"`
}
```

**3. Load and use:**

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/yuseferi/envyaml"
)

func main() {
    var cfg Config
    
    if err := envyaml.LoadConfig("config.yml", &cfg); err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Connected to %s:%d\n", cfg.Host, cfg.Port)
}
```

## 🎨 Features

| Feature | Description |
|---------|-------------|
| 🔄 **Seamless Integration** | Combine YAML files with environment variables effortlessly |
| ✅ **Required Variables** | Mark critical env vars as required with automatic validation |
| 🏷️ **Struct Tags** | Use familiar `yaml` and `env` struct tags |
| 🛡️ **Type Safety** | Full Go type safety with automatic type conversion |
| 📦 **Zero Config** | Works out of the box with sensible defaults |
| 🪶 **Lightweight** | Minimal dependencies, maximum performance |

## 📚 Examples

### Required Environment Variables

Mark sensitive variables as required to fail fast if they're missing:

```go
type Config struct {
    Host     string `yaml:"host" env:"HOST"`
    Port     int    `yaml:"port" env:"PORT"`
    Password string `yaml:"password" env:"DB_PASSWORD,required"` // 👈 Required!
}

var cfg Config
err := envyaml.LoadConfig("config.yml", &cfg)
if err != nil {
    // Error: failed to parse environment variables: env: required environment variable "DB_PASSWORD" is not set
    log.Fatal(err)
}
```

### Complete Working Example

```go
package main

import (
    "fmt"
    "log"
    "os"
    
    "github.com/yuseferi/envyaml"
)

type DatabaseConfig struct {
    Host     string `yaml:"host" env:"DB_HOST"`
    Port     int    `yaml:"port" env:"DB_PORT"`
    Username string `yaml:"username" env:"DB_USER"`
    Password string `yaml:"password" env:"DB_PASSWORD,required"`
    Database string `yaml:"database" env:"DB_NAME"`
}

type Config struct {
    Database DatabaseConfig `yaml:"database"`
    Debug    bool           `yaml:"debug" env:"DEBUG"`
}

func main() {
    // Set environment variables (in production, these come from your environment)
    os.Setenv("DB_PASSWORD", "super_secret_password")
    
    var cfg Config
    if err := envyaml.LoadConfig("config.yml", &cfg); err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }
    
    fmt.Printf("Database: %s@%s:%d/%s\n", 
        cfg.Database.Username,
        cfg.Database.Host, 
        cfg.Database.Port,
        cfg.Database.Database,
    )
}
```

With this `config.yml`:

```yaml
database:
  host: localhost
  port: 5432
  username: admin
  password: ${DB_PASSWORD}
  database: myapp

debug: false
```

**Output:**
```
Database: admin@localhost:5432/myapp
```

## 🏗️ How It Works

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│   YAML File     │     │   Environment   │     │   Go Struct     │
│                 │     │   Variables     │     │                 │
│  host: localhost│     │                 │     │  Host: localhost│
│  port: 3606     │ ──► │  DB_PASSWORD=   │ ──► │  Port: 3606     │
│  password: ${..}│     │  "secret123"    │     │  Password: ...  │
└─────────────────┘     └─────────────────┘     └─────────────────┘
        │                       │                       │
        └───────────────────────┴───────────────────────┘
                            envYaml
```

1. **Read** - Parse your YAML configuration file
2. **Merge** - Overlay environment variables using struct tags
3. **Validate** - Ensure required variables are present
4. **Return** - Provide a fully populated, type-safe config struct

## 🛠️ Development

This project uses [Task](https://taskfile.dev) for managing development tasks.

```bash
# Build the project
task build

# Run tests
task test

# Run tests with coverage
task test-coverage

# Clean generated files
task clean

# Run all tasks
task all
```

## 🤝 Contributing

We love contributions! ❤️

1. 🍴 Fork the repository
2. 🌿 Create your feature branch (`git checkout -b feature/amazing-feature`)
3. 💾 Commit your changes (`git commit -m 'Add amazing feature'`)
4. 📤 Push to the branch (`git push origin feature/amazing-feature`)
5. 🎉 Open a Pull Request

Please feel free to:
- 🐛 Report bugs
- 💡 Suggest new features
- 📖 Improve documentation
- ⭐ Star the project if you find it useful!

## 📄 License

This project is licensed under the **GNU General Public License v3.0** - see the [LICENSE](LICENSE) file for details.

---

<div align="center">

**Made with ❤️ by [Yusef Mohamadi](https://github.com/yuseferi)**

If this project helped you, consider giving it a ⭐!

[![GitHub stars](https://img.shields.io/github/stars/yuseferi/envyaml?style=social)](https://github.com/yuseferi/envyaml/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/yuseferi/envyaml?style=social)](https://github.com/yuseferi/envyaml/network/members)

</div>
