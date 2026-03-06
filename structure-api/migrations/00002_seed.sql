-- +goose Up
-- Создаём департаменты
INSERT INTO departments (name, parent_id) VALUES ('Backend', NULL);
INSERT INTO departments (name, parent_id) VALUES ('Frontend', NULL);
INSERT INTO departments (name, parent_id) VALUES ('Platform', 1);

-- Создаём сотрудников
INSERT INTO employees (department_id, full_name, position) VALUES (1, 'Ivan Ivanov', 'Go Developer');
INSERT INTO employees (department_id, full_name, position) VALUES (1, 'Olga Smirnova', 'Backend Developer');
INSERT INTO employees (department_id, full_name, position) VALUES (2, 'Anna Petrova', 'React Developer');
INSERT INTO employees (department_id, full_name, position) VALUES (3, 'Sergey Kuznetsov', 'Platform Engineer');

-- +goose Down
DELETE FROM employees;
DELETE FROM departments;