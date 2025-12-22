FROM python:3.11-slim
WORKDIR /app
RUN pip install cs50 flask flask-session pytz requests
COPY . /app
EXPOSE 5000
CMD ["python", "app.py"]
