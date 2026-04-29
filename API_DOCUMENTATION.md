# Documentación de la API: Clientes Grupo Proteger

Esta documentación describe los aspectos funcionales, reglas de negocio, modelos de datos y endpoints de la API **Clientes Grupo Proteger**, desarrollada en Go bajo una arquitectura limpia (Clean Architecture).

---

## 1. Arquitectura y Tecnologías

El proyecto está diseñado usando los principios de **Clean Architecture**, dividiendo las responsabilidades en capas (Dominio, Casos de uso/Servicios, Repositorios y Handlers).

- **Lenguaje:** Go (Golang)
- **Framework Web:** `go-chi/chi/v5` para el enrutamiento HTTP.
- **Bases de Datos:**
  - **MySQL:** Base de datos relacional principal (entidades: Clientes, Afiliados, Representantes Legales, Credenciales, Usuarios).
  - **MongoDB:** Base de datos NoSQL utilizada para almacenar el historial de fechas de ingreso y retiro de los afiliados.
- **Logger:** `go.uber.org/zap` para logs estructurados de alto rendimiento.
- **Autenticación:** JWT (JSON Web Tokens) gestionado a través de un Middleware (requerido para todas las rutas bajo `/api/v1` a excepción del login).

---

## 2. Entidades del Dominio y Reglas de Negocio

El sistema administra la información de clientes corporativos y sus respectivos afiliados (empleados o asociados), junto con las credenciales de acceso a diferentes plataformas.

### 2.1. Client (Cliente)
Representa a una empresa o entidad corporativa.
- **Campos principales:** ID, Nombre, NIT, Dirección, Teléfono, Email, ARL, y `IdRep` (Referencia al Representante Legal).
- **Regla de Negocio:** Un cliente debe estar asociado a un Representante Legal (`IdRep`).

### 2.2. LegalRep (Representante Legal)
Persona natural que representa a uno o varios clientes corporativos.
- **Campos principales:** ID, Nombre, CC (Cédula de Ciudadanía).

### 2.3. Affiliate (Afiliado)
Persona natural asociada a un Cliente.
- **Campos principales:** ID, Nombre, CC, EPS, Pensión, Riesgo, Caja (de compensación), Ingresos, Fecha de Nacimiento, Estado (`Status`), Fecha Último Pago, y `IdClient` (Referencia al Cliente).
- **Regla de Negocio:** 
  - Todo afiliado pertenece a un cliente (`IdClient`).
  - El historial de ingreso/retiro (`History` / `DateHistory`) se registra en MongoDB para mantener un rastreo de fechas sin afectar la tabla transaccional en MySQL.

### 2.4. Credentials (Credenciales)
Credenciales de acceso a portales u organizaciones externas pertenecientes a un cliente.
- **Campos principales:** ID, `IdClient`, Organización, Usuario, Contraseña.

### 2.5. History (Historial - MongoDB)
Documento en MongoDB que almacena las fechas en las que un afiliado ingresa o sale de cobertura.
- **Campos principales:** `_id` (Cédula o ID), `name`, `cc`, y un arreglo de `history` (Fecha de ingreso y de finalización).

---

## 3. Endpoints de la API

La API se expone bajo el prefijo `/api/v1`.
**Importante:** Todas las rutas, excepto el `/login`, están protegidas por el `AuthMiddleware` mediante JWT. Se debe incluir el header `Authorization: Bearer <token>`.

### Autenticación
- `POST /api/v1/login`
  - Inicia sesión validando credenciales y retorna un JWT.

### Clientes (`/clients`)
- `GET /api/v1/clients`: Obtiene la lista de todos los clientes.
- `GET /api/v1/clients/{id}`: Obtiene un cliente por su ID interno.
- `GET /api/v1/clients/nit/{nit}`: Busca un cliente por su NIT.
- `GET /api/v1/clients/legal-rep/{idLegalRep}`: Obtiene todos los clientes asociados a un ID de representante legal específico.
- `POST /api/v1/clients`: Crea un nuevo cliente.
- `PUT /api/v1/clients`: Actualiza la información de un cliente existente.
- `DELETE /api/v1/clients/{id}`: Elimina un cliente por su ID.

### Afiliados (`/affiliates`)
- `GET /api/v1/affiliates`: Obtiene la lista general de afiliados.
- `GET /api/v1/affiliates/{id}`: Obtiene un afiliado por su ID.
- `GET /api/v1/affiliates/cc/{cc}`: Busca un afiliado por su Cédula.
- `GET /api/v1/affiliates/client/{clientId}`: Obtiene la lista de afiliados que pertenecen a un cliente en particular.
- `POST /api/v1/affiliates`: Crea un afiliado. (Almacena también la fecha de ingreso en MongoDB).
- `PUT /api/v1/affiliates`: Actualiza un afiliado existente.
- `DELETE /api/v1/affiliates/{id}`: Elimina un afiliado.

### Representantes Legales (`/legal-reps`)
- `GET /api/v1/legal-reps`: Listado de representantes legales.
- `GET /api/v1/legal-reps/{id}`: Detalle de un representante.
- `GET /api/v1/legal-reps/cc/{cc}`: Búsqueda de un representante por Cédula.
- `POST /api/v1/legal-reps`: Creación de un representante legal.
- `PUT /api/v1/legal-reps`: Actualización de datos.
- `DELETE /api/v1/legal-reps/{id}`: Eliminación de un representante.

### Credenciales (`/credentials`)
- `GET /api/v1/credentials`: Lista de credenciales en el sistema.
- `GET /api/v1/credentials/{id}`: Obtiene una credencial por su ID.
- `GET /api/v1/credentials/client/{idClient}`: Obtiene todas las credenciales asociadas a un Cliente específico.
- `POST /api/v1/credentials`: Guarda una nueva credencial de un cliente.
- `PUT /api/v1/credentials`: Actualiza una credencial.
- `DELETE /api/v1/credentials/{id}`: Borra una credencial.

---

## 4. Middleware y Seguridad
- **Logger & Recoverer:** Se utiliza middleware de `chi` para asegurar que las caídas imprevistas (panics) se capturen adecuadamente y se genere un log de toda petición entrante.
- **CORS:** Configurado para aceptar orígenes cruzados (`*`) y los métodos principales (`GET`, `POST`, `PUT`, `DELETE`, `OPTIONS`).
- **Autenticación:** Intercepta todas las peticiones con `md.AuthMiddleware(authService)`. Si la validación del Token expira o es incorrecta, deniega el acceso con status `401 Unauthorized`.
