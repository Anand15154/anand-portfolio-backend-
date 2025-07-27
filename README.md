# Portfolio Backend

A Go-based REST API backend for a portfolio website built with Gin framework and MongoDB Atlas.

## Features

- **Contact Form Management**: Handle contact form submissions with email notifications
- **Project Management**: CRUD operations for portfolio projects
- **Authentication**: JWT-based authentication for admin routes
- **Rate Limiting**: Prevent spam and abuse
- **Email Notifications**: SMTP integration for contact form notifications
- **CORS Support**: Cross-origin resource sharing configuration
- **Logging**: Request logging and error tracking
- **Cloud Database**: MongoDB Atlas integration

## Tech Stack

- **Framework**: Gin (Go web framework)
- **Database**: MongoDB Atlas (Cloud)
- **Authentication**: JWT
- **Email**: SMTP
- **Configuration**: Environment variables with .env support

## Prerequisites

- Go 1.21 or higher
- MongoDB Atlas account (already configured)
- SMTP credentials (optional, for email notifications)
- Node.js (for database initialization)

## Installation

1. Clone the repository and navigate to the backend directory:
```bash
cd backend
```

2. Install Go dependencies:
```bash
go mod tidy
```

3. Copy the environment file and configure it:
```bash
cp env.example .env
```

4. The `.env` file is already configured with your MongoDB Atlas connection:
```env
# Server Configuration
PORT=8080
GIN_MODE=debug

# MongoDB Configuration (MongoDB Atlas)
MONGODB_URI=mongodb+srv://anandtiwari3399:Anand.01@anand-portfolio.jyac0ao.mongodb.net/
MONGODB_DATABASE=portfolio_db

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key-here
JWT_EXPIRY=24h

# CORS Configuration
ALLOWED_ORIGINS=http://localhost:5173,http://localhost:3000

# Email Configuration (optional)
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your-email@gmail.com
SMTP_PASSWORD=your-app-password
```

5. Initialize the MongoDB Atlas database:
```bash
make init-atlas
```

## Running the Application

### Development
```bash
go run cmd/main.go
```

### Production
```bash
go build -o portfolio-backend cmd/main.go
./portfolio-backend
```

### Docker (Recommended)
```bash
# Build and run with Docker Compose
docker-compose up -d
```

The server will start on `http://localhost:8080` (or the port specified in your .env file).

## API Endpoints

### Health Check
- `GET /health` - Check server status

### Authentication
- `POST /api/v1/auth/login` - Admin login
  ```json
  {
    "username": "admin",
    "password": "password"
  }
  ```

### Contact Management
- `POST /api/v1/contacts/` - Submit contact form (rate limited)
  ```json
  {
    "name": "John Doe",
    "email": "john@example.com",
    "subject": "Hello",
    "message": "Your message here"
  }
  ```
- `GET /api/v1/contacts/` - Get all contacts (admin only)
- `GET /api/v1/contacts/:id` - Get specific contact (admin only)
- `PUT /api/v1/contacts/:id/read` - Mark contact as read (admin only)
- `DELETE /api/v1/contacts/:id` - Delete contact (admin only)

### Project Management
- `POST /api/v1/projects/` - Create new project (admin only)
  ```json
  {
    "title": "Project Title",
    "description": "Project description",
    "image_url": "https://example.com/image.jpg",
    "live_url": "https://example.com",
    "github_url": "https://github.com/user/repo",
    "technologies": ["React", "Node.js", "MongoDB"],
    "category": "Web Development",
    "featured": true
  }
  ```
- `GET /api/v1/projects/` - Get all projects
- `GET /api/v1/projects/featured` - Get featured projects
- `GET /api/v1/projects/:id` - Get specific project
- `PUT /api/v1/projects/:id` - Update project (admin only)
- `DELETE /api/v1/projects/:id` - Delete project (admin only)

## Authentication

For admin routes, include the JWT token in the Authorization header:
```
Authorization: Bearer <your-jwt-token>
```

Default admin credentials:
- Username: `admin`
- Password: `password`

**Important**: Change these credentials in production!

## Rate Limiting

Contact form submissions are rate-limited to 10 requests per minute per IP address.

## Email Configuration

To enable email notifications for contact form submissions:

1. For Gmail, enable 2-factor authentication and generate an app password
2. Update the SMTP configuration in your `.env` file
3. The email service will automatically send notifications when contact forms are submitted

## Database Setup

### MongoDB Atlas Configuration

The application is configured to use MongoDB Atlas with the following connection string:
```
mongodb+srv://anandtiwari3399:Anand.01@anand-portfolio.jyac0ao.mongodb.net/
```

### Database Initialization

To initialize your MongoDB Atlas database with sample data:

```bash
# Using Makefile (recommended)
make init-atlas

# Or manually
npm install
npm run init-db
```

This will:
- Create the required collections (`contacts`, `projects`)
- Set up database validation rules
- Create indexes for better performance
- Insert sample projects

## Project Structure

```
backend/
├── cmd/
│   └── main.go              # Application entry point
├── configs/
│   └── config.go            # Configuration management
├── internal/
│   ├── database/
│   │   └── mongodb.go       # Database connection and utilities
│   ├── handlers/
│   │   ├── auth_handler.go  # Authentication handlers
│   │   ├── contact_handler.go # Contact form handlers
│   │   └── project_handler.go # Project management handlers
│   ├── middleware/
│   │   ├── auth.go          # JWT authentication middleware
│   │   ├── cors.go          # CORS middleware
│   │   ├── logging.go       # Request logging middleware
│   │   └── rate_limit.go    # Rate limiting middleware
│   ├── models/
│   │   ├── contact.go       # Contact data models
│   │   └── project.go       # Project data models
│   └── services/
│       ├── contact_service.go # Contact business logic
│       ├── email_service.go   # Email service
│       └── project_service.go # Project business logic
├── env.example              # Environment variables template
├── go.mod                   # Go module file
├── package.json             # Node.js dependencies for DB init
├── init-atlas-db.js         # MongoDB Atlas initialization script
├── Dockerfile               # Container configuration
├── docker-compose.yml       # Development environment
├── Makefile                 # Development tasks
└── README.md               # This file
```

## Development

### Adding New Features

1. Create models in `internal/models/`
2. Implement business logic in `internal/services/`
3. Create handlers in `internal/handlers/`
4. Add routes in `cmd/main.go`
5. Add middleware if needed in `internal/middleware/`

### Testing

Run tests:
```bash
go test ./...
```

### Database Operations

The application will automatically connect to MongoDB Atlas and create the necessary collections when it first runs.

## Deployment

### Docker (Recommended)

The application is already configured for Docker deployment:

```bash
# Build and run
docker-compose up -d

# Or build manually
docker build -t portfolio-backend .
docker run -p 8080:8080 --env-file .env portfolio-backend
```

### Environment Variables for Production

Make sure to set these environment variables in production:
- `GIN_MODE=release`
- `JWT_SECRET` (use a strong, random secret)
- `MONGODB_URI` (your MongoDB Atlas connection string)
- `ALLOWED_ORIGINS` (your frontend domain)

## Security Considerations

1. **JWT Secret**: Use a strong, random secret in production
2. **CORS**: Configure allowed origins properly
3. **Rate Limiting**: Adjust limits based on your needs
4. **Input Validation**: All inputs are validated using Gin's binding
5. **Database**: MongoDB Atlas provides built-in security features
6. **HTTPS**: Always use HTTPS in production

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License. # anand-portfolio-backend-
