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
					withEnv(readFile('.env').split('\n') as List) {
						sh "echo ${version}"
					}
				}
				sh 'env'
				sh 'go test -v test/handlers_test.go'
			}
    	}
	}
}
