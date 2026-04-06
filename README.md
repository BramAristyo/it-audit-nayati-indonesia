# IT Audit Collector

A high-performance Windows system and hardware information collector built with Go. Designed for internal IT infrastructure auditing at **Nayati Indonesia**, this tool automates the gathering of detailed hardware specifications and system configurations, storing them directly into a centralized MySQL database.

## Key Features

- **Hardware Inventory:** Automatic detection of CPU, RAM, GPU/VGA, Storage (Disk), and Monitors.
- **System Information:** Gathers Motherboard, BIOS, and OS-level details.
- **Peripheral Audit:** Support for connected printer identification.
- **Database Integration:** Seamlessly pushes collected data to MySQL for reporting and analysis.
- **Windows 7 Support:** Engineered to run on legacy systems (Windows 7 SP1+) as well as modern versions.

## Tech Stack

- **Language:** Go (Golang)
- **Database:** MySQL 5.7+
- **Infrastructure:** Docker (for containerized builds)
- **Legacy Compatibility:** Optimized for Windows 7 SP1+ (requires Go 1.20.14 for builds).

## Quick Start

### Build with Docker (Recommended)
This method ensures the correct environment and dependencies are used for the build process.
```cmd
REM Run the provided build script
build.bat
```
The executable will be generated at `build\it-audit-collector.exe`.

### Manual Build
Ensure you have **Go 1.20.14** installed to maintain compatibility with Windows 7.
```cmd
go mod download
go build -o it-audit-collector.exe .\cmd\app\main.go
```

## Configuration

The application currently supports configuration via environment variables:

- `DB_HOST`: Database host (default: `127.0.0.1:3306`)
- `DB_USER`: Database user (default: `root`)
- `DB_PASS`: Database password (default: empty)
- `DB_NAME`: Database name (default: `db_it_audit`)

## Project Structure

```text
it-audit-collector/
├── cmd/app/              # Application entry point
├── internal/
│   ├── collector/        # Hardware and system data collection logic
│   ├── config/           # Configuration management
│   ├── database/         # MySQL connection and repository operations
│   └── sender/           # Data transmission logic (HTTP/Database)
├── models/               # Data structures and models
└── build/                # Output directory for compiled binaries
```

## Roadmap & Future Improvements

This project was developed rapidly for internal use at **Nayati Indonesia**. As a result, there is several technical debt items currently being addressed:

- [ ] **Enhanced Error Handling:** Improve system-wide error catching and logging for better troubleshooting in diverse environments.
- [ ] **Robust Configuration Management:** Move away from hardcoded defaults in `internal\database\connection.go` to a more flexible configuration system (e.g., YAML/JSON files).
- [ ] **Security Hardening:** Implement secure handling of database credentials. Currently, some sensitive data is hardcoded for development speed; this will be transitioned to a secure environment variable or vault-based approach.
- [ ] **Network Resilience:** Improve the data transmission logic to handle intermittent network connectivity during audits.

## License

**Internal Use Only** - Developed for Nayati Indonesia IT Audit purposes.
