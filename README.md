# Keshuook Web Archive

This is a Go-based tool that lets you browse old versions of [keshuook.github.io](https://keshuook.github.io). Instead of crawling the site, it pulls historical files directly from GitHub's commit history and serves them through a custom router.

## How it works

Most archives save static snapshots. This project uses your Git repo as the database. When you request a date, the backend fetches the files from that specific commit in real-time.

### Key Features

* **Link Fixing:** Injects a JS script into every HTML page to rewrite absolute links in an attempt to prevent the user from accidentally clicking a link and leaving the archive.
* **Clean URLs:** Mimics GitHub Pages logic-it automatically checks for `index.html` or `.html` if you leave the extension off the URL.

## Tech Stack

* **Frontend:** `HTML/JS/CSS`
* **Backend:** `Go` with `net/http` and the `Github API`

## Setup

### 1. **Clone the repository**

```bash
git clone https://github.com/keshuook/keshuook-web-archive.git
```

### 2. **Configure:**

Modify `cmd/keshuook-web-archive/main.go` to point to your specific repository (if you wish).

Create a `.env` file in the root directory and add your GitHub token:

```env
GITHUB_AUTH_TOKEN=your_token_here
```

### 3. **Run:**

```bash
go run cmd/keshuook-web-archive/main.go
```

The server runs on `http://localhost:3000`. A demo is available at [https://keshuook-web-archive.onrender.com](https://keshuook-web-archive.onrender.com)
