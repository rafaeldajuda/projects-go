# INSTALANDO DOCKER E DOCKER COMPOSE

* Ref: https://www.nerdlivre.com.br/instalando-docker-e-docker-compose-no-ubuntu-24-04/

```sh
sudo apt update && sudo apt upgrade -y
```

## Atualizando o Sistema

```sh
sudo apt update && sudo apt upgrade -y
```

## Docker

1. Pré-requisitos: Instale pacotes que permitem o uso de HTTPS para downloads seguros:
```sh
sudo apt install apt-transport-https ca-certificates curl software-properties-common
```

2. Adicionando a chave GPG e o repositório oficial do Docker:
```sh
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
```
3. Instalando o Docker:
```sh
sudo apt update
sudo apt install docker-ce
```

4. Verificando a instalação:
```sh
docker --version
```

Para usar o Docker sem sudo, adicione seu usuário ao grupo apropriado:
```sh
sudo usermod -aG docker ${USER}
su - ${USER}
```

## Docker Compose

1. Criando o diretório para plugins:
```sh
mkdir -p ~/.docker/cli-plugins/
```

2. Baixando e configurando o Docker Compose:
```sh
curl -SL https://github.com/docker/compose/releases/download/v2.27.1/docker-compose-linux-x86_64 -o ~/.docker/cli-plugins/docker-compose
chmod +x ~/.docker/cli-plugins/docker-compose
```
4. Verificando a instalação:
```sh
docker compose version
```