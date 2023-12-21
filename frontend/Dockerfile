# Use an appropriate base image for your frontend
# For example:
FROM node:latest

# Set the working directory inside the container
WORKDIR /app/frontend

# Copy necessary files into the container
COPY . .

# Install dependencies and build your frontend
RUN npm install
RUN npm run build

EXPOSE 5173/tcp

# Set the entry point or command to run the frontend
CMD ["npm", "run", "dev", "--", "--host"]
