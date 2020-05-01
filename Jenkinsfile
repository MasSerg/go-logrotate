pipeline {
  agent any
  stages {
    stage('go get') {
      steps {
        sh 'go get -d ./'
      }
    }
    stage('go build') {
      steps {
        sh 'go build'
      }
    }
    stage('env') {
      steps {
        sh 'cp ./.env-dist ./.env'
      }
    }
    stage('go go-logrotate') {
      steps {
        sh './go-logrotate_master'
      }
    }
  }
}
