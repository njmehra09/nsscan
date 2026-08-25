#!/data/data/com.termux/files/usr/bin/bash

set -e

REPO="https://github.com/njmehra09/nsscan.git"
INSTALL_DIR="$HOME/.local/bin"

echo "[+] Installing NSSCAN..."

pkg update -y
pkg install -y git golang

TMP_DIR="$(mktemp -d)"
trap 'rm -rf "$TMP_DIR"' EXIT

echo "[+] Downloading NSSCAN..."

git clone "$REPO" "$TMP_DIR/nsscan"

cd "$TMP_DIR/nsscan"

echo "[+] Downloading Go dependencies..."
go mod download

echo "[+] Building NSSCAN..."
go build -o nsscan nsscan.go

mkdir -p "$INSTALL_DIR"

cp nsscan "$INSTALL_DIR/nsscan"
chmod +x "$INSTALL_DIR/nsscan"

if ! grep -qxF 'export PATH="$HOME/.local/bin:$PATH"' "$HOME/.bashrc" 2>/dev/null; then
    echo 'export PATH="$HOME/.local/bin:$PATH"' >> "$HOME/.bashrc"
fi

export PATH="$HOME/.local/bin:$PATH"

echo
echo "================================"
echo "  NSSCAN installed successfully!"
echo "================================"
echo
echo "Run NSSCAN with:"
echo
echo "    nsscan"
echo
