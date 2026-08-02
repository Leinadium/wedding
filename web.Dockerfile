FROM node:alpine AS builder
WORKDIR /app
COPY web/package*.json .
RUN npm ci
COPY web/ .

# overriding translations
RUN if [ -f "./src/assets/override.texts.json" ]; then cp ./src/assets/override.texts.json ./src/assets/texts.json; fi;

ARG VITE_API_URL
ARG VITE_FRONT_URL
ENV VITE_API_URL=$VITE_API_URL
ENV VITE_FRONT_URL=$VITE_FRONT_URL

RUN npm run build
RUN npm prune --production

FROM nginx:alpine
COPY infra/nginx.conf /etc/nginx/templates/default.conf.template
COPY --from=builder /app/dist /usr/share/nginx/html
EXPOSE 3000
