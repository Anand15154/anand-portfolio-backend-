// MongoDB initialization script
// This script runs when the MongoDB container starts for the first time

// Switch to the portfolio database
db = db.getSiblingDB('portfolio_db');

// Create collections with validation
db.createCollection('contacts', {
    validator: {
        $jsonSchema: {
            bsonType: "object",
            required: ["name", "email", "subject", "message"],
            properties: {
                name: {
                    bsonType: "string",
                    description: "must be a string and is required"
                },
                email: {
                    bsonType: "string",
                    pattern: "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$",
                    description: "must be a valid email address and is required"
                },
                subject: {
                    bsonType: "string",
                    description: "must be a string and is required"
                },
                message: {
                    bsonType: "string",
                    description: "must be a string and is required"
                },
                created_at: {
                    bsonType: "date",
                    description: "must be a date"
                },
                read: {
                    bsonType: "bool",
                    description: "must be a boolean"
                }
            }
        }
    }
});

db.createCollection('projects', {
    validator: {
        $jsonSchema: {
            bsonType: "object",
            required: ["title", "description"],
            properties: {
                title: {
                    bsonType: "string",
                    description: "must be a string and is required"
                },
                description: {
                    bsonType: "string",
                    description: "must be a string and is required"
                },
                image_url: {
                    bsonType: "string",
                    description: "must be a string"
                },
                live_url: {
                    bsonType: "string",
                    description: "must be a string"
                },
                github_url: {
                    bsonType: "string",
                    description: "must be a string"
                },
                technologies: {
                    bsonType: "array",
                    items: {
                        bsonType: "string"
                    },
                    description: "must be an array of strings"
                },
                category: {
                    bsonType: "string",
                    description: "must be a string"
                },
                featured: {
                    bsonType: "bool",
                    description: "must be a boolean"
                },
                created_at: {
                    bsonType: "date",
                    description: "must be a date"
                },
                updated_at: {
                    bsonType: "date",
                    description: "must be a date"
                }
            }
        }
    }
});

// Create indexes for better performance
db.contacts.createIndex({ "created_at": -1 });
db.contacts.createIndex({ "read": 1 });
db.contacts.createIndex({ "email": 1 });

db.projects.createIndex({ "created_at": -1 });
db.projects.createIndex({ "featured": 1 });
db.projects.createIndex({ "category": 1 });

// Insert sample projects
db.projects.insertMany([
    {
        title: "E-Commerce Platform",
        description: "A full-stack e-commerce platform built with React, Node.js, and MongoDB. Features include user authentication, product management, shopping cart, and payment integration.",
        image_url: "https://via.placeholder.com/600x400/2563eb/ffffff?text=E-Commerce",
        live_url: "https://example-ecommerce.com",
        github_url: "https://github.com/username/ecommerce-platform",
        technologies: ["React", "Node.js", "MongoDB", "Express", "Stripe"],
        category: "Web Development",
        featured: true,
        created_at: new Date(),
        updated_at: new Date()
    },
    {
        title: "Task Management App",
        description: "A collaborative task management application with real-time updates, drag-and-drop functionality, and team collaboration features.",
        image_url: "https://via.placeholder.com/600x400/7c3aed/ffffff?text=Task+Manager",
        live_url: "https://example-taskmanager.com",
        github_url: "https://github.com/username/task-manager",
        technologies: ["Vue.js", "Firebase", "Vuex", "Vuetify"],
        category: "Web Development",
        featured: true,
        created_at: new Date(),
        updated_at: new Date()
    },
    {
        title: "Weather Dashboard",
        description: "A weather dashboard that displays current weather conditions and forecasts using data from multiple weather APIs.",
        image_url: "https://via.placeholder.com/600x400/059669/ffffff?text=Weather+App",
        live_url: "https://example-weather.com",
        github_url: "https://github.com/username/weather-dashboard",
        technologies: ["JavaScript", "HTML5", "CSS3", "OpenWeather API"],
        category: "Web Development",
        featured: false,
        created_at: new Date(),
        updated_at: new Date()
    },
    {
        title: "Portfolio Website",
        description: "A responsive portfolio website showcasing projects and skills with modern design and smooth animations.",
        image_url: "https://via.placeholder.com/600x400/dc2626/ffffff?text=Portfolio",
        live_url: "https://example-portfolio.com",
        github_url: "https://github.com/username/portfolio",
        technologies: ["React", "Tailwind CSS", "Framer Motion", "Vite"],
        category: "Web Development",
        featured: true,
        created_at: new Date(),
        updated_at: new Date()
    }
]);

print("Database initialization completed successfully!");
print("Sample projects have been inserted.");
print("Collections created: contacts, projects");
print("Indexes created for better performance."); 