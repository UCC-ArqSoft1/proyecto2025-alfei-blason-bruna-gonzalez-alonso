CREATE DATABASE IF NOT EXISTS backend;
USE backend;

DROP TABLE IF EXISTS inscripciones;
DROP TABLE IF EXISTS horarios;
DROP TABLE IF EXISTS act_deportivas;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios (
    id_usuario INT AUTO_INCREMENT PRIMARY KEY,
    nombre_usuario VARCHAR(255) UNIQUE,
    contrasenia_hash VARCHAR(255),
    nombre VARCHAR(255),
    apellido VARCHAR(255),
    dni INT,
    mail VARCHAR(255),
    is_admin BOOLEAN,
    foto VARCHAR(255)
);

CREATE TABLE horarios (
    id_horario INT AUTO_INCREMENT PRIMARY KEY,
    id_actividad INT,
    dia VARCHAR(255),
    horario_inicio TIME,
    horario_fin TIME,
    cupos INT
);

CREATE TABLE act_deportivas (
    id_actividad INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(255),
    nombre_profesor VARCHAR(255),
    foto VARCHAR(255),
    descripcion VARCHAR(255)
);

CREATE TABLE inscripciones (
    id_inscripcion INT AUTO_INCREMENT PRIMARY KEY,
    id_usuario INT,
    id_actividad INT,
    id_horario INT

);

