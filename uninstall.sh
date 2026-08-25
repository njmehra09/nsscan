#!/data/data/com.termux/files/usr/bin/bash

echo "[+] Removing NSSCAN..."

rm -f "$HOME/.local/bin/nsscan"

echo
echo "[+] NSSCAN binary removed."
echo

echo "If you also want to remove the source repository, run:"
echo
echo "    rm -rf ~/nsscan"
echo
echo "Done!"
