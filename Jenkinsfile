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
		sh 'cp env-example .env'
        script {
			loadEnvironmentVariablesFromFile(".env")
        }
		sh 'env'
        sh 'go test -v test/handlers_test.go'
      }
    }

  }
}
