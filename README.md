# PDF Toolbox

[![Go Report Card](https://goreportcard.com/badge/github.com/jojomak13/pdf-toolbox)](https://goreportcard.com/report/github.com/jojomak13/pdf-toolbox)
[![Latest Release](https://img.shields.io/github/v/release/jojomak13/pdf-toolbox)](https://github.com/jojomak13/pdf-toolbox/releases/latest)
[![License](https://img.shields.io/github/license/jojomak13/pdf-toolbox)](LICENSE)

A simple Go application for merging multiple PDF files into one with support for uploading to S3.

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

## Environment Variables

Set the following environment variables before running the application:

```
APP_NAME="PDF-Toolbox"
PORT=8080

# S3 Config
S3_KEY=""
S3_SECRET=""
S3_REGION=""
S3_BUCKET=""

OUTPUT_FILE_NAME="out.pdf"
```

## License

This project is licensed under the [MIT License](LICENSE).
