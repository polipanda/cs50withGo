FROM python:3.11-slim
WORKDIR /app
COPY . /app
RUN pip install flask
RUN pip install cs50
EXPOSE 5000
CMD ["python", "app.py"]
