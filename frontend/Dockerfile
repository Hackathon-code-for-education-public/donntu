FROM node:alpine
WORKDIR /usr/app
COPY ./package*.json .npmrc ./
RUN npm install --omit=dev
COPY ./ ./
RUN npm run build
EXPOSE 3000
USER node
CMD [ "npm", "start" ]