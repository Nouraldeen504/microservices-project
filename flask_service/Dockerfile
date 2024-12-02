# flask_service/Dockerfile
FROM python:3.13-slim

# Set the working directory
WORKDIR /app

# Copy requirements and install
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy the rest of the application
COPY . .

# Expose the port
EXPOSE 5000

# Command to run the app
CMD ["python", "app.py"]