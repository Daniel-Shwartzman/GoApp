pipeline {
    agent any
    environment {
        DOCKERHUB_CREDENTIALS = 'docker-credentials'
        DOCKER_ACCESS_TOKEN = credentials('docker-credentials')
        DOCKER_USERNAME = 'dshwartzman5'
        AWS_ACCESS_KEY_ID = credentials('aws-access-key-id')
        AWS_SECRET_ACCESS_KEY = credentials('aws-secret-access-key')
    }
    triggers {
        githubPush(branch: 'main')
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

        stage('Push Docker Image') {
            steps {
                script {
                    bat "docker tag dshwartzman5/go-jenkins-dockerhub-repo:latest dshwartzman5/go-jenkins-dockerhub-repo:latest"
                    bat "docker push dshwartzman5/go-jenkins-dockerhub-repo:latest"
                }
            }
        }

        stage('Terraform Apply') {
            steps {
                script {
                    dir('terraform') {
                        bat 'terraform init'
                        bat 'terraform apply -auto-approve'
                    }
                }
            }
        }

        stage('Deploy Container') {
            steps {
                script {
                    // SSH into the EC2 instance and deploy the Docker container
                    withCredentials([string(credentialsId: 'ec2-ssh-key', variable: 'SSH_KEY')]) {
                        dir('terraform') {
                            // Read the instance public IP from Terraform output
                            def instanceIP = bat(script: 'terraform output -raw instance_public_ip', returnStatus: true).trim()

                            // Use the retrieved IP in the SSH command
                            bat "ssh -o StrictHostKeyChecking=no -i $SSH_KEY ec2-user@${instanceIP} 'docker pull dshwartzman5/go-jenkins-dockerhub-repo:latest && docker run -d -p 8081:8081 dshwartzman5/go-jenkins-dockerhub-repo:latest'"
                        }
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

    post {
        always {
            withCredentials([string(credentialsId: 'discord-credential', variable: 'WEBHOOK_URL')]) {
                script {
                    def buildStatus = currentBuild.currentResult
                    def buildStatusMessage = buildStatus == 'SUCCESS' ? 'Build Succeeded' : 'Build Failed'
                    discordSend description: buildStatusMessage, link: env.BUILD_URL, result: buildStatus, title: JOB_NAME, webhookURL: WEBHOOK_URL
                }
            }
        }
    }
}
