package stubs

// SievertYmlTpl is a stub for the sievert.yml file
var sievertYmlTpl = `
ip: "127.0.0.1"
services:
  - mysql
  #- sqlite
  #- postgresql
  #- elasticsearch
  #- redis
  #- memcached

sites:
  - map: joinfinity.lan
    to: /home/public/html
    template: laravel
  - map: improveyourcity.lan
    to: /home/public/html_iyc/web/mtc
    template: symfony
  - map: api.improveyourcity.lan
    to: /home/public/html_iyc/web/open010
    template: symfony
databases:
  mysql:
    - user: joinfinity
      database: joinfinity_local
      password: secret
    - user: wordpress
      database: wordpress_local
      password: secret
variables:
    - key: APP_ENV
      value: local
`
