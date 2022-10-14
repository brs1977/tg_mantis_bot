build:
	docker build -t mantis_bot:v1 .
run:
	docker run -it --env-file .env.prod --rm mantis_bot:v1
shell:
	docker run -it --env-file .env.prod --rm mantis_bot:v1 sh