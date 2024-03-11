FROM node:18.17.0 as build

WORKDIR /app

COPY package.json ./

RUN npm install

COPY . .

RUN mkdir ./public

RUN npm run build

FROM node:18.17.0-alpine as release

WORKDIR /app

COPY --from=build --chown=nextjs:nodejs /app/.next ./.next
COPY --from=build /app/next.config.js ./
COPY --from=build /app/node_modules ./node_modules
COPY --from=build /app/package.json ./
COPY --from=build /app/public ./public

EXPOSE 3000

CMD ["npm", "start"]