package stubs

// NginxSymfonyTpl is a stub for a Nginx virtual host for Symfony2 projects
var nginxSymfonyTpl = `
server {
    listen {{.Port}};
    listen {{.SslPort}} ssl http2;
    server_name {{.ServerName}};
    root {{.ServerRoot}};

    ssl_certificate     /etc/nginx/ssl/{{.ServerName}}.crt;
    ssl_certificate_key /etc/nginx/ssl/{{.ServerName}}.key;

    charset utf-8;

    location / {
        try_files $uri /app_dev.php$is_args$args;
    }

    # DEV
    location ~ ^/(app_dev|config)\.php(/|$) {
        #fastcgi_pass unix:/var/run/php5-fpm.sock;
        fastcgi_pass unix:/var/run/php/php7.0-fpm.sock;
        fastcgi_split_path_info ^(.+\.php)(/.*)$;
        include fastcgi_params;
        fastcgi_param SCRIPT_FILENAME $realpath_root$fastcgi_script_name;
        fastcgi_param DOCUMENT_ROOT $realpath_root;
        fastcgi_intercept_errors off;
        fastcgi_buffer_size 16k;
        fastcgi_buffers 4 16k;
    }

    # PROD
    location ~ ^/app\.php(/|$) {
        #fastcgi_pass unix:/var/run/php5-fpm.sock;
        fastcgi_pass unix:/var/run/php/php7.0-fpm.sock;
        fastcgi_split_path_info ^(.+\.php)(/.*)$;
        include fastcgi_params;
        fastcgi_param SCRIPT_FILENAME $realpath_root$fastcgi_script_name;
        fastcgi_param DOCUMENT_ROOT $realpath_root;
        fastcgi_intercept_errors off;
        fastcgi_buffer_size 16k;
        fastcgi_buffers 4 16k;
        internal;
    }

    location ~ \.php$ {
      return 404;
    }

    location = /favicon.ico { access_log off; log_not_found off; }
    location = /robots.txt  { access_log off; log_not_found off; }

    access_log off;
    error_log  /var/log/nginx/{{.ServerName}}-error.log error;

    sendfile off;

    client_max_body_size 100m;
}

`
