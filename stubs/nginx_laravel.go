package stubs

// NginxLaravelTpl is a stub for a Nginx virtual host for Laravel projects
var NginxLaravelTpl = `
server {
    listen {{.Port}};
    listen {{.SslPort}} ssl http2;
    server_name {{.ServerName}};
    root {{.ServerRoot}};

    ssl_certificate     /etc/nginx/ssl/{{.ServerName}}.crt;
    ssl_certificate_key /etc/nginx/ssl/{{.ServerName}}.key;

    index index.html index.htm index.php;

    charset utf-8;

    location / {
        try_files $uri $uri/ /index.php$is_args$args;
    }

    location ~ \.php$ {
        fastcgi_split_path_info ^(.+\.php)(/.+)$;
        fastcgi_pass unix:/var/run/php5-fpm.sock;
        fastcgi_index index.php;
        include fastcgi_params;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        fastcgi_intercept_errors off;
        fastcgi_buffer_size 16k;
        fastcgi_buffers 4 16k;
    }

    location ~ /\.ht {
        deny all;
    }

    location = /favicon.ico { access_log off; log_not_found off; }
    location = /robots.txt  { access_log off; log_not_found off; }

    access_log off;
    error_log  /var/log/nginx/{{.ServerName}}-error.log error;

    sendfile off;

    client_max_body_size 100m;
}
`
