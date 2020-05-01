pipeline {
  agent any
  stages {
    stage('go build') {
      steps {
        sh 'go build'
      }
    }
    stage('go go-logrotate') {
      steps {
        sh 'go go-logrotate'
      }
    }
  }
}
