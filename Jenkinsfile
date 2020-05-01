pipeline {
  agent any
  environment {
    WORKSPACE = '/home/deploy/go-logrotate'
    }
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
