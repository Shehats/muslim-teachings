FROM mysql:5.7
ADD arabic_quran.sql /tmp/arabic_quran.sql
RUN /bin/bash -c "/usr/bin/mysqld_safe --skip-grant-tables &" && \
  sleep 5 && \
  mysql -u root -e "CREATE DATABASE Quran_db" && \
  mysql -u root Quran_db < /tmp/arabic_quran.sql
EXPOSE 3306