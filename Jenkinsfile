def remote = [:]
pipeline {
    agent any
    parameters {
        gitParameter name: 'branch', type: 'PT_BRANCH', sortMode: 'DESCENDING_SMART', selectedValue: 'NONE', quickFilterEnabled: true
    } 
    environment {
        REPO = "zabella/go_workers"
        DOCKER_IMAGE = 'golang'
        DOCKER_TAG = 'latest'
        HOST = "3.94.85.77"
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
                git (url: 'https://github.com/Bella0708/go_workers', branch: '${branch}')
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
