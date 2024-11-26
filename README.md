# Game Library  Management System

## Context

You're building a small library management system for video games. The system needs to handle basic operations for managing games and developers while following clean architecture principles using the Repository Pattern.

## == Part 1 ==

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
