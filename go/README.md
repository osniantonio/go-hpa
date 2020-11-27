# Kubernetes e hpa

## URL da imagem publicada no docker hub
https://hub.docker.com/r/osniantonio/go-hpa

## Comando usado para criar a imagem
docker build -t osniantonio/go-hpa .

## Comando para baixar a imagem
docker pull osniantonio/go-hpa

## Comando para executar
docker run -d -p 8000:8000 osniantonio/go-hpa