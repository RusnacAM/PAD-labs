FROM python:3.11-slim-bookworm

WORKDIR /gateway

COPY ./requirements.txt /gateway

RUN pip3 install -r requirements.txt

COPY . .

EXPOSE 5050

ENV FLASK_APP=main.py

CMD ["python3", "-m", "flask", "run", "--host=0.0.0.0", "--port=5050"]