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
                def testContainer = docker.image("${DOCKER_IMAGE}:latest").run("-d")
                // Here you would run your tests against the running container.
                // Replace 'your-test-command' with the actual command to run your tests.
                def testResult = testContainer.exec(['sh', '-c', 'your-test-command'])
                echo "${testResult}"
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
