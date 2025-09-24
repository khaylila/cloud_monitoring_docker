# Praktek cloud monitoring menggunakan docker (INSTALASI)

## clone data terlebih dahulu github.com/khaylila/cloud_monitoring_docker
`git clone https://github.com/khaylila/cloud_monitoring_docker`

## masuk ke dalam folder main
`cd /opt/temp/praktek_cc/ubuntu_go_exporter`

## masuk kedalam folder golang_app untuk membuat image
`cd golang_app`
`docker build -t golang_app .`

### masuk kedalam folder grafana, lalu ubah data email pada grafana.ini (optional jika ingin menggunakan alerting)
#### cari parameter berikut, lalu sesuaikan dengan email
##### sebelum
​```sh
[smtp]
enabled = false
host = localhost:25
user = 
password = 
from_address = 
from_name = 
​```
##### sesudah
​```sh
[smtp]
enabled = true
host = smtp.gmail.com:587
user = hujanturun@gmail.com
password = lorem ipsum
from_address = hujanturun@gmail.com
from_name = Grafana
​```
## kembali ke folder main dan lakukan docker compose
`cd /opt/temp/praktek_cc/ubuntu_go_exporter`
`docker-compose up -d`

