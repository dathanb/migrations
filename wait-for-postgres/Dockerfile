from alpine

RUN apk update
RUN apk add bash

RUN mkdir /app
COPY wait-for-it.sh /app

CMD /app/wait-for-it.sh -h fakestack-postgres -p 5432
