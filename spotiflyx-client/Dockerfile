# Utiliser une image de base Node.js
FROM node:20-alpine

# Définir le répertoire de travail dans l'image Docker
WORKDIR /app

# Copier le package.json et le package-lock.json dans le répertoire de travail
COPY . .

# Installer les dépendances
RUN npm install

# Compiler l'application Vue
RUN npm run build-only

RUN npm install --global serve

COPY serve.json /app/dist/serve.json

# Exposer le port 8080
EXPOSE 8080

WORKDIR /app/dist

# Démarrer l'application
CMD ["serve", "-l", "8080"]
