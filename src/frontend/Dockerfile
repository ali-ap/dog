FROM nginx

WORKDIR /usr/src/app

RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash -
RUN apt-get install -y nodejs

COPY package*.json ./

RUN npm install 

COPY . .

ARG REACT_APP_API_URL

ENV REACT_APP_API_URL $REACT_APP_API_URL

RUN npm run build

RUN rm -r /usr/share/nginx/html/*

RUN cp -a build/. /usr/share/nginx/html
