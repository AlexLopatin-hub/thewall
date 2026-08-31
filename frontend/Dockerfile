FROM node:26-alpine AS build-stage

WORKDIR /app

COPY . .

RUN npm install
RUN npm run build

FROM nginx:1.31.4 AS deploy-stage

COPY --from=build-stage /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
