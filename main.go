package main

import (
	"certificate-extractor/services"
	"flag"
	"fmt"
	"os"
)

func main() {
	domain := flag.String("domain", "", "Domain to extract the certificate from")
	format := flag.String("format", "pem", "Output format: pem or b64")
	flag.Parse()

	if *domain == "" {
		fmt.Println("Error: --domain flag is required")
		os.Exit(1)
	}

	fmt.Printf("Extracting certificate for domain: %s\n", *domain)

	switch *format {
	case "pem":
		pemFile := "cert.pem"
		if err := services.GenerateCertificateFile(*domain); err != nil {
			fmt.Printf("Error generating certificate: %v\n", err)
			os.Exit(1)
		}
		if err := services.ConvertCertificateToPEM("cert.cer", pemFile); err != nil {
			fmt.Printf("Error converting to PEM: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("PEM certificate saved to: %s\n", pemFile)
	case "b64":
		if err := services.GenerateCertificateFile(*domain); err != nil {
			fmt.Printf("Error generating certificate: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Base64 certificate saved to: cert.cer.b64")
	default:
		fmt.Println("Error: Unsupported format. Use 'pem' or 'b64'")
		os.Exit(1)
	}
}
