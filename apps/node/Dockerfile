FROM node:19-alpine

WORKDIR /home/node/app

COPY package*.json .

RUN npm install

COPY . .

EXPOSE 8000

CMD ["npm", "start"]
