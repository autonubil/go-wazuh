# Project Identity
### Project Type
The project appears to be an **App**, specifically focusing on the Wazuh management system.

### Primary Language
The primary language used in this repository is **Go** (Golang).

### Core Tech Stack
- **Go**: The main programming language.
- **OpenAPI/Swagger**: Used for defining the API specifications, as evident from the `spec.yaml` file.
- **GitHub Actions**: CI/CD workflows are defined using GitHub Actions, specifically seen in the `.github/workflows/go.yml`.
- **Grafana**: Integration with Grafana is indicated by the presence of a dashboard configuration file (`grafana/wodle-agent-dashboard.json`).
- **Prometheus**: Prometheus integration is suggested by the `prometheus.go` file under the `ossec/` directory.

### Main Purpose
This repository contains the core components and configurations for a Go-based application that interfaces with the Wazuh management system. It includes API definitions, client generation scripts, configuration files, and various modules responsible for managing agents, handling security policies, and integrating with monitoring tools like Grafana and Prometheus. The project aims to provide a robust framework for automating and managing Wazuh operations through a RESTful API.

# Architecture & Logic
Based on the provided directory structure and file contents, this codebase appears to follow a **Modular** architectural style. The modular architecture is evident from the separation of concerns across different directories, each representing a distinct module with specific functionalities (e.g., `ossec`, `rest`, `sysinfo`, `wazuh`). Each module contains related files and components that perform tasks associated with their domain.

### Key Characteristics of Modular Architecture:
- **Separation of Concerns**: The codebase is divided into multiple directories, each responsible for a specific aspect of the application (e.g., security management, data indexing, system information).
- **Reusability**: Modules can be reused across different parts of the application or even in other applications.
- **Maintainability**: Changes to one module do not affect others as long as the interface remains consistent.

### Naming Conventions and Patterns:
1. **File Naming Conventions**:
   - **Descriptive File Names**: Files have descriptive names that indicate their purpose (e.g., `agent_keys.go` for handling agent keys, `indexer.go` for indexing functionality).
   - **Test File Naming**: Test files typically end with `_test.go`, following Go's convention (e.g., `agent_test.go`, `wazuh_gen_code_test.go`). This makes it easy to identify and run tests.

2. **Directory Structure**:
   - **Logical Grouping**: Directories are logically grouped based on functionality, making it easy for developers to find related files.
   - **Consistent Path Structure**: Each module has a consistent path structure within the repository, aiding in navigation and organization (e.g., `rest/`, `ossec/`).

3. **Go Package Naming**:
   - **Package Names Reflect Directory Names**: The package names in Go files typically reflect their directory names or functionality, ensuring clarity and consistency.
     // Example from ossec/authd.go
     package ossec

     // Example from rest/wazuh.go
     package rest

### Additional Observations:
- **Use of Interfaces**: The presence of interface files (e.g., `client_interfaces.go`) suggests that the codebase uses interfaces to define and enforce contracts between modules, which is a common pattern in modular architectures.
- **Test Coverage**: There are test files for most of the main logic files (e.g., `agent_test.go` alongside `agent.go`), indicating a focus on testing and ensuring quality.

By adhering to these conventions and patterns, the codebase maintains clarity, maintainability, and scalability. This approach is particularly beneficial in larger projects where multiple developers may be working simultaneously.

# Coding Guidelines
### Proactive Coding Guidelines

1. **Avoid Weak Cryptographic Algorithms:**
   - **Guideline:** Do not use MD5 or other weak cryptographic hash algorithms (e.g., SHA1) in your code.
   - **Rationale:** MD5 and similar algorithms are vulnerable to collision attacks, making them unsuitable for secure cryptographic operations. Use stronger alternatives like SHA256 or SHA3 instead.
   - **Example Violation:**
     // Incorrect Usage:
     hash := md5.New()
   - **Correct Usage:**
     // Correct Usage:
     hash := sha256.New()

2. **Use Secure Randomness for Cryptographic Operations:**
   - **Guideline:** Use `crypto/rand` instead of `math/rand` for generating random numbers in cryptographic contexts.
   - **Rationale:** The `math/rand` package is not cryptographically secure and should only be used for non-security purposes. For security-sensitive code, use the `crypto/rand` package to ensure true randomness.
   - **Example Violation:**
     // Incorrect Usage:
     rand.Seed(time.Now().UnixNano())
     randomValue := rand.Intn(100)
   - **Correct Usage:**
     // Correct Usage:
     var randomBytes [4]byte
     _, err := crypto/rand.Read(randomBytes[:])
     if err != nil {
         log.Fatal(err)
     }
     randomValue := int(binary.LittleEndian.Uint32(randomBytes[:]))

3. **Ensure TLS Configuration Uses Secure Protocols:**
   - **Guideline:** Set the `MinVersion` to `tls.VersionTLS13` in your TLS configurations to enforce the use of secure protocols.
   - **Rationale:** Older versions of TLS (like TLS 1.0 and 1.1) are outdated and vulnerable to various attacks. Modern applications should default to TLS 1.3, with older versions disabled unless there is a specific, documented need for backward compatibility with legacy systems.
   - **Example Violation:**
     // Incorrect Usage:
     config := &tls.Config{
         InsecureSkipVerify: true,
     }
   - **Correct Usage:**
     // Correct Usage:
     config := &tls.Config{
         MinVersion: tls.VersionTLS13,
         CipherSuites: []uint16{
             tls.TLS_AES_256_GCM_SHA384,
             tls.TLS_CHACHA20_POLY1305_SHA256,
             tls.TLS_AES_128_GCM_SHA256,
         },
     }

By adhering to these guidelines, the code will be more secure and resistant to common cryptographic vulnerabilities.