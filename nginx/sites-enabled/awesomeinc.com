limit_req_zone $binary_remote_addr zone=ip:10m rate=5r/s;
server {	

        root /home/ubuntu/awesome_inc/backend/main;
	index index.html index.htm index.nginx-debian.html;

        server_name awesomeinc.com www.awesomeinc.com;

        location / {
                try_files $uri $uri/ =404;
        }

	location /new_signup {
		limit_req zone=ip burst=8 delay=4;
                proxy_set_header Host $host;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_pass https://localhost:8081;
        }

	location /contact {
		limit_req zone=ip burst=8 delay=4;
                proxy_set_header Host $host;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_pass https://localhost:8081;
        }

	location /test {
		limit_req zone=ip burst=12 delay=8;
                proxy_set_header Host $host;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_pass https://localhost:8081;
        }

        location /device_alert {
		limit_req zone=ip burst=12 delay=8;
                proxy_set_header Host $host;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_pass https://localhost:8080;
        }

    listen [::]:443 ssl ipv6only=on; # managed by Certbot
    listen 443 ssl; # managed by Certbot
    ssl_certificate /etc/letsencrypt/live/www.awesomeinc.com/fullchain.pem; # managed by Certbot
    ssl_certificate_key /etc/letsencrypt/live/www.awesomeinc.com/privkey.pem; # managed by Certbot
    include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot





}
server {
    if ($host = awesomeinc.com) {
        return 301 https://$host$request_uri;
    } # managed by Certbot


    if ($host = www.awesomeinc.com) {
        return 301 https://$host$request_uri;
    } # managed by Certbot


	listen 80;
	listen [::]:80;

        server_name awesomeinc.com www.awesomeinc.com;
    return 404; # managed by Certbot




}
