pipeline {
  agent any
  environment {
    WORKSPACE = '/home/deploy/jenkins'
    JENKINS_HOME = '/home/deploy/jenkins'
    }
  stages {
    stage('Change working directory...') {
      steps {
        dir('/home/deploy/jenkins/project') {
          sh 'pwd'
        }
      }
    }
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
