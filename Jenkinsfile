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
    stage('go go-logrotate') {
      steps {
        sh './go-logrotate_master'
      }
    }
  }
}
