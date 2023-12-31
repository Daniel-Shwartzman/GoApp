pipeline {
 agent any
 environment {
   DOCKER_IMAGE = 'dshwartzman5/go-jenkins-dockerhub-repo:latest'
   DOCKERHUB_CREDENTIALS = 'docker-credentials' // Refer to the credentials added in Jenkins
 }
 stages {
   stage('Pull Docker Image') {
       steps {
           script {
             docker.image("${DOCKER_IMAGE}").pull()
           }
       }
   }
    stage('Run Tests') {
    steps {
    script {
        def dockerArgs = '-p 8081:8081 -v /c/programdata/jenkins/.jenkins/workspace/goapppipeline/:/app'
        docker.image("${DOCKER_IMAGE}").run(dockerArgs) {
            sh 'cd /app && go test'
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
