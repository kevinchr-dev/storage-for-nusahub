# Simple Object Storage API

Project sederhana untuk object storage menggunakan Go Fiber yang memungkinkan user untuk upload file dan mendapatkan URL untuk mengakses file tersebut.

## 🚀 Features

- ✅ Upload file dengan mudah
- ✅ Mendapatkan URL untuk akses file
- ✅ Download file
- ✅ List semua file yang tersimpan
- ✅ Delete file
- ✅ Support file size hingga 100MB
- ✅ Docker & Docker Compose ready

## 📋 Prerequisites

- Docker & Docker Compose (untuk menjalankan dengan Docker)
- Go 1.21+ (jika ingin menjalankan tanpa Docker)

## 🏃‍♂️ Cara Menjalankan

### Dengan Docker Compose (Recommended)

1. Clone atau download project ini

2. Jalankan dengan Docker Compose:
```bash
docker-compose up -d
```

3. API akan berjalan di `http://localhost:3000`

4. Untuk melihat logs:
```bash
docker-compose logs -f
```

5. Untuk stop aplikasi:
```bash
docker-compose down
```

### Tanpa Docker

1. Install dependencies:
```bash
go mod download
```

2. Jalankan aplikasi:
```bash
go run main.go
```

3. API akan berjalan di `http://localhost:3000`

## 📝 API Endpoints

### 1. Home
```
GET /
```
Response:
```json
{
  "message": "Welcome to Simple Object Storage API",
  "version": "1.0.0"
}
```

### 2. Upload File
```
POST /upload
Content-Type: multipart/form-data
```

Parameter:
- `file`: File yang akan diupload

Contoh menggunakan curl:
```bash
curl -X POST http://localhost:3000/upload \
  -F "file=@/path/to/your/file.jpg"
```

Response:
```json
{
  "message": "File uploaded successfully",
  "filename": "550e8400-e29b-41d4-a716-446655440000.jpg",
  "original_name": "file.jpg",
  "size": 123456,
  "url": "http://localhost:3000/files/550e8400-e29b-41d4-a716-446655440000.jpg",
  "storage_url": "http://localhost:3000/storage/550e8400-e29b-41d4-a716-446655440000.jpg",
  "uploaded_at": "2025-10-21T10:30:00Z"
}
```

### 3. Download File
```
GET /files/:filename
```

Contoh:
```bash
curl http://localhost:3000/files/550e8400-e29b-41d4-a716-446655440000.jpg -o downloaded-file.jpg
```

Atau buka langsung di browser:
```
http://localhost:3000/files/550e8400-e29b-41d4-a716-446655440000.jpg
```

### 4. List All Files
```
GET /files
```

Response:
```json
{
  "total_files": 2,
  "files": [
    {
      "filename": "550e8400-e29b-41d4-a716-446655440000.jpg",
      "size": 123456,
      "modified": "2025-10-21T10:30:00Z",
      "url": "http://localhost:3000/files/550e8400-e29b-41d4-a716-446655440000.jpg"
    }
  ]
}
```

### 5. Delete File
```
DELETE /files/:filename
```

Contoh:
```bash
curl -X DELETE http://localhost:3000/files/550e8400-e29b-41d4-a716-446655440000.jpg
```

Response:
```json
{
  "message": "File deleted successfully",
  "filename": "550e8400-e29b-41d4-a716-446655440000.jpg"
}
```

### 6. Akses File Langsung (Static)
```
GET /storage/:filename
```

Anda juga bisa mengakses file langsung melalui browser:
```
http://localhost:3000/storage/550e8400-e29b-41d4-a716-446655440000.jpg
```

## 📁 Struktur Project

```
storage/
├── handlers/
│   └── file_handler.go    # Handler untuk upload/download/list/delete
├── uploads/                # Folder penyimpanan file (auto-created)
├── main.go                 # Entry point aplikasi
├── go.mod                  # Go dependencies
├── go.sum                  # Go dependencies checksum
├── Dockerfile              # Docker configuration
├── docker-compose.yml      # Docker Compose configuration
└── README.md              # Dokumentasi
```

## 🔧 Environment Variables

- `PORT`: Port untuk menjalankan aplikasi (default: 3000)

## 📦 Persistent Storage

File yang diupload akan disimpan di folder `./uploads` dan akan tetap ada meskipun container di-restart karena menggunakan volume mounting di Docker Compose.

## 🛠 Development

Untuk development, Anda bisa menjalankan dengan hot reload menggunakan Air:

1. Install Air:
```bash
go install github.com/cosmtrek/air@latest
```

2. Jalankan:
```bash
air
```

## 📄 License

MIT License

## 👨‍💻 Author

Simple Object Storage API - Built with ❤️ using Go Fiber
