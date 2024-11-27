# The task to do is:

# Game Library  Management System

## Context

You're building a small library management system for video games. The system needs to handle basic operations for managing games and developers while following clean architecture principles using the Repository Pattern.

## Requirements

- Use NoSQL database

### Developer Properties

- Name
- Main HQ Location

### Game Properties

- Title
- Developer
- Genre
- Publication Year
- Available (boolean)

### Required Operations

For Games:

1. Get all games
2. Get game by ID
3. Add a new game
4. Update a game’s availability
5. Delete a game
6. Find games by developer

For Developers:

1. Get all developers
2. Get developer by ID
3. Add a new developer
4. Update developer information
5. Delete developer (should handle associated games)

### Architecture Requirements

- Use the **Repository Pattern** to abstract the data access layer


# Steps to Set up the Project

## Requirments

To set up the project you need **Docker** and **Docker Compose** installed.

## Follow these steps to get your project running using Docker Compose.

## 1. Clone the Project

Clone the project repository to your local machine:

```bash
git clone https://github.com/Dosik13/game-library-management-system.git
```

## 2. Set up the .env file

Set up your .env file (There is a .env.example file that shows the information needed)

## 3. Build the Docker Images

```bash
docker-compose build
```

## 4. Start the Containers

```bash
docker-compose up
```
