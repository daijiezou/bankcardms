FROM 10.10.10.149:80/gemini-platform/gemini-base:v1.0.0

COPY Shanghai /etc/localtime
COPY start.sh /usr/start.sh
COPY bank-card-ms /usr/bank-card-ms


