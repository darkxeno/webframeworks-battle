FROM mongo:3.3

COPY restore-db.sh /restore-db.sh
COPY mongodb-test.archive /mongodb-test.archive

ENTRYPOINT ["/restore-db.sh"]

EXPOSE 27017
CMD ["mongod"]