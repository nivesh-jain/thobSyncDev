# **ThobSync CLI Tool**

ThobSync is a CLI tool designed for managing file uploads, downloads, and role-based access control (RBAC) for collaborative environments. It supports operations with MinIO and provides secure role management via API tokens.

---

## **Features**
- **File Management:**
  - Upload files to MinIO buckets.
  - Download files from MinIO buckets.
  - List files and buckets.
- **Role-Based Access Control (RBAC):**
  - Roles: `Admin`, `Editor`, `Viewer`.
  - API token system for secure role assignment.
- **Admin Management:**
  - Generate tokens for specific roles.
  - List all tokens and their associated roles.
  - Assign roles to users.
- **Configuration:**
  - Manage settings via `config.yaml`.
  - Secure user and role management with `users.json` and `tokens.json`.

---

## **Repository Structure**

```
thobSyncDev/
├── cmd/                      # CLI commands
│   ├── generate-token.go     # Command to generate API tokens
│   ├── list-tokens.go        # Command to list all tokens
│   ├── init.go               # Command to initialize CLI with API token
│   ├── upload-file.go        # Command to upload files
│   ├── download-file.go      # Command to download files
│   └── root.go               # Root command and initialization
│
├── config/                   # Configuration management
│   └── config.go             # Viper setup and default values
│
├── internal/                 # Internal logic
│   ├── auth/                 # Authentication and token management
│   │   ├── auth.go           # Authentication logic
│   │   ├── tokens.go         # Token validation logic
│   ├── minio/                # MinIO operations
│   │   ├── client.go         # MinIO client initialization
│   │   └── operations.go     # File and bucket operations
│
├── config.yaml               # Configuration file
├── users.json                # User credentials and roles
├── tokens.json               # Generated API tokens
├── roles.json                # Role-permission mappings
├── main.go                   # Entry point of the application
├── README.md                 # Documentation
├── go.mod                    # Go module file
└── go.sum                    # Dependency lock file
```

---

## **Installation**

### **Prerequisites**
- Install [Go](https://golang.org/doc/install) (version 1.18 or higher).
- Docker (if using MinIO via Docker).

### **Setup MinIO**
1. Run MinIO using Docker:
   ```bash
   docker run -p 9000:9000 -p 9001:9001 \
     -e MINIO_ROOT_USER=admin \
     -e MINIO_ROOT_PASSWORD=admin123 \
     quay.io/minio/minio server /data --console-address ":9001"
   ```
2. Access the MinIO console at [http://localhost:9001](http://localhost:9001).

---

## **Usage**

### **1. Initialize the CLI**
Authenticate with your API token:
```bash
go run main.go init --token <your-api-token>
```

### **2. File Management**
- **Upload a File:**
  ```bash
  go run main.go upload-file --bucket <bucket-name> --file <file-path>
  ```
- **Download a File:**
  ```bash
  go run main.go download-file --bucket <bucket-name> --object <object-name> --dest <destination-path>
  ```
- **List Files in a Bucket:**
  ```bash
  go run main.go list-files --bucket <bucket-name>
  ```

### **3. Admin Operations**
- **Generate a Token:**
  ```bash
  go run main.go generate-token --role <role-name>
  ```
- **List All Tokens:**
  ```bash
  go run main.go list-tokens --admin <admin-username> --password <admin-password>
  ```
- **Assign a Role:**
  ```bash
  go run main.go assign-role --admin <admin-username> --password <admin-password> --username <user> --role <role>
  ```

---

## **Configuration**

### **Config Files**
1. **`config.yaml`:**
   Stores CLI configuration such as MinIO endpoint and user role.
   ```yaml
   minio:
     endpoint: localhost:9000
     accessKeyID: admin
     secretAccessKey: admin123
     useSSL: false
   ```

2. **`users.json`:**
   Stores user credentials and roles.
   ```json
   {
       "users": {
           "admin": {
               "password": "hashed_admin_password",
               "role": "Admin"
           }
       }
   }
   ```

3. **`tokens.json`:**
   Manages generated tokens and associated roles.
   ```json
   {
       "tokens": {
           "admin-token-123": "Admin",
           "editor-token-456": "Editor"
       }
   }
   ```

---

## **Development**

### **Run the CLI Tool**
```bash
go run main.go <command>
```

### **Run Tests**
Add unit tests and run them:
```bash
go test ./...
```

---

## **Contributing**
1. Fork the repository.
2. Create a new branch for your feature.
3. Commit and push your changes.
4. Submit a pull request.

---

## **License**
This project is licensed under the MIT License.
