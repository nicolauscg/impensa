FROM node:14.5-alpine AS build

WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH
COPY package.json /app/package.json
RUN npm install --silent
RUN npm install react-scripts@2.1.8 -g --silent
COPY . /app
EXPOSE 80
RUN npm run build

FROM nginx:1.16.0-alpine

COPY --from=build /app/build /var/www
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx/nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
