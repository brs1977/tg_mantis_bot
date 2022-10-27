build:
	docker build -t mantis_bot:v1 .
run:
	# docker run -it --env-file .env.prod --restart unless-stopped mantis_bot:v1
	docker run -d --restart unless-stopped --env-file .env.prod mantis_bot:v1
shell:
	docker run -it --env-file .env.prod --rm mantis_bot:v1 sh
test_curl:
	curl -d "from=2022-10-06&to=2022-10-13&report=%CF%EE%EA%E0%E7%E0%F2%FC" -X POST http://129.200.0.18/mantis/mantisreport.php