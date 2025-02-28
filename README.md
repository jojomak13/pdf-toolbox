# PDF Toolbox

[![Go Report Card](https://goreportcard.com/badge/github.com/jojomak13/pdf-toolbox)](https://goreportcard.com/report/github.com/jojomak13/pdf-toolbox)
[![Latest Release](https://img.shields.io/github/v/release/jojomak13/pdf-toolbox)](https://github.com/jojomak13/pdf-toolbox/releases/latest)
[![License](https://img.shields.io/github/license/jojomak13/pdf-toolbox)](LICENSE)
[![Docker Pulls](https://img.shields.io/docker/pulls/jojomak/pdf-toolbox)](https://hub.docker.com/r/jojomak/pdf-toolbox)
[![Docker Image Size](https://img.shields.io/docker/image-size/jojomak/pdf-toolbox)](https://hub.docker.com/r/jojomak/pdf-toolbox)
[![Docker Stars](https://img.shields.io/docker/stars/jojomak/pdf-toolbox)](https://hub.docker.com/r/jojomak/pdf-toolbox)

A simple Go application for merging multiple PDF files into one and generate pdf from html with support for uploading to S3.

## Installation

Download the latest release from [here](https://github.com/jojomak13/pdf-toolbox/releases/latest).

## Usage

### Endpoint: `POST /merge`

#### Request Body
```json
{
    "file_path": "invoices/02-2025-20-18_0_invoices.pdf",
    "urls": [
        "https://pdfobject.com/pdf/sample.pdf",
        "https://my-storage.amazonaws.com/16-02-2025-20-18_0_sample.pdf"
    ]
}
```

#### Response (Success)
```json
{
    "data": {
        "url": "https://my-storage.amazonaws.com/invoices/02-2025-20-18_0_invoices.pdf"
    },
    "message": "success",
    "status": true
}
```

### Endpoint: `POST /html`

This endpoint accepts raw HTML content and converts it into a PDF file.

#### Headers
- `FILE-PATH`: The path where the generated PDF will be saved.

#### Request Body
```html
<html>
<head>
    <title>Sample PDF</title>
</head>
<body>
    <h1>Hello, PDF!</h1>
</body>
</html>
```

#### Response (Success)
```json
{
    "data": {
        "url": "https://my-storage.amazonaws.com/generated/sample.pdf"
    },
    "message": "success",
    "status": true
}
```

## 🐳 Docker Deployment

### Quick Start
```bash
docker-compose up -d
```

### Prerequisites
- Docker and Docker Compose installed on your system
- AWS S3 credentials (for storage functionality)

### Configuration
Create a `.env` file in your project root with the following variables:

```env
# AWS S3 Configuration (Required)
S3_KEY=your-aws-access-key
S3_SECRET=your-aws-secret-key
S3_REGION=your-aws-region
S3_BUCKET=your-bucket-name
```

### Environment Variables

|Variable|Required|Default|Description|
|----------|----------|---------|-------------|
| APP_NAME | No | PDF-Toolbox | Application name |
| PORT | No | 8080 | Port the application listens on |
| S3_KEY | Yes | - | AWS Access Key ID |
| S3_SECRET | Yes | - | AWS Secret Access Key |
| S3_REGION | Yes | - | AWS Region |
| S3_BUCKET | Yes | - | S3 Bucket name |
| OUTPUT_FILE_NAME | No | out.pdf | Default output filename |

## License

This project is licensed under the [MIT License](LICENSE).
