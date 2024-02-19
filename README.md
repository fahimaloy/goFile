# GoFile - Simple File Server with Systemd Integration

GoFile is a lightweight Go package that provides a simple file server to serve static files over HTTP. This package is designed to offer basic functionality, making it convenient to serve files on your VPS for website hosting or access files locally across different devices.

## Basic Functionalities

1. **Serve Static Files**: GoFile allows you to easily serve static files from a specified directory over HTTP. This is particularly useful for hosting files on your VPS and providing direct links to access them.

2. **Access Files Locally**: When used locally, GoFile enables you to access desired files from your laptop on any device within your local network.

## Installation

To use GoFile, you need to have Go installed on your system. If Go is not installed, you can install it on Ubuntu using the following command:

```bash
sudo apt install golang
```

Install the required dependency, godotenv, using the following command:

```bash
go install github.com/joho/godotenv@latest
```

Clone the GoFile repository to your local machine:

```bash
git clone https://github.com/fahimaloy/goFile.git
```

Navigate to the project directory:
```bash
cd goFile
```

Build the executable file:

```bash
go build main.go
```

#Configuration
Create a `.env` file in the project directory with the following content:

```
MEDIA_DIR=/path/to/your/files
SERVER_ADDR=localhost:8080
```

Adjust the values accordingly, specifying the path to your media files and the desired server address.

#Usage
Run the file server using the following command:

```bash
./main
```
This will start the server, and you can access your files by navigating to http://localhost:8080 in your web browser.


#Systemd Integration (Ubuntu)
To run the file server as a systemd service on Ubuntu, follow these steps:

1. Create a systemd service file:
```bash
sudo nano /etc/systemd/system/gofile.service
```
2. Add the following content to the file, adjusting the paths as necessary:

```
[Unit]
Description=GoFile Server

[Service]
ExecStart=/path/to/your/goFile/main
WorkingDirectory=/path/to/your/goFile
Restart=always

[Install]
WantedBy=multi-user.target
```
3.Save and exit the text editor.

4.Reload systemd to apply changes:
``` bash
sudo systemctl daemon-reload
```
5. Start and enable the file server service:
``` bash
sudo systemctl start gofile
sudo systemctl enable gofile
```
Now, GoFile will serve your files in the background and start automatically on system boot.

Feel free to customize the configuration and systemd service file based on your preferences and system setup.

#####Thanks from Fahim Ahmed & Omicronical