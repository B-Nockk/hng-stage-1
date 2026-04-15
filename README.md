# HNG Stage 1: Go Minimalist API

A fast, zero-dependency Go API built for the HNG DevOps track. Containerized with a multi-stage Docker build and served behind an Nginx reverse proxy.

**Live URL:** <https://hng14-stage0-nockk.duckdns.org>

---

## Running Locally

First, clone the repository:

```bash
git clone https://github.com/b-nockk/hng-stage-1.git
cd hng-stage-1
```

Then choose one of the following options.

### **Option A — Docker (Recommended)**

**Requires:** Docker and Docker Compose


```bash
docker compose up -d --build
```

### **Option B — Go Directly**

**Requires:** Go 1.24+

```bash
go run main.go
```
Both options expose the API at `http://localhost:8080`.

---

## Endpoints

### `GET /`

Returns a status message confirming the API is active.

```json
{
  "message": "API is running"
}
```

### `GET /health`

Returns the operational health status of the service.

```json
{
  "message": "healthy"
}
```

### `GET /me`

Returns the developer's profile information.

```json
{
  "email": "ob.ogheneochuko@gmail.com",
  "github": "https://github.com/B-Nockk",
  "name": "Benson OgheneOchuko"
}
```
