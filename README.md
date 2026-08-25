NSSCAN

NSSCAN is a fast HTTP/HTTPS host discovery and response-information scanner written in Go.

It can process individual domains, IP/CIDR ranges, and input files, then display and save useful HTTP response information such as IP address, port, status code, and server.

«⚠️ Authorized-use only: Use NSSCAN only against systems, domains, IP addresses, and networks that you own or have explicit permission to test.»

---

✨ Features

- 🚀 Concurrent scanning with configurable threads
- 🌐 HTTP and HTTPS checking
- 📡 Domain and IP target support
- 📦 CIDR range scanning
- 📄 File-based target input
- 🔢 Large CIDR support
- 🔌 Single custom-port scanning
- 🔀 Multiple HTTP methods
- 🚦 Configurable status-code filtering
- 🖥️ Server/banner detection
- 🌍 IP resolution for domain targets
- 💾 Real-time result saving
- 🎨 Color-coded terminal output
- ⏱️ Configurable request timeout
- 🛑 Graceful scan interruption

---

📋 Requirements

- Android + Termux or another supported Go environment
- Go
- Internet/network access
- Go modules

---

📥 Installation

Clone the repository:

git clone https://github.com/YOUR_USERNAME/YOUR_REPOSITORY.git
cd YOUR_REPOSITORY

Install dependencies:

go mod tidy

Build NSSCAN:

gofmt -w nsscan.go
go build -o ~/go/bin/nsscan nsscan.go

Run:

~/go/bin/nsscan

If "~/go/bin" is in your PATH:

nsscan

---

🚀 Usage

Run NSSCAN without arguments:

nsscan

The interactive interface will ask for:

1. Target
2. Number of threads
3. Timeout
4. Output filename
5. Output location
6. Custom port
7. HTTP method
8. Status codes to skip

---

🎯 Target Types

1. Single Domain

example.com

2. Single IP

192.0.2.10

3. CIDR Range

192.0.2.0/24

For example, a "/16" IPv4 range contains up to 65,536 addresses.

4. Target File

You can provide a text file containing domains, IPs, and CIDR ranges.

Example:

example.com
example.org
192.0.2.10
192.0.2.0/24

---

🔌 Custom Port

You can specify one custom port instead of the default HTTP/HTTPS ports.

Example:

8443

For a target, NSSCAN will attempt the custom port using HTTP and, if necessary, HTTPS.

Other examples:

8080
8000
8443

«Currently, use a single custom port per scan. Multiple comma-separated custom ports are not intended by the current implementation.»

---

🌐 Default Port Behaviour

When no custom port is specified, NSSCAN attempts:

HTTP  → 80
HTTPS → 443

If the HTTP request fails, it attempts HTTPS.

---

🧰 HTTP Methods

NSSCAN supports:

HEAD
GET
POST
PUT
DELETE
CONNECT
OPTIONS
TRACE
PATCH

Default:

HEAD

---

🚦 Status Code Filtering

You can specify status codes that should be skipped.

Example:

302,307

This allows those responses to be excluded from the saved results.

Leave the field blank if you don't want to skip status codes.

---

🧵 Threads

The default number of concurrent workers is:

100

You can choose another value during startup.

Example:

200

Higher concurrency can increase resource usage and may cause more connection failures, so use an appropriate value for your authorized testing environment.

---

⏱️ Timeout

Default timeout:

10 seconds

You can change it during startup.

Example:

5

---

💾 Output

NSSCAN saves successful responses to the selected output file.

Example:

example.com | IP: 93.184.216.34 | PORT: 80 | CODE: 200 | SERVER: nginx

The saved result contains:

Field| Description
Domain| Target hostname
IP| Resolved IP address
PORT| Port that responded
CODE| HTTP status code
SERVER| Server response header

---

🖥️ Terminal Output

The terminal displays information in a compact format:

IP               PORT CODE SERVER         HOST
------------------------------------------------
93.184.216.34    80   200  nginx          example.com

During CIDR processing, NSSCAN also reports how many addresses were loaded:

🔍 Processing CIDR: 192.0.2.0/16
✅ Loaded 65536 IPs from this CIDR

---

📂 Output Location

You can choose where the result file should be saved.

For example:

Termux

or:

Storage

The output filename can also be customized.

Example:

scan-results.txt

---

🛑 Interrupting a Scan

Press:

CTRL + C

NSSCAN will stop the scan and display the progress reached so far.

Already discovered results are written to the output file during scanning.

---

⚙️ Build From Source

Format the source:

gofmt -w nsscan.go

Download/update dependencies:

go mod tidy

Build:

go build -o ~/go/bin/nsscan nsscan.go

Run:

nsscan

---

📦 Project Structure

nsscan/
├── nsscan.go
├── go.mod
├── go.sum
└── README.md

---

🔧 Dependencies

NSSCAN currently uses:

- "github.com/fatih/color" — terminal colors
- "golang.org/x/term" — terminal information

Dependencies are managed through Go Modules.

---

📝 Example Workflow

Start the scanner:

nsscan

Enter a target:

example.com

Choose threads:

100

Choose timeout:

10

Choose output:

results.txt

Leave custom port blank to use the default HTTP/HTTPS behaviour.

Leave status-code filtering blank if no codes should be skipped.

---

⚠️ Responsible Use

NSSCAN is intended for:

- Your own infrastructure
- Authorized security testing
- Lab environments
- Educational networking experiments
- Systems where you have explicit permission

Do not use it to scan networks or services without authorization.

The author is not responsible for misuse of this software.

---

📜 License

Add your preferred open-source license here.

For example:

MIT License

---

👤 Author

NS Hacker / NSSCAN

Project:

NSSCAN

Telegram:

@NS_SCAN

«Replace the repository URL, author information, and license section with your actual project details before publishing.»
