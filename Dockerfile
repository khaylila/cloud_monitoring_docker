# Tahap 1: Gunakan Ubuntu 24.04 sebagai dasar
FROM ubuntu:24.04

# Set environment agar tidak ada prompt interaktif saat instalasi
ENV DEBIAN_FRONTEND=noninteractive

# Update dan install dependensi yang dibutuhkan
RUN apt-get update && apt-get install -y \
    wget \
    ca-certificates \
    supervisor \
    && rm -rf /var/lib/apt/lists/*

#--- Instalasi Go ---
# Tentukan versi Go dan arsitektur
ENV GO_VERSION 1.25.1
ENV GO_ARCH amd64

# Download, ekstrak, dan install Go
RUN wget https://golang.org/dl/go${GO_VERSION}.linux-${GO_ARCH}.tar.gz -O /tmp/go.tar.gz && \
    tar -C /usr/local -xzf /tmp/go.tar.gz && \
    rm /tmp/go.tar.gz

# Set PATH untuk Go
ENV PATH="/usr/local/go/bin:${PATH}"

#--- Instalasi Node Exporter ---
# Tentukan versi Node Exporter dan arsitektur
ENV NODE_EXPORTER_VERSION 1.9.1
ENV NODE_EXPORTER_ARCH amd64

# Download, ekstrak, dan install Node Exporter
RUN wget https://github.com/prometheus/node_exporter/releases/download/v${NODE_EXPORTER_VERSION}/node_exporter-${NODE_EXPORTER_VERSION}.linux-${NODE_EXPORTER_ARCH}.tar.gz -O /tmp/node_exporter.tar.gz && \
    tar -C /tmp -xzf /tmp/node_exporter.tar.gz && \
    mv /tmp/node_exporter-${NODE_EXPORTER_VERSION}.linux-${NODE_EXPORTER_ARCH}/node_exporter /usr/local/bin/ && \
    rm -rf /tmp/node_exporter.tar.gz /tmp/node_exporter-${NODE_EXPORTER_VERSION}.linux-${NODE_EXPORTER_ARCH}

#--- Build Aplikasi Go ---
# Buat direktori kerja untuk aplikasi
WORKDIR /app

# Salin kode sumber aplikasi
COPY app/ .

# Build aplikasi Go. Outputnya akan bernama 'hello-api'
RUN go build -o hello-api .

#--- Konfigurasi Akhir ---
# Salin file konfigurasi Supervisor
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

# Ekspos port untuk API (8080) dan Node Exporter (9100)
EXPOSE 8080
EXPOSE 9100

# Jalankan Supervisor saat kontainer dimulai
CMD ["/usr/bin/supervisord"]