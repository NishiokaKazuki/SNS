FROM node:latest

WORKDIR /HEW2020/client
COPY . .
RUN yarn install --ignore-engines --network-timeout 1000000
CMD ["yarn", "start"]
EXPOSE 3000