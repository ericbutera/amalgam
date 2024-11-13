# Use the Node base image to handle TypeScript compilation
FROM node:18-alpine AS builder

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm install

COPY . .

RUN npm run bundle

FROM grafana/k6:latest
WORKDIR /app

# Copy the transpiled JavaScript files to the runtime image
COPY --from=builder /app/dist /app/dist

# Run the k6 script
# CMD ["npm", "test"]
ENTRYPOINT [ "k6", "run", "/app/dist/load-test.js" ]
