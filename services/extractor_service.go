package services

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func GenerateCertificateFile(domain string) error {
	fmt.Printf("Generating certificate file for domain: %s\n", domain)

	certsDir := "certs"
	if err := os.MkdirAll(certsDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create certs directory: %v", err)
	}

	certFilename := filepath.Join(certsDir, "cert.cer")
	cmd1 := exec.Command("bash", "-c",
		fmt.Sprintf("echo | openssl s_client -connect %s:443 -servername %s 2>/dev/null | openssl x509 > %s", domain, domain, certFilename))

	err := cmd1.Run()
	if err != nil {
		return fmt.Errorf("failed to get certificate: %v", err)
	}

	if _, err := os.Stat(certFilename); os.IsNotExist(err) {
		return fmt.Errorf("certificate file was not created")
	}

	base64Filename := filepath.Join(certsDir, "cert.cer.b64")
	cmd2 := exec.Command("bash", "-c",
		fmt.Sprintf("base64 %s > %s", certFilename, base64Filename))

	err = cmd2.Run()
	if err != nil {
		return fmt.Errorf("failed to convert certificate to base64: %v", err)
	}

	if _, err := os.Stat(base64Filename); os.IsNotExist(err) {
		return fmt.Errorf("base64 certificate file was not created")
	}

	fmt.Printf("Certificate saved to: %s\n", certFilename)
	fmt.Printf("Base64 certificate saved to: %s\n", base64Filename)

	return nil
}

func ConvertCertificateToPEM(inputFile string, outputFile string) error {
	certsDir := "certs"
	if err := os.MkdirAll(certsDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create certs directory: %v", err)
	}

	inputFile = filepath.Join(certsDir, inputFile)
	outputFile = filepath.Join(certsDir, outputFile)
	fmt.Printf("Converting %s to PEM format: %s\n", inputFile, outputFile)

	for i := 0; i < 5; i++ {
		if _, err := os.Stat(inputFile); err == nil {
			break
		}
		time.Sleep(1 * time.Second)
		if i == 4 {
			return fmt.Errorf("input file %s not found after waiting", inputFile)
		}
	}

	cmd := exec.Command("bash", "-c",
		fmt.Sprintf("openssl x509 -in %s -out %s -outform PEM", inputFile, outputFile))

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to convert certificate to PEM: %v", err)
	}

	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		return fmt.Errorf("PEM file was not created")
	}

	fmt.Printf("PEM certificate saved to: %s\n", outputFile)
	return nil
}
