# Utiliser une image de base Node.js
FROM node:14-alpine AS build

# Définir le répertoire de travail dans l'image Docker
WORKDIR /app

# Copier le package.json et le package-lock.json dans le répertoire de travail
COPY package*.json ./

# Installer les dépendances
RUN npm install

# Copier le reste des fichiers de l'application dans le répertoire de travail
COPY . .

# Compiler l'application Vue
RUN npm run build-only

# Etape de production
FROM nginx:alpine

# Copier les fichiers de build de l'étape précédente dans le répertoire de base de Nginx
COPY --from=build /app/dist /usr/share/nginx/html

# Exposer le port 80
EXPOSE 80

# Commande de démarrage du serveur Nginx
CMD ["nginx", "-g", "daemon off;"]
