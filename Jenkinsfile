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

				def varsFile = "var_file.txt"
				def content = readFile varsFile
				def lines = content.split("\n")
				for(l in lines){
        			String variable = "${l.split(" ")[1].split("=")[0]}"
        			String value = l.split(" ")[1].split("=")[1]
					sh ("echo env.$variable = \\\"$value\\\" >> var_to_exp.groovy") 
 				}
				load var_to_exp.groovy

				sh 'rm -rf var_to_exp.groovy'

				sh 'printenv'
				sh 'go test -v test/handlers_test.go'
			}
    	}
	}
}
