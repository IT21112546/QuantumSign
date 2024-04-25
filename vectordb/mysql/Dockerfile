FROM python:3-slim

COPY . /app
WORKDIR /app

RUN pip install -r requirements.txt
CMD ["python", "-u", "main.py"]
