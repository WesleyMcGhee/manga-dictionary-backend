package database

var Schema = `
  CREATE TABLE user (
    username text,
    email text,
    password text,
  );
`
