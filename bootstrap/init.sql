CREATE TABLE "users"
(
    "id"         serial PRIMARY KEY,
    "first_name" varchar(50) NOT NULL,
    "last_name"  varchar(50) NOT NULL,
    "email"      varchar(150) UNIQUE,
    "password"   varchar(255),
    "created_at" timestamp DEFAULT now(),
    "updated_at" timestamp
);

COMMENT
    ON TABLE "users" IS '
Tabla de usuarios
';


-- Insertar datos de prueba

INSERT INTO users ("first_name", "last_name", "email", "password")
VALUES ('Juan', 'Pérez', 'admin@admin', '$2a$10$hPbh.Z57coBKvzvzt87yCOCFFV4NmPFmtjdOmsGQNqES4jxqgRBES');



