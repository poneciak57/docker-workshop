
run:
	docker network create mynet || true
	docker build -t my-python-server .
	docker run -d --rm --name server1 --network mynet my-python-server python server1.py
	docker run -d --rm --name server2 --network mynet my-python-server python server2.py
	docker run -d --rm --name nginx-proxy -p 8080:80 --network mynet -v $(PWD)/nginx.conf:/etc/nginx/nginx.conf:ro nginx:alpine
	@echo "App running on http://localhost:8080/server1/ and http://localhost:8080/server2/"

clean:
	docker stop server1 server2 nginx-proxy || true
	docker network rm mynet || true
	docker rmi my-python-server || true

.PHONY: run clean