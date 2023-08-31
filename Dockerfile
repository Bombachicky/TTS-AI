FROM node:18-alpine
WORKDIR /tts-ai
COPY . .
RUN npm i
CMD npm run dev
