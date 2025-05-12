# Game Log Hub

A simple game log management system built with Go and Gin framework. The system currently focuses on collecting and displaying login errors.

## Features

- Collection of login errors from games
- Web interface for viewing and managing login error logs
- SQLite database for data storage
- RESTful API for integration with game clients

## Tech Stack

- **Backend**: Go with Gin framework
- **Database**: SQLite
- **Frontend**: HTML, CSS, JavaScript with Bootstrap

## Project Structure

```
game_log_hub/
├── api/            # API routes and handlers
├── config/         # Configuration files
├── controllers/    # Request handlers
├── database/       # Database connection and models
├── middleware/     # HTTP middleware
├── models/         # Data models
├── public/         # Static files for frontend
├── utils/          # Utility functions
├── main.go         # Application entry point
└── go.mod          # Go module definition
```

## API Endpoints

- `GET /api/login-errors` - Get a paginated list of login errors
- `GET /api/login-errors/:id` - Get a specific login error by ID
- `POST /api/login-errors/` - Create a new login error record (note the trailing slash)
- `DELETE /api/login-errors/:id` - Delete a login error record

## Running the Application

1. Clone the repository
2. Build the application:
   ```bash
   go build -o game_log_hub
   ```
3. Run the application:
   ```bash
   ./game_log_hub
   ```
4. Access the web interface at http://localhost:8080

## Setting up as a System Service

To set up Game Log Hub as a system service that starts automatically on boot:

1. Build the application:
   ```bash
   go build -o game_log_hub
   ```

2. Create a systemd service file:
   ```bash
   sudo cp game_log_hub.service /etc/systemd/system/
   sudo systemctl daemon-reload
   ```

3. Enable and start the service:
   ```bash
   sudo systemctl enable game_log_hub.service
   sudo systemctl start game_log_hub.service
   ```

4. Check service status:
   ```bash
   sudo systemctl status game_log_hub.service
   ```

5. Service management commands:
   ```bash
   # Stop the service
   sudo systemctl stop game_log_hub.service
   
   # Restart the service
   sudo systemctl restart game_log_hub.service
   
   # View logs
   sudo journalctl -u game_log_hub.service
   ```

## Sample API Request

Creating a login error record:

```bash
curl -X POST http://192.168.1.200:8080/api/login-errors/ -H "Content-Type: application/json" -d '{
  "user_id": "12345",
  "user_name": "player123",
  "error_type": "authentication_failed",
  "error_msg": "Invalid username or password",
  "ip": "192.168.1.100",
  "platform": "Windows",
  "device": "Desktop"
}'
``` 