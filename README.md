# Certificate Helper

Certificate Helper is a tool developed in Go (version 1.22) designed to simplify the generation and management of certificates for domains. Currently, the tool allows:

- Extracting certificates from domains.
- Converting certificates to useful formats such as Base64 and PEM.
- Saving generated certificates in an organized structure.

## Future Features

- **Integration with Cloudflare**: Query and manage certificates directly through the Cloudflare API.
- **API for domain queries**: Allow users to monitor their domain certificates in a centralized way.
- **Schedule Cron Jobs**: Allow users schedule cron jobs to refresh certificates before expiration date;
## Requirements

- Go 1.22 or higher.
- OpenSSL installed on the system.

## How to Use

1. Clone the repository:
   ```bash
   git clone https://github.com/Mamonha/certificate-helper.git
   ```
2. Navigate to the project directory:
   ```bash
   cd certificate-helper
   ```
3. Run the tool:
   ```bash
   go run main.go --domain=<domain> --format=<format>
   ```
   - Replace `<domain>` with the desired domain.
   - Replace `<format>` with `pem` or `b64` to choose the certificate format.

## Contribution

Contributions are welcome! Feel free to open issues and submit pull requests with improvements or new features.

---

**Note**: This project is under active development, and new features will be added soon.