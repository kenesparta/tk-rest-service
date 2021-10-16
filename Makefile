l/build:
	ls

l/up:
	docker-compose down --remove-orphans --rmi all
	docker-compose up --detach --remove-orphans --force-recreate