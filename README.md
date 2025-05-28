# Backdown 🗂️⬇️

A simple utility for backing up directories and uploading the resulting archive to a WebDAV server ☁️📦

To run it, you must provide the following flags:

* `--folder` — the directory to be archived  
* `--username` — your WebDAV username  
* `--password` — your WebDAV password  
* `--webdav-url` — the target WebDAV folder where the archive will be uploaded

## Project Build 🛠️

* `make build linux` builds the utility for Linux (x86_64) 🐧  
* `make build darwin` builds the utility for macOS with M1 chips 🍏  
* `make build windows` builds the utility for Windows (x86_64) 🪟

## Installation via GitHub 📥

```bash
go install github.com/pokatomnik/backdown
```
