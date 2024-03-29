CREATE TABLE categories (
    id varchar(20) NOT NULL,
    name varchar(50) NOT NULL,
    created_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP 
        ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id), 
    UNIQUE KEY ux_name (name)
) ENGINE=InnoDB;