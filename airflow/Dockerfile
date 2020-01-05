FROM python:3.7

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
ENV AIRFLOW_HOME=/workspace/airflow
ENV AIRFLOW__CORE__LOAD_EXAMPLES=False

RUN airflow initdb
