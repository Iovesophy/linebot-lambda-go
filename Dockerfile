FROM node:14.18.1 AS dev-env
WORKDIR /app
RUN npm -g install serverless
CMD [ "./slsinit.sh" ]