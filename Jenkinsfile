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
        script {
        	load "./pipeline/basics/extenvvariable/env.groovy"
            echo "${env.env_var1}"
            echo "${env.env_var2}"
        }
		sh 'env'
        sh './scripts/run-test.sh'
      }
    }

  }
}
