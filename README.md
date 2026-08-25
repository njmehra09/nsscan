# 🚀 NSSCAN

**NSSCAN** is a fast and lightweight HTTP/HTTPS host discovery and response scanner written in **Go**.

It supports domain, IP, CIDR, and file-based input and displays useful response information such as **IP address, port, HTTP status code, and server**.

> ⚠️ **Authorized Use Only:** Use NSSCAN only on systems, domains, IP addresses, and networks that you own or have explicit permission to test.

---

## ✨ Features

- 🚀 Concurrent scanning
- 🌐 HTTP/HTTPS support
- 🔍 Domain scanning
- 🎯 IP address scanning
- 📡 CIDR range scanning
- 📂 File-based target input
- 🔌 Custom port scanning
- 📊 HTTP status-code detection
- 🖥️ Server header detection
- 🌍 Domain IP resolution
- 💾 Real-time result saving
- 🎨 Color-coded terminal output
- ⏱️ Configurable timeout
- 🧵 Configurable threads
- 🚦 Status-code filtering
- 🛑 Graceful scan interruption

---

## 📱 Requirements

NSSCAN can be built and used in **Termux**.

- Android
- Termux
- Git
- Go
- Internet connection

---

# 📥 Installation

### 1. Update Termux

```bash
pkg update && pkg upgrade
```

### 2. Install Git and Go

```bash
pkg install git golang
```

### 3. Clone the repository

```bash
git clone https://github.com/njmehra09/nsscan.git
```

### 4. Enter the project directory

```bash
cd nsscan
```

### 5. Install Go dependencies

```bash
go mod tidy
```

### 6. Format the source code

```bash
gofmt -w nsscan.go
```

### 7. Build NSSCAN

```bash
go build -o ~/go/bin/nsscan nsscan.go
```

---

# ▶️ Run NSSCAN

Run directly:

```bash
~/go/bin/nsscan
```

If `~/go/bin` is already in your PATH:

```bash
nsscan
```

---

# 🎯 Target Input

NSSCAN supports multiple target types.

### Domain

```text
example.com
```

### IP Address

```text
192.0.2.10
```

### CIDR Range

```text
192.0.2.0/24
```

Example `/16`:

```text
192.0.0.0/16
```

A normal IPv4 `/16` contains **65,536 addresses**.

### File

A `.txt` file can contain domains, IP addresses, and CIDR ranges:

```text
example.com
example.org
192.0.2.10
192.0.2.0/24
```

---

# 🔌 Custom Port

NSSCAN supports a custom port.

Example:

```text
8443
```

Another example:

```text
8080
```

When a custom port is supplied, NSSCAN attempts HTTP first and then HTTPS if the HTTP request fails.

> **Note:** The current implementation is designed for one custom port at a time.

---

# 🌐 Default Ports

When the custom-port field is left blank:

```text
HTTP  → 80
HTTPS → 443
```

If HTTP fails, HTTPS is attempted.

---

# 🧵 Threads

Default:

```text
100
```

You can choose the number of concurrent workers during startup.

Example:

```text
200
```

Use a reasonable value for your device and authorized testing environment.

---

# ⏱️ Timeout

Default:

```text
10 seconds
```

Example:

```text
5
```

---

# 📡 HTTP Methods

NSSCAN supports:

```text
HEAD
GET
POST
PUT
DELETE
CONNECT
OPTIONS
TRACE
PATCH
```

Default method:

```text
HEAD
```

---

# 🚦 Status Code Filtering

You can optionally skip specific HTTP status codes.

Example:

```text
302,307
```

Leave the field blank if you don't want to skip any status codes.

---

# 🖥️ Terminal Output

Example:

```text
IP               PORT CODE SERVER         HOST
------------------------------------------------
93.184.216.34    80   200  nginx          example.com
```

During CIDR processing:

```text
🔍 Processing CIDR: 192.0.2.0/16
✅ Loaded 65536 IPs from this CIDR
```

---

# 💾 Saved Results

Results are saved automatically while scanning.

Example:

```text
example.com | IP: 93.184.216.34 | PORT: 80 | CODE: 200 | SERVER: nginx
```

Saved information:

| Field | Description |
|---|---|
| Domain | Target hostname |
| IP | Resolved IP address |
| PORT | Responding port |
| CODE | HTTP status code |
| SERVER | Server response header |

---

# 📁 Output File

During startup, NSSCAN asks for an output filename.

Example:

```text
results.txt
```

You can choose whether the file should be saved in Termux or device storage.

---

# 🛑 Stop a Scan

Press:

```text
CTRL + C
```

NSSCAN will stop gracefully and display the progress reached so far.

Results discovered before interruption are already saved to the output file.

---

# 🔧 Build From Source

After modifying the source code:

```bash
gofmt -w nsscan.go
```

Update dependencies:

```bash
go mod tidy
```

Build:

```bash
go build -o ~/go/bin/nsscan nsscan.go
```

Run:

```bash
nsscan
```

---

# 🔄 Update NSSCAN

If you already cloned the repository:

```bash
cd nsscan
```

Pull the latest changes:

```bash
git pull
```

Update dependencies:

```bash
go mod tidy
```

Format the source:

```bash
gofmt -w nsscan.go
```

Rebuild:

```bash
go build -o ~/go/bin/nsscan nsscan.go
```

Run:

```bash
nsscan
```

---

# 📦 Project Structure

```text
nsscan/
├── nsscan.go
├── go.mod
├── go.sum
└── README.md
```

---

# 📚 Dependencies

NSSCAN currently uses:

- `github.com/fatih/color`
- `golang.org/x/term`

Dependencies are managed using **Go Modules**.

---

# 🧪 Example Workflow

Start NSSCAN:

```bash
nsscan
```

Enter a target:

```text
example.com
```

Threads:

```text
100
```

Timeout:

```text
10
```

Output file:

```text
results.txt
```

Custom port:

```text
8443
```

HTTP method:

```text
HEAD
```

Skip status codes:

```text
302,307
```

---

# ⚠️ Responsible Use

NSSCAN is intended for:

- Your own servers
- Your own domains
- Authorized security testing
- Lab environments
- Educational networking
- Networks where you have explicit permission

Do **not** scan systems or networks without authorization.

The author is not responsible for misuse of this software.

---

# 👤 Author

**NJ Mehra**

GitHub:

https://github.com/njmehra09

Repository:

https://github.com/njmehra09/nsscan

---

# 📜 License

This project currently does not specify a license.

If you want others to legally reuse, modify, or distribute the project, add an appropriate open-source license.

---

## ⭐ Support

If you find NSSCAN useful, consider giving the repository a ⭐ on GitHub.

**NSSCAN — Fast • Simple • Accurate**
