pipeline {
    agent any

    environment {
        REPO = "zabella/go_workers"
        DOCKER_HUB_CREDENTIALS = credentials('docker-hub-credentials')
        DOCKER_IMAGE = 'golang'
        DOCKER_TAG = 'latest'
        HOST = "51.250.110.75"
    }

    stages {
        stage('Configure credentials') {
            steps {
                withCredentials([sshUserPrivateKey(credentialsId: 'jenkins_ssh_key', keyFileVariable: 'private_key', usernameVariable: 'username')]) {
                    script {
            remote.name = "${env.HOST}"
            remote.host = "${env.HOST}"
            remote.user = "$username"
            remote.identity = readFile("$private_key")
            remote.allowAnyHosts = true
          }
                }
            }
        }

        stage('Clone Repository') {
            steps {
                git 'https://github.com/Bella0708/go_workers'
            }
        }

       

        stage('Build and Push Docker Image') {
            steps {
                script {
                    def image = docker.build("${env.REPO}:${env.BUILD_ID}")
                    docker.withRegistry('https://registry-1.docker.io', 'hub_token') {
                        image.push()
                    }
                }
            }
        }
    }
}
