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
                    // Run the container with the tests
                    bat 'docker run -d -p 8081:8081 --name test-container dshwartzman5/go-jenkins-dockerhub-repo:latest'

                    // Copy the test results from the container to the workspace
                    bat 'docker cp test-container:/app/test_results.txt .'

                    // Stop and remove the container
                    bat 'docker stop test-container'
                    bat 'docker rm test-container'

                    // Read the test results
                    def testResults = readFile('test_results.txt')

                    // Check if the tests passed
                    if (testResults.contains('FAIL')) {
                        error 'Tests failed'
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
