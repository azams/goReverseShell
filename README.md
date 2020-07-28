# goReverseShell
Simple reverse shell using Golang.

# Installation:
* Install Golang - https://golang.org/doc/install
* git clone https://github.com/azams/goReverseShell
* Don't forget to change the IP and PORT

# Compile
Example for targeting Windows 32bit
> GOOS=windows GOARCH=386 go build goReverseShell.go

Example for targeting Windows 64bit
> GOOS=windows GOARCH=x64 go build goReverseShell.go

Example for targeting Linux 32bit
> GOOS=linux GOARCH=386 go build goReverseShell.go

Example for targeting Linux 64bit
> GOOS=linux GOARCH=x64 go build goReverseShell.go

Documentation: https://golang.org/doc/install/source#environment

VirusTotal: https://www.virustotal.com/gui/file/d72953e4cd473bd86607a3e31bd7207b410b1cd80194d6a637054c07c7c79b63/detection
