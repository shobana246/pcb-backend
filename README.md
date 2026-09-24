# PCB Placement Tracking: Backend API

A Go backend providing CRUD operations for a PCB (Printed Circuit Board) component placement and inspection tracking system. Built as a job assignment.

Tracks **Projects** (PCB boards), the **Components** placed on each board, and **Placement Issues** found during inspection.

> All source code is inside the [`backend/`](backend) folder.

---

## Tech Stack

| Layer | 
|---|---|
| Language | Go |
| Router | [Gin](https://github.com/gin-gonic/gin) |
| ORM | [GORM](https://gorm.io/) |
| Database | MySQL |
| Architecture | Layered: Handler, Service, Repository |

---

## How a request flows

```
Router (cmd/main.go)
  -> Handler     reads the JSON, checks its shape (validation), sends the reply
  -> Service     business rules (does the project exist? is this status change allowed?)
  -> Repository  database queries only (GORM)
  -> MySQL
```

| Layer | Folder | Responsibility |
|---|---|---|
| Models | `backend/models` | Table structs, plus `request.go` (what the client is allowed to send, with validation rules) |
| Handlers | `backend/handlers` | HTTP in and out, status codes |
| Service | `backend/service` | Business rules |
| Repository | `backend/repository` | Database queries |

---

## Setup

1. Install Go and MySQL.
2. Create the database and tables: `mysql -u root -p < schema.sql`
3. Copy `.env.example` to `backend/.env` and fill in your values.
4. Install dependencies: `go mod tidy`
5. Run the server from the `backend/` folder: `go run cmd/main.go`

The server starts on `http://localhost:8080`.

---

## API Endpoints

### Projects

| Method | Endpoint | Description |
|---|---|---|
| POST | `/projects` | Create a project |
| GET | `/projects` | List projects (deleted ones are hidden) |
| GET | `/projects/:id` | Get one project |
| PUT | `/projects/:id` | Update a project |
| DELETE | `/projects/:id` | Soft-delete a project and its components (one transaction) |

```bash
curl -X POST localhost:8080/projects -H "Content-Type: application/json" \
  -d '{"project_name":"Board A","revision":"Rev B","updated_by":1}'
```

### Components

| Method | Endpoint | Description |
|---|---|---|
| POST | `/components` | Create a component |
| GET | `/components?project_id=` | List components of a project |
| GET | `/components/:id` | Get one component |
| PUT | `/components/:id` | Update a component (`project_id` cannot be changed) |
| DELETE | `/components/:id` | Soft-delete a component |

```bash
curl -X POST localhost:8080/components -H "Content-Type: application/json" \
  -d '{"project_id":1,"component_name":"C236","package":"0603","category":"Passive",
       "placement_side":"Top","rotation":90,"x_position":12.5,"y_position":30.25,
       "height":0.5,"supplier":"Murata","part_number":"GRM188",
       "tolerance_position":0.1,"tolerance_rotation":5}'
```

### Placement Issues

| Method | Endpoint | Description |
|---|---|---|
| POST | `/issues` | Create an issue (starts as `open`) |
| GET | `/issues` | List issues (`reference` = component name) |
| GET | `/issues/:id` | Get one issue |
| PUT | `/issues/:id` | Update the issue status |
| DELETE | `/issues/:id` | Delete an issue (hard delete) |

```bash
curl -X POST localhost:8080/issues -H "Content-Type: application/json" \
  -d '{"component_id":1,"issue":"Rotation mismatch","severity":"warning"}'

curl -X PUT localhost:8080/issues/1 -H "Content-Type: application/json" \
  -d '{"status":"in review"}'
```

---

## Validation rules

Checked in the request structs (`backend/models/request.go`). A failure returns **400**.

| Field | Rule |
|---|---|
| Project `project_name` / `revision` | required, max 150 / 100 characters |
| Project `updated_by` | required, greater than 0 |
| Component `component_name` | required, max 100 |
| Component `category` | one of `Passive`, `IC`, `Connector`, `Mechanical` |
| Component `placement_side` | `Top` or `Bottom` |
| Component `status` | optional; `verified`, `warning`, `critical` or `unverified` (default `unverified`) |
| Component `rotation` | 0 up to (not including) 360 |
| Component `x_position`, `y_position` | -999 to 999 |
| Component `height`, `tolerance_position`, `tolerance_rotation` | 0 to 99 |
| Issue `severity` | `warning` or `critical` |
| Issue `status` (update) | `open`, `in review` or `resolved` |

## Business rules

Checked in the service layer (`backend/service`).

- Blank or space-only names are rejected.
- `updated_by` must be an existing user.
- A project name + revision must be unique among live projects (**409**).
- A component can only be added to a project that exists and is not deleted (**404**).
- A component name must be unique within its project (**409**).
- An issue can only be added to a component that exists and is not deleted (**404**).
- Issue status changes must follow: `open` to `in review`; `in review` to `open` or `resolved`; `resolved` to `in review`. Anything else returns **409**.
- Deleting a project soft-deletes its components in one transaction, so both succeed or neither does.

## Status codes

| Code | Meaning |
|---|---|
| 200 | Success |
| 400 | Invalid input or invalid id |
| 404 | Record not found (or deleted) |
| 409 | Conflict: duplicate name or status change not allowed |
| 500 | Server or database error (details are logged, not returned) |

---

## Key Design Decisions

- **Soft delete vs. hard delete:** Projects and Components use soft delete (`is_deleted` flag) since deleting either can cascade to a meaningful amount of dependent data (components, issues). Placement Issues use a hard delete, since nothing depends on them — there's no cascading risk to protect against.
- **No full CRUD for `users`:** Users exist only as reference data (linked via foreign keys from Projects and Placement Issues). Full user management was scoped out to focus on the core placement-tracking functionality.
- **IDs over duplicated names:** Foreign keys store IDs only. Where a readable name is useful in a response (e.g. a project's `updated_by` user name, or an issue's linked component name as `reference`), it's fetched via a JOIN at query time rather than stored as a duplicate column — this keeps the data from silently going out of sync.
- **No relational "nested" GET routes:** Component listing is filtered via a query parameter (`?project_id=`) rather than a full "get all" endpoint, since a single project can have hundreds of components — returning everything unfiltered isn't realistic or useful.

See `database_schema_reference.md` for the full table structures and relationships.

---

## Out of Scope (by design, given assignment timeline)

- Authentication / JWT
- Full CRUD for `users`
- Automated tests
- Pagination
- Docker
