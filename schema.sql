CREATE TABLE urls
(
  id INT AUTO_INCREMENT PRIMARY KEY,
  redirect_name varchar(55) NOT NULL UNIQUE,
  original_url varchar(55) NOT NULL UNIQUE
);
