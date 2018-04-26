FROM node:latest

ENV NODE_ENV $app_env

RUN mkdir /frontend
WORKDIR /frontend

ENV PATH /frontend/node_modules/.bin:$PATH

COPY package.json /frontend/package.json
RUN npm install --silent
RUN npm install react-scripts@1.1.1 -g --silent

# If dev settings will use create-react start script for hot code relaoding via docker-compose shared volume
# If production setting will build optimized static files and serve using http-server
CMD if [ ${NODE_ENV} = production ]; \
	then \
	npm install -g http-server && \
	npm build && \
	cd build && \
	hs -p 3000; \
	else \
	npm start; \
	fi

EXPOSE 3000