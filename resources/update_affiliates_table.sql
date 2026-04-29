USE grupo_proteger;

-- 1. Agregar las nuevas columnas
ALTER TABLE affiliates 
ADD COLUMN entry_date DATE,
ADD COLUMN last_payment_date DATE;

-- 2. Actualizar los registros existentes con una fecha por defecto (para evitar nulos)
UPDATE affiliates 
SET entry_date = '2023-01-01', 
    last_payment_date = '2026-01-01';

-- 3. Eliminar la antigua columna 'income'
ALTER TABLE affiliates 
DROP COLUMN income;
