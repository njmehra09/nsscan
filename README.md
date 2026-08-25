🚀 NSSCAN

NSSCAN is a fast and lightweight HTTP/HTTPS host discovery and response scanner written in Go.

It supports domain, IP, CIDR, and file-based input and displays useful response information such as IP, Port, HTTP Status Code, and Server.

«⚠️ Authorized Use Only: Use NSSCAN only on systems, domains, IP addresses, and networks that you own or have explicit permission to test.»

---

✨ Features

- 🚀 Concurrent scanning
- 🌐 HTTP/HTTPS support
- 🔍 Domain scanning
- 🎯 IP address scanning
- 📡 CIDR range scanning
- 📂 File-based target input
- 🔌 Single custom-port scanning
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

📱 Requirements

NSSCAN can be built and used in Termux.

Requirements

- Android
- Termux
- Go
- Internet connection

---

📥 Installation

1. Install Termux packages

pkg update && pkg upgrade

Install Git and Go:

pkg install git golang

---

2. Clone the repository

git clone https://github.com/njmehra09/nsscan.git

Enter the project directory:

cd nsscan

---

3. Install Go dependencies

go mod tidy

---

4. Build NSSCAN

gofmt -w nsscan.go
go build -o ~/go/bin/nsscan nsscan.go

---

▶️ Run NSSCAN

Run directly:

~/go/bin/nsscan

If "~/go/bin" is available in your PATH:

nsscan

---

🎯 Target Input

NSSCAN supports several target types.

Domain

example.com

IP Address

192.0.2.10

CIDR

192.0.2.0/24

Example "/16":

192.0.0.0/16

A normal IPv4 "/16" contains 65,536 addresses.

File

You can also provide a ".txt" file containing targets.

Example:

example.com
example.org
192.0.2.10
192.0.2.0/24

---

🔌 Custom Port

NSSCAN supports scanning a single custom port.

Example:

8443

Other examples:

8080

8000

When a custom port is supplied, NSSCAN attempts HTTP first and then HTTPS if the HTTP request fails.

«Note: The current implementation is designed for one custom port at a time. Use separate scans for different custom ports.»

---

🌐 Default Ports

When the custom-port field is left blank, NSSCAN uses the normal HTTP/HTTPS behaviour:

HTTP  → 80
HTTPS → 443

If HTTP fails, HTTPS is attempted.

---

🧵 Threads

Default:

100

You can choose the number of concurrent workers during startup.

Example:

200

Use a reasonable value for your device and authorized network.

---

⏱️ Timeout

Default:

10 seconds

You can change the timeout during startup.

Example:

5

---

📡 HTTP Methods

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

Default method:

HEAD

---

🚦 Status Code Filtering

You can optionally skip specific HTTP status codes.

Example:

302,307

This prevents those selected responses from being included in the saved results.

Leave the field blank to keep all status codes.

---

🖥️ Terminal Output

NSSCAN displays:

IP               PORT CODE SERVER         HOST
------------------------------------------------
93.184.216.34    80   200  nginx          example.com

The scanner also displays CIDR loading progress:

🔍 Processing CIDR: 192.0.2.0/16
✅ Loaded 65536 IPs from this CIDR

---

💾 Saved Results

Results are saved automatically while scanning.

Example:

example.com | IP: 93.184.216.34 | PORT: 80 | CODE: 200 | SERVER: nginx

Each saved result contains:

Field| Description
Domain| Target hostname
IP| Resolved IP address
PORT| Responding port
CODE| HTTP status code
SERVER| Server response header

---

📁 Output File

During startup, NSSCAN asks for the output filename.

Example:

results.txt

You can also select whether the file should be saved in Termux or device storage.

---

🛑 Stop a Scan

Press:

CTRL + C

NSSCAN will stop gracefully and display the progress reached so far.

Results discovered before interruption are already saved to the output file.

---

🔧 Build From Source

After modifying the source code:

gofmt -w nsscan.go

Update dependencies:

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

📚 Dependencies

NSSCAN currently uses:

- fatih/color — terminal colors
- golang.org/x/term — terminal information

Dependencies are managed using Go Modules.

---

🧪 Example

Start:

nsscan

Enter:

example.com

Threads:

100

Timeout:

10

Output:

results.txt

Custom port:

8443

HTTP method:

HEAD

Skip codes:



The scanner will then test the target using the selected configuration.

---

🔄 Update From GitHub

If you already cloned the repository:

cd nsscan
git pull

Then rebuild:

gofmt -w nsscan.go
go mod tidy
go build -o ~/go/bin/nsscan nsscan.go

Run:

nsscan

---

⚠️ Responsible Use

NSSCAN is intended for:

- Your own servers
- Your own domains
- Authorized security testing
- Lab environments
- Educational networking
- Networks where you have explicit permission

Do not scan systems or networks without authorization.

The author is not responsible for misuse of this software.

---

👤 Author

NJ Mehra

GitHub:

https://github.com/njmehra09

Repository:

https://github.com/njmehra09/nsscan

---

📜 License

This project currently does not specify a license.

If you want others to legally reuse, modify, or distribute the project, add an appropriate open-source license to the repository.

---

⭐ Support

If you find NSSCAN useful, consider giving the repository a ⭐ on GitHub.

NSSCAN — Fast • Simple • Accurate
