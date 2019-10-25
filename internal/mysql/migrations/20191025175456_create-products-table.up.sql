CREATE TABLE products (
    id varchar(20) NOT NULL,
    store_id varchar(20) NOT NULL,
    category_id varchar(20) NOT NULL,
    name varchar(100) NOT NULL,
    price decimal(19, 4) NOT NULL,
    description text NOT NULL,
    quantity int NOT NULL,
    created_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP 
        ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id),
    INDEX ix_strcat (store_id, category_id),
    INDEX ix_cat (category_id)
) ENGINE=InnoDB;