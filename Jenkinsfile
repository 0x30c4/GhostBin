pipeline {
  agent any
  stages {
    stage('Checkout Code') {
      steps {
        git(url: 'https://github.com/0x30c4/GhostBin', branch: 'main')
      }
    }

    stage('Test') {
      steps {
        sh './scripts/run-test.sh'
      }
    }

  }
}