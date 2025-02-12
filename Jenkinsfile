pipeline {
    agent any

    environment {
        REPO = "zabella/go_workers"
        DOCKER_HUB_CREDENTIALS = credentials('docker-hub-credentials')
        DOCKER_IMAGE = 'golang'
        DOCKER_TAG = 'latest'
    }

    stages {
        stage('Configure credentials') {
            steps {
                withCredentials([sshUserPrivateKey(credentialsId: 'jenkins_ssh_key', keyFileVariable: 'private_key', usernameVariable: 'username')]) {
                    // Add your steps here
                }
            }
        }

        stage('Clone Repository') {
            steps {
                git 'https://github.com/Bella0708/go_workers'
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    docker.build("${env.DOCKER_IMAGE}:${env.DOCKER_TAG}")
                }
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
