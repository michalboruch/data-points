FROM python:3.9-slim

ENV PYTHONUNBUFFERED=1

RUN apt-get update

COPY pip-requirements.txt /usr/src/app/pip-requirements.txt

RUN pip install --upgrade \
    pip install --no-cache-dir -r /usr/src/app/pip-requirements.txt

WORKDIR /usr/src/app

COPY . .

CMD ["./run.sh"]
