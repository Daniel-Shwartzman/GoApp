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
        stage('Build Docker Image') {
            steps {
                script {
                    echo "Logging into Docker Hub"
                    bat "docker login -u $DOCKER_USERNAME -p $DOCKER_ACCESS_TOKEN"
                    echo "Building Docker Image"
                    bat "docker build -t dshwartzman5/go-jenkins-dockerhub-repo:latest ."
                }
            }
        }

        stage('Run Tests') {
            steps {
                script {
                    // Your existing test stage steps here
                }
            }
        }

        stage('Push Docker Image') {
            steps {
                script {
                    bat "docker tag dshwartzman5/go-jenkins-dockerhub-repo:latest dshwartzman5/go-jenkins-dockerhub-repo:latest"
                    bat "docker push dshwartzman5/go-jenkins-dockerhub-repo:latest"
                }
            }
        }
    }

    post {
    success {
        script {
            emailext subject: 'Pipeline Successful',
                        body: 'The Jenkins pipeline has completed successfully.',
                        recipientProviders: [culprits(), developers()],
                        to: 'dshwartzman5@gmail.com'
        }
    }

    failure {
        script {
            emailext subject: 'Pipeline Failed',
                        body: 'The Jenkins pipeline has failed. Please review the build logs for details.',
                        recipientProviders: [culprits(), developers()],
                        to: 'dshwartzman5@gmail.com'
        }
    }

    always {
        script {
            cleanWs()
        }
    }
}

