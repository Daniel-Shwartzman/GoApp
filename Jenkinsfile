pipeline {
  agent any
  environment {
      DOCKER_IMAGE = 'dshwartzman5/go-jenkins-dockerhub-repo'
      DOCKERHUB_CREDENTIALS = 'docker-credentials' // Refer to the credentials added in Jenkins
  }
  stages {
      stage('Pull Docker Image') {
          steps {
              script {
                docker.image("${DOCKER_IMAGE}:latest").pull()
              }
          }
      }
        stage('Run Tests') {
            steps {
            script {
                docker.image("dshwartzman5/go-jenkins-dockerhub-repo:latest").inside('.') {
                    sh 'go test'
                }
            }
            }
            }


            stage('Tag Docker Image') {
          steps {
              script {
                docker.image("${DOCKER_IMAGE}:latest").tag("${DOCKER_IMAGE}:${env.BUILD_ID}")
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
