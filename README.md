certverify.go is a small go snippet that can be used to validate the certificate It will return, if the provided certificate is malformed or not.
To run this snippet ypou need to have the go lib installed on your machine.
### Download and Install

#### Binary Distributions
Official binary distributions are available at https://go.dev/dl/.
After downloading a binary release, visit https://go.dev/doc/install
for installation instructions.

### Check Go version 
Verify that you've installed Go by opening a command prompt and typing the following command:
$ go version

Confirm that the command prints the installed version of Go.

### To verify the certificate
Decode the base64 certificate and add the decodded certificate to the certverify.go in the 


### Run  the certverify.go
To run this snippet you simply use $ go run <go-file-name> 
for example:
$ go run certverify.go
