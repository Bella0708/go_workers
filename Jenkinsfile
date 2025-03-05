def remote = [:]
def git_url = "git@github.com:Bella0708/go_workers.git"
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
        TOKEN = credentials('tokentell')
        CHAT_ID = "683028341"
        LINK = "<a href=\\\"${BUILD_URL}\\\">${JOB_NAME} #${BUILD_NUMBER}</a>"
    }

    stages {
        stage('Configure credentials') {
            steps {
                withCredentials([sshUserPrivateKey(credentialsId: 'jenkins_key', keyFileVariable: 'private_key', usernameVariable: 'username')]) {
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
                
               checkout([$class: 'GitSCM', branches: [[name: "${branch}"]], doGenerateSubmoduleConfigurations: false, extensions: [], submoduleCfg: [], userRemoteConfigs: [[credentialsId: 'jenkins_key', url: "$git_url"]]])
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
    post {
      success {
        script {
          sh "curl -X POST -H 'Content-Type: application/json' -d '{\"chat_id\": \"${CHAT_ID}\", \"text\": \"${LINK}\n🟢 Deploy succeeded! \", \"parse_mode\": \"HTML\", \"disable_notification\": false}' \"https://api.telegram.org/bot${TOKEN}/sendMessage\""
        }
      }
      failure {
        script {
          sh "curl -X POST -H 'Content-Type: application/json' -d '{\"chat_id\": \"${CHAT_ID}\", \"text\": \"${LINK}\n🔴 Deploy failure! \", \"parse_mode\": \"HTML\", \"disable_notification\": false}' \"https://api.telegram.org/bot${TOKEN}/sendMessage\""
        }
      }
      aborted {
        script {
          sh "curl -X POST -H 'Content-Type: application/json' -d '{\"chat_id\": \"${CHAT_ID}\", \"text\": \"${LINK}\n⚪️ Deploy aborted! \", \"parse_mode\": \"HTML\", \"disable_notification\": false}' \"https://api.telegram.org/bot${TOKEN}/sendMessage\""
        }
      }    
    }
}
