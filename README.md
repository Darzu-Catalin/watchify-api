# Movie Recommendation System API

This repository contains the API for a Movie Recommendation System built using Go. The API provides endpoints to retrieve movie information, manage user data, and recommend movies based on different criteria.

## Table of Contents
- [About](#about)
- [Features](#features)
- [Technologies Used](#technologies-used)

## About

This project aims to provide a movie recommendation service, which retrieves movie details (like genre, cast, crew, etc.) and offers personalized recommendations based on user preferences and interactions.

## Features

- **Retrieve Movies**: Fetch a list of movies from the database.
- **Movie Details**: Fetch detailed information about a specific movie.
- **User Management**: Create and manage users who can interact with the movie database.
- **Movie Recommendations**: Personalized recommendations based on user data.

## Technologies Used

- **Go (Golang)**: The core language for building the API.
- **MySQL**: Database to store movie and user information.
- **Gin**: A web framework for Go that is used to build the HTTP API.
- **Gorm**: ORM for interacting with MySQL.
