CREATE TABLE product(
	id bigint GENERATED ALWAYS AS IDENTITY,
	name varchar NOT NULL,
	price bigint NOT NULL,
	description varchar(1000) NOT NULL,
	quantity int NOT NULL,
	created_at timestamp default now()
);
