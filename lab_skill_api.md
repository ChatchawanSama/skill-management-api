# Skill management API

## Instructions

1. MUST create a new project using Go [Gin Gonic framework](https://github.com/gin-gonic/gin?tab=readme-ov-file#running-gin)
2. Could pair with a partner / solo
3. MUST write e2e tests for the API using [Playwright](https://playwright.dev/docs/api-testing)
4. Should write unit tests for the API using Go testing package
5. MUST use PostgreSQL as the database
6. MUST use Docker compose to run database
7. MUST take DATABASE_URL as an environment variable
8. submit link to the repository via email or chat group
9. You can add any techniques or tools you think are necessary to make the project better

## Context

This API is used to manage skills. It allows to create, read, update and delete skills.
API functionalities include:
- `GET /api/v1/skills/:key` - Get a skill by key
- `GET /api/v1/skills` - Get all skills
- `POST /api/v1/skills` - Create a skill
- `PUT /api/v1/skills/:key` - Update a skill
- `PATCH /api/v1/skills/:key/actions/name` - Update the name of a skill
- `PATCH /api/v1/skills/:key/actions/description` - Update the description of a skill
- `PATCH /api/v1/skills/:key/actions/logo` - Update the logo of a skill
- `PATCH /api/v1/skills/:key/actions/tags` - Update the tags of a skill
- `DELETE /api/v1/skills/:key` - Delete a skill

## Skill table

```sql
-- create skill table
CREATE TABLE skill (
	key TEXT PRIMARY KEY,
	name TEXT NOT NULL DEFAULT '',
	description TEXT NOT NULL DEFAULT '',
	logo TEXT NOT NULL DEFAULT '',
	tags TEXT [] NOT NULL DEFAULT '{}'
);
```

## API Specs

1. `GET /api/v1/skills/:key`

description: Get a skill by key

example: GET /api/v1/skills/go

response:

```json
{
	"status": "success",
	"data": {
		"key": "go",
		"name": "Go",
		"description": "Go is a statically typed, compiled programming language designed at Google.",
		"logo": "https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg",
		"tags": ["programming language", "system"]
	}
}
```

failure response:

```json
{
	"status": "error",
	"message": "Skill not found"
}
```

2. `GET /api/v1/skills`

description: Get all skills

example: GET /api/v1/skills

response:

```json
{
    "status":"success",
	"data": [
		{
			"key": "go",
			"name": "Go",
			"description": "Go is a statically typed, compiled programming language designed at Google.",
			"logo": "https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg",
			"tags": ["programming language", "system"]
		},
		{
			"key": "nodejs",
			"name": "Node.js",
			"description": "Node.js is an open-source, cross-platform, JavaScript runtime environment that executes JavaScript code outside of a browser.",
			"logo": "https://upload.wikimedia.org/wikipedia/commons/d/d9/Node.js_logo.svg",
			"tags": ["runtime", "javascript"]
		}
	]
}
```

3. `POST /api/v1/skills`

description: Create a skill

example: POST /api/v1/skills

payload:

```json
{
	"key": "python",
	"name": "Python",
	"description": "Python is an interpreted, high-level, general-purpose programming language.",
	"logo": "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
	"tags": ["programming language", "scripting"]
}
```

response:

```json
{
	"status": "success",
	"data": {
		"key": "python",
		"name": "Python",
		"description": "Python is an interpreted, high-level, general-purpose programming language.",
		"logo": "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
		"tags": ["programming language", "scripting"]
	}
}
```

failure response:

```json
{
	"status": "error",
	"message": "Skill already exists"
}
```

4. `PUT /api/v1/skills/:key`

description: Update a skill

example: PUT /api/v1/skills/python

payload:

```json
{
	"name": "Python 3",
	"description": "Python 3 is the latest version of Python programming language.",
	"logo": "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
	"tags": ["data"]
}
```

response:

```json
{
	"status": "success",
	"data": {
		"key": "python",
		"name": "Python 3",
		"description": "Python 3 is the latest version of Python programming language.",
		"logo": "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
		"tags": ["data"]
	}
}
```

failure response:

```json
{
	"status": "error",
	"message": "not be able to update skill"
}
```

5. `PATCH /api/v1/skills/:key/actions/name`

description: Update the name of a skill

example: PATCH /api/v1/skills/python/actions/name

payload:

```json
{
	"name": "Python 3"
}
```

response:

```json
{
	"status": "success",
	"data": {
		"key": "python",
		"name": "Python 3",
		"description": "Python is an interpreted, high-level, general-purpose programming language.",
		"logo": "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
		"tags": ["programming language", "scripting"]
	}
}
```

failure response:

```json
{
	"status": "error",
	"message": "not be able to update skill name"
}
```

6. `PATCH /api/v1/skills/:key/actions/description`

description: Update the description of a skill

example: PATCH /api/v1/skills/python/actions/description

payload:

```json
{
	"description": "Python 3 is the latest version of Python programming language."
}
```

response:

```json
{
	"status": "success",
	"data": {
		"key": "python",
		"name": "Python",
		"description": "Python 3 is the latest version of Python programming language.",
		"logo": "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
		"tags": ["programming language", "scripting"]
	}
}
```

failure response:

```json
{
	"status": "error",
	"message": "not be able to update skill description"
}
```

7. `PATCH /api/v1/skills/:key/actions/logo`

description: Update the logo of a skill

example: PATCH /api/v1/skills/python/actions/logo

payload:

```json
{
	"logo": "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg"
}
```

response:

```json
{
	"status": "success",
	"data": {
		"key": "python",
		"name": "Python",
		"description": "Python is an interpreted, high-level, general-purpose programming language.",
		"logo": "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
		"tags": ["programming language", "scripting"]
	}
}
```

failure response:

```json
{
	"status": "error",
	"message": "not be able to update skill logo"
}
```

8. `PATCH /api/v1/skills/:key/actions/tags`

description: Update the tags of a skill

example: PATCH /api/v1/skills/python/actions/tags

payload:

```json
{
	"tags": ["programming language", "data"]
}
```

response:

```json
{
	"status": "success",
	"data": {
		"key": "python",
		"name": "Python",
		"description": "Python is an interpreted, high-level, general-purpose programming language.",
		"logo": "https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg",
		"tags": ["programming language", "data"]
	}
}
```

failure response:

```json
{
	"status": "error",
	"message": "not be able to update skill tags"
}
```

9. `DELETE /api/v1/skills/:key`

description: Delete a skill

example: DELETE /api/v1/skills/python

response:

```json
{
	"status": "success",
	"message": "Skill deleted"
}
```

failure response:

```json
{
	"status": "error",
	"message": "not be able to delete skill"
}
```