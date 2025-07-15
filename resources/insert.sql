USE grupo_proteger;
INSERT INTO legal_rep (id, name, cc) VALUES
(1, 'Carlos Rodríguez', '1002003000'),
(2, 'Laura Gómez', '1002004000');
INSERT INTO clients (id, name, nit, address, phone, email, id_rep, arl) VALUES
(1, 'AgroInnova S.A.S', '900123456-7', 'Calle 45 #32-10, Medellín', '3011234567', 'contacto@agroinnova.com', 1, 'Colpatria Riesgos'),
(2, 'TecnoRiego Ltda', '901987654-3', 'Carrera 21 #15-30, Cali', '3029876543', 'info@tecnoriego.com', 2, 'SURA Riesgos');
INSERT INTO affiliates (id, name, cc, eps, status, id_client, pension, risk, birthdate, caja, income) VALUES
(1, 'Juan Pérez', '1020304050', 'Sura EPS', 'ACTIVE', 1, 'Protección', '0.5', '1990-06-15', 'Comfama', 2500000.00),
(2, 'Ana Torres', '1122334455', 'Nueva EPS', 'RETIRED', 1, 'Colpensiones', '2', '1985-09-10', 'Comfama', 2800000.00),
(3, 'Luis Gómez', '2233445566', 'Sanitas', 'ACTIVE', 2, 'Porvenir', '3', '1995-02-25', 'Cafam', 2000000.00);
INSERT INTO credentials (id, id_client, organization, user, pass) VALUES
(1, 1, 'Plataforma Pila', 'agroadmin', 'P@ssAgro123'),
(2, 1, 'Sistema Nómina', 'agronomina', 'N0minaAgro!'),
(3, 2, 'Plataforma Pila', 'tecnoadmin', 'T3cnoPila'),
(4, 2, 'Sistema Nómina', 'tecnoriego_nom', 'RiegoN0mina!');
INSERT INTO user (user, password) VALUES
('admin', 'admin123'),  -- Para pruebas; idealmente usar hash en producción
('agro_user', 'Agro456'),
('riego_user', 'Riego789');

