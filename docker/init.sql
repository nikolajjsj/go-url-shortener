SELECT 'CREATE DATABASE shortener'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'shortener')\gexec