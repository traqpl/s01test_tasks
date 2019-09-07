all: gobuild maven dockerbuild deploy

gobuild: 
	go build -o kafka_check nagios_kafka_topic_check.go

maven:
	git clone https://github.com/TechPrimers/docker-mysql-spring-boot-example.git
	cd docker-mysql-spring-boot-example && mvn package 

dockerbuild:
	cd docker-mysql-spring-boot-example && doker build . -t users-mysql
    docker tag users-mysql traqtest/s01:latest
    docker push traqtest/s01:latest

deploy:
	gcloud container clusters get-credentials s01 --zone europe-west1-b --project dev-242510
	kubectl create secret generic mysqlrootpassword --from-literal=password="password"
	kubectl create secret generic mysqluserpassword --from-literal=password="password"
	kubectl create namespace springboot
	kubect create -f k8s/dev/mysql-pv.yaml
	kubect create -f k8s/dev/mysql.yaml
	kubect create -f k8s/dev/springboot.yaml
