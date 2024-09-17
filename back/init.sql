create table recipes (
    id serial primary key,
    nome varchar(255),
    tempo_preparo int,
    custo_aproximado float
);

create table ingredients (
    id serial primary key,
    recipe_id int,
    nome varchar(255)
);