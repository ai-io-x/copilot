**Build and run the Docker container** from the `build-ops` directory:

    ```sh
    cd build-ops
    docker-compose up --build
    ```

**Multi-stage Docker build**
You can use a multi-stage Docker build to first build and test the application in one stage, and then copy the JAR file to another stage for deployment. Finally, you can push the built image to GitHub Container Registry.


To build, test, and deploy the application, follow these steps:

1. **Build and test the Docker image**:

    ```sh
    cd build-ops
    docker-compose build
    ```

2. **Run the application**:

    ```sh
    docker-compose up
    ```

3. **Push the Docker image to GitHub Container Registry**:

    First, log in to GitHub Container Registry:

    ```sh
    echo $GITHUB_TOKEN | docker login ghcr.io -u USERNAME --password-stdin
    ```

    Replace `USERNAME` with your GitHub username and ensure you have a GitHub token with the appropriate permissions.

    Tag the Docker image:

    ```sh
    docker tag build-ops_app:latest ghcr.io/USERNAME/REPOSITORY:TAG
    ```

    Replace `USERNAME` with your GitHub username, `REPOSITORY` with your repository name, and `TAG` with the desired tag.

    Push the Docker image:

    ```sh
    docker push ghcr.io/USERNAME/REPOSITORY:TAG
    ```
