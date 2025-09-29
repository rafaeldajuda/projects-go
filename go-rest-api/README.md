# install node.js

```
https://nodejs.org/en/download

# Download and install nvm:
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.3/install.sh | bash
# in lieu of restarting the shell
\. "$HOME/.nvm/nvm.sh"
# Download and install Node.js:
nvm install 22
# Verify the Node.js version:
node -v # Should print "v22.20.0".
# Verify npm version:
npm -v # Should print "10.9.3".
```

```
# Agora abra um terminal na pasta do projeto e digite o seguinte comando para instalar as dependências do React:
npm install
# Atualize as dependências do npm com o seguinte comando:
npm update
# Para finalizar, ainda no terminal digite o comando abaixo para subir o servidor do React e aguarde:
NODE_OPTIONS=--openssl-legacy-provider npm start
```