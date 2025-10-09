FROM python:3.7

RUN pip install virtualenv

RUN virtualenv venv

ENTRYPOINT bin/bash

RUN /bin/bash -c "source venv/bin/activate"

RUN pip install apache-beam[gcp]
