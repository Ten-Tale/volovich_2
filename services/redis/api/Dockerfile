FROM python:3.12.2-slim-bullseye

WORKDIR /app

RUN apt-get update && apt-get -y install libpq-dev gcc

COPY requirements ./

RUN pip3 install -r requirements

COPY . .

CMD ["python3", "main.py"]