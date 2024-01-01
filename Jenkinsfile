pipeline {
    agent any
    environment {
        DOCKER_IMAGE = 'dshwartzman5/go-jenkins-dockerhub-repo:latest'
        DOCKERHUB_CREDENTIALS = 'docker-credentials'
        DOCKER_ACCESS_TOKEN = credentials('docker-credentials')
        DOCKER_USERNAME = 'dshwartzman5'
    }
    stages {
        stage('Pull Docker Image') {
            steps {
                script {
                    echo "Logging into Docker Hub"
                    bat "docker login -u $DOCKER_USERNAME -p $DOCKER_ACCESS_TOKEN"
                    echo "Pulling Docker Image"
                    bat "docker pull ${DOCKER_IMAGE}"
                }
            }
        }

        stage('Run Tests') {
            steps {
                script {
                    // Use the Docker image for testing
                    def dockerImage = "${DOCKER_IMAGE}"

                    // Run tests within the Docker container
                    docker.image(dockerImage).inside {
                        // Tests are already executed in the container
                    }

                    // Read the test results file and check if the tests passed
                    def testResults = readFile('/app/test_results.txt')
                    if (testResults =~ /FAIL/) {
                        error('Tests failed')
                    }
                }
            }
        }

        stage('Tag Docker Image') {
            steps {
                script {
                    docker.image("${DOCKER_IMAGE}").tag("${DOCKER_IMAGE}:${env.BUILD_ID}")
                }
            }
        }

        stage('Push Docker Image') {
            steps {
                script {
                    docker.withRegistry('https://registry.hub.docker.com', "${DOCKERHUB_CREDENTIALS}") {
                        docker.image("${DOCKER_IMAGE}:${env.BUILD_ID}").push()
                    }
                }
            }
        }

        stage('Cleanup') {
            steps {
                cleanWs()
            }
        }
    }
}
