# Project Greenlight

## Folder Structure

- bin
contains application compiled binaries for deployment on production server 
- cmd/api 
application specific code for api. 
- internal
various packages used by the api, any code which isn't application specific
- migrations 
SQL migration files for database
- remote
configuration files and setup scripts for production server

## Endpoint

| Method | URL Pattern | Handler | Action |
|:----------:|:----------:|:----------:|
| GET | /v1/healthcheck | healthcheckHandler | Show application information 
| GET | /v1/movies | listMoviesHandler | Show the details of all movies
| POST | /v1/movies | createMovieHandler | Create new movie listing
| GET | /v1/movies/:id | showMovieHandler | Show the details of a specific movie
| PUT | /v1/movies/:id | editMovieHandler | Update the details of a specific movie
| DELETE | /v1/movies/:id | deleteMovieHandler | Delete a specefic movie listing



## Local Deployment
```bash
go run ./cmd/api -port=3030 -env=development
```

This project is for educational purpose.
Source: Let's Go Further By Alex Edwards
On page: 87
