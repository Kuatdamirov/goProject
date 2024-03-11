# goProject
Проект представляет собой веб-приложение, которое позволяет создавать, просматривать, редактировать и удалять автомобили. 
# Структура API
POST /cars 
GET /cars/:id
PUT /cars/:id
DELETE /cars/:id
# DB Structure
// Use DBML to define your database structure
// Docs: https://dbml.dbdiagram.io/docs

Table carDealers {
  id bigserial [primary key]
  created_at timestamp
  updated_at timestamp
  title text
  description text
  coordinates text
  address text
  country text
}

Table car {
  id bigserial [primary key]
  created_at timestamp
  updated_at timestamp
  title text
  description text
  year int
}

// many-to-many
Table carDealers_and_car {
  id bigserial [primary key]
  created_at timestamp
  updated_at timestamp
  carDealers bigserial
  car bigserial
}

# Участники команды
Дамиров Куат 22B030335
Байбатчаев Азамат 22B030319
../schema.dbml
