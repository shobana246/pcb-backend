# PCB Placement Tracking — Backend API

A Go backend providing CRUD operations for a PCB (Printed Circuit Board) component placement and inspection tracking system, inspired by a 3D placement-viewer style dashboard. Built as a job assignment.

Tracks **Projects** (PCB boards), the **Components** placed on each board, and **Placement Issues** found during inspection.

---

## Tech Stack


| Language | Go |
| Router | [Gin](https://github.com/gin-gonic/gin) |
| ORM | [GORM](https://gorm.io/) |
| Database | MySQL |
| Architecture | Layered — Handler → Service → Repository |


---

## Project Structure

```
backend/
├── cmd/
│   └── main.go            # entry point, route registration
├── database/
│   └── database.go        # DB connection (GORM + MySQL)
├── models/                # struct definitions
├── handlers/               # HTTP request/response layer
├── service/                 # business logic layer
├── repository/              # database query layer (GORM)
├── .env                    # DB credentials (not committed)
```

Request flow: **Router → Handler → Service → Repository → Database**

---

## Setup

### 1. Prerequisites
- Go installed
- MySQL server running

### 2. Create the database
Run the schema SQL (tables: `users`, `projects`, `component`, `placement_issues`) against your MySQL instance. See `database_schema_reference.md` for full column details and relationships.

### 3. Configure environment variables
Create a `.env` file inside `backend/` (same folder you'll run the app from):
```
DB_USER=root
DB_PASSWORD=yourpassword
DB_HOST=localhost
DB_PORT=3306
DB_NAME=pcb_placement_db
```

### 4. Install dependencies
```bash
go mod tidy
```

### 5. Run the server
```bash
go run cmd/main.go
```
Server starts on `http://localhost:8080`.

---

## API Endpoints

### Projects
| Method | Endpoint | Description |
|---|---|---|
| POST | `/projects` | Create a project |
| GET | `/projects` | List all projects |
| GET | `/projects/:id` | Get one project by ID |
| PUT | `/projects/:id` | Update a project |
| DELETE | `/projects/:id` | Soft-delete a project (and its components) |

### Components
| Method | Endpoint | Description |
|---|---|---|
| POST | `/components` | Create a component |
| GET | `/components?project_id=` | List components under a project |
| GET | `/components/:id` | Get one component by ID |
| PUT | `/components/:id` | Update a component |
| DELETE | `/components/:id` | Soft-delete a component |

### Placement Issues
| Method | Endpoint | Description |
|---|---|---|
| POST | `/issues` | Create a placement issue |
| GET | `/issues` | List all issues |
| GET | `/issues/:id` | Get one issue by ID |
| PUT | `/issues/:id` | Update an issue's status |
| DELETE | `/issues/:id` | Delete an issue (hard delete) |

Full sample request bodies for each endpoint are documented in the testing notes shared alongside this project.

---

## Key Design Decisions

- **Soft delete vs. hard delete:** Projects and Components use soft delete (`is_deleted` flag) since deleting either can cascade to a meaningful amount of dependent data (components, issues). Placement Issues use a hard delete, since nothing depends on them — there's no cascading risk to protect against.
- **No full CRUD for `users`:** Users exist only as reference data (linked via foreign keys from Projects and Placement Issues). Full user management was scoped out to focus on the core placement-tracking functionality.
- **IDs over duplicated names:** Foreign keys store IDs only. Where a readable name is useful in a response (e.g. a project's `updated_by` user name, or an issue's linked component name as `reference`), it's fetched via a JOIN at query time rather than stored as a duplicate column — this keeps the data from silently going out of sync.
- **No relational "nested" GET routes:** Component listing is filtered via a query parameter (`?project_id=`) rather than a full "get all" endpoint, since a single project can have hundreds of components — returning everything unfiltered isn't realistic or useful.

See `database_schema` for the full table structures and relationships.

---

## Out of Scope (by design, given assignment timeline)

- Authentication / JWT
- Full CRUD for `users`
- Automated tests
- Pagination
- Docker
