pipeline {
  agent any
  stages {
    stage('go build') {
      steps {
        sh 'go get -d ./'
        sh 'go build'
      }
    }
    stage('go go-logrotate') {
      steps {
        sh './go-logrotate_master'
      }
    }
  }
}
