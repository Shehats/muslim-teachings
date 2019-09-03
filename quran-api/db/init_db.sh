#!/bin/bash
/usr/bin/mysqld_safe --skip-grant-tables &
sleep 5
mysql -u root -e "CREATE DATABASE QuranDB"
mysql -u root QuranDB < /tmp//arabic_quran.sql