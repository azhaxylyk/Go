# Ascii-art-web-dockerize
## Description

The "ASCII Art Web Dockerize" project is a web server developed in the Go programming language that enables users to convert text into ASCII art. The project is containerized using Docker for ease of deployment and dependency isolation.

Key Features:

A Go-based web server adhering to programming best practices.
Includes a Dockerfile, image, and container for the project.
Utilizes metadata for Docker objects and manages unused objects efficiently.

###  Usage: How to Run

To run the Ascii-art-web-dockerize server, follow these steps:

    Clone the repository: git clone git@git.01.alem.school:ayerkebul/ascii-art-web-dockerize.git


Navigate to the project directory:

    cd ascii-art-web-dockerize
    cd cmd

Build the Docker image from the Dockerfile:

    docker build -t ascii-art-web-dockerize .

Run the container from the built image:

    docker run -p 8080:8080 ascii-art-web-dockerize

    Open a web browser and go to http://localhost:8080 to access the application.

### Docker objects 
    Dockerfile: Defines the instructions to build the project's Docker image.
    Image: Contains everything necessary to run the web server, including the executable file and static files.
    Container: A running instance of the image that isolates the application and its dependencies from the rest of the system.

#### Cleaning up unused docker objects. To remove unused Docker images, containers, and networks, execute the following command: 
    docker system prune


### Authors
    ayerkebul, azhaxylyk