# filepath: build-ops/Dockerfile
FROM maven:3.9.9-eclipse-temurin-23 AS build

# Set the working directory
WORKDIR /app

# Copy the pom.xml and source code
COPY pom.xml .
COPY src ./src

# Package the application
RUN mvn clean package

# List the contents of the /app/target directory to verify the JAR file
# RUN ls -l /app/target

# Use a smaller base image for the runtime
FROM openjdk:24-jdk-slim

# Set the working directory
WORKDIR /app-deploy

# Copy the built jar from the build stage
COPY --from=build /app/target/*.jar app.jar

# Expose the application port
EXPOSE 8080

# Run the application
ENTRYPOINT ["java", "-jar", "/app-deploy/app.jar"]
