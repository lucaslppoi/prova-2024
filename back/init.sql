create table recipes (
    id int,
    nome varchar(255),
    tempo_preparo int,
    custo_aproximado float
);

create table ingredientes (
    id int,
    recipe_id int,
    nome varchar(255)
);