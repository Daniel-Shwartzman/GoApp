pipeline {
    agent any
    environment {
        DOCKER_IMAGE = 'dshwartzman5/go-jenkins-dockerhub-repo:latest'
        DOCKERHUB_CREDENTIALS = 'docker-credentials' // Refer to the credentials added in Jenkins
        DOCKER_ACCESS_TOKEN = credentials('docker-credentials') // Assuming your credential ID is 'docker-credentials'
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
                    def dockerArgs = '-p 8081:8081 -v C:\\Users\\Daniel Schwartzman\\Desktop\\myproject:/app'
                    docker.image("${DOCKER_IMAGE}").withRun(dockerArgs) {
                        // Read the test_results.txt file
                        def testResults = readFile('/app/test_results.txt')
                        // Check if the tests passed
                        if (testResults =~ /FAIL/) {
                            error('Tests failed')
                        }
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
