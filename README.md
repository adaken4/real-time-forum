# Real-Time Forum

A single page real-time forum application with live messaging and user interactions.

## Features

- Real-time messaging
- User authentication
- Live post updates
- Interactive discussions

## Tech Stack

- **Backend**: Go
- **Frontend**: HTML, CSS, JavaScript
- **Database**: SQLite
- **WebSockets**: For real-time communication

## Setup

### Prerequisites

- Go 1.22+
- Modern web browser

### Installation

1. Clone the repository:
   ```bash
   git clone <https://github.com/adaken4/real-time-forum>
   cd real-time-forum
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   make run
   ```

4. Open your browser and navigate to `http://localhost:8080`

## Development

- Backend server runs on port 8080
- Static files served from `frontend/static` directory
- Database file: `rt-forum.db`

## Project Structure

```
.
├── cmd/             # Main application entry point (server/main.go)
├── frontend/        # All static assets (HTML, CSS, JS)
│   ├── index.html   # Main application shell
│   └── static/      # CSS, JS, etc.
├── internal/        # Private application logic (Go standards)
│   ├── app/         # Handlers, Services, and Middleware
│   ├── domain/      # Core business objects/models (Post, User, etc.)
│   ├── infra/       # Infrastructure (DB, Storage, WebSockets Hub)
│   └── ...          # Auth, Config, Realtime components
├── pkg/             # Reusable, external-facing utilities
├── migrations/      # Database migration scripts
├── rt-forum.db      # SQLite database file
└── ...              # Project files (LICENSE, Makefile, README.md)
```
