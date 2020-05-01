pipeline {
  agent any
  environment {
    WORKSPACE = '/home/deploy/jenkins'
    JENKINS_HOME = '/home/deploy/jenkins'
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
