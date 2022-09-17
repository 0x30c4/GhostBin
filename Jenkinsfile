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
				sh 'mkdir test/testdata'
				script {
					def props = readProperties  file: ".env"
    				keys= props.keySet()
    				for(key in keys) {
        				value = props["${key}"]
        				env."${key}" = "${value}"
    				}

				}
				sh 'ls -la test'
				sh 'go test -v test/handlers_test.go'
				sh 'rm -rf test/testdata'
			}
    	}
	}
}
