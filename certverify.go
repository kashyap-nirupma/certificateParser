package main

import (
    "crypto/x509"
    "encoding/pem"
    "fmt"
)
// 	// Verifying with a custom list of root certificates. enter the certificate data below to verify.
func main() {
    cert := []byte(`-----BEGIN CERTIFICATE-----
-----END CERTIFICATE-----`)
    for len(cert) > 0 {
        b, next := pem.Decode(cert)
        if b == nil {
            fmt.Println("Unable to decode cert into a pem block. Cert is either empty or invalid.", string(cert))
            break
        }
        _, err := x509.ParseCertificate(b.Bytes)
        if err != nil {
            fmt.Println("Malformed Cert, not syncing. Err:", err)
            cert = next
            continue
        }
        cert = next
        fmt.Println("Success")
    }
    fmt.Println("End")
}
 
