pipeline {
    agent any
    environment {
        DOCKERHUB_CREDENTIALS = 'docker-credentials'
        DOCKER_ACCESS_TOKEN = credentials('docker-credentials')
        DOCKER_USERNAME = 'dshwartzman5'
    }
    triggers {
        githubPush()
    }
    stages {
        stage('Pull and Build Docker Image') {
            steps {
                script {
                    echo "Logging into Docker Hub"
                    bat "docker login -u $DOCKER_USERNAME -p $DOCKER_ACCESS_TOKEN"
                    echo "Building Docker Image"
                    bat "docker build -t dshwartzman5/go-jenkins-dockerhub-repo:${env.BUILD_ID} ."
                }
            }
        }

        stage('Run Tests') {
            steps {
                script {
                    // Run the container with the tests
                    bat 'docker run -d -p 8081:8081 --name test-container dshwartzman5/go-jenkins-dockerhub-repo:${env.BUILD_ID}'

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

        stage('Push Docker Image') {
            steps {
                script {
                    bat "docker tag dshwartzman5/go-jenkins-dockerhub-repo:${env.BUILD_ID} dshwartzman5/go-jenkins-dockerhub-repo:${env.BUILD_ID}"
                    bat "docker push dshwartzman5/go-jenkins-dockerhub-repo:${env.BUILD_ID}"
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
