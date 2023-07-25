# FROM node:13.12.0-alpine
FROM node:16.17.1-alpine3.15

ARG WS_API_URL

ENV REACT_APP_API_URL=${WS_API_URL}

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

RUN npm install -g serve
# CMD ["npm", "run", "build"]

# RUN serve -s build -l 3000
