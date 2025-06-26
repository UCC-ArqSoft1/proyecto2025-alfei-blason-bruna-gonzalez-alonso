# Imagen base
FROM node:latest

# Directorio de trabajo
WORKDIR /app

# Copiar package.json y package-lock.json
COPY package*.json ./

# Instalar dependencias
RUN npm install

# Copiar el código fuente
COPY . .

# Construir la aplicación
RUN npm run build

# Instalar servidor ligero para servir la aplicación
RUN npm install -g serve

# Exponer el puerto
EXPOSE 3000

# Comando para ejecutar la aplicación
CMD ["serve", "-s", "build", "-l", "3000"]