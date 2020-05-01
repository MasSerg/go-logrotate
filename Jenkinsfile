pipeline {
  agent any
  stages {
    stage('composer install') {
      steps {
        sh 'composer install'
      }
    }
    stage('docker build & up') {
      steps {
        sh 'docker-compose build'
        sh 'docker-compose up -d'
      }
    }
  }
}
