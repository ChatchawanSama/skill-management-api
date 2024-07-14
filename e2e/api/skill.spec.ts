import { test, expect } from '@playwright/test';
import { Client } from 'pg';

const BASE_URL = 'http://localhost:8080/api/v1/skills';

let client: Client;

test.describe('Skill API Tests', () => {

  test.beforeAll(async () => {
    // Initialize PostgreSQL client
    client = new Client({
      connectionString: 'postgresql://skillapi:skillapi@localhost:5432/skill?sslmode=disable'
    });

    await client.connect();

    // Insert initial test data
    const queryText = `
      INSERT INTO skill (key, name, description, logo, tags)
      VALUES
      ('skill-1', 'Test Skill 1', 'Description for Test Skill 1', 'http://example.com/logo1.png', ARRAY['tag1', 'tag2']),
      ('skill-2', 'Test Skill 2', 'Description for Test Skill 2', 'http://example.com/logo2.png', ARRAY['tag3', 'tag4']);
    `;
    await client.query(queryText);
  });

  test.afterAll(async () => {
    // Delete test data
    const queryText = `
      DELETE FROM skill WHERE key IN ('skill-1', 'skill-2');
    `;
    await client.query(queryText);

    // Close PostgreSQL client
    await client.end();
  });

  test('GET /api/v1/skills', async ({ request }) => {
    const response = await request.get(BASE_URL);
    expect(response.status()).toBe(200);
    const responseBody = await response.json();
    expect(responseBody.status).toBe('success');
    expect(responseBody.data.length).toBeGreaterThan(0);
  });

  test('POST /api/v1/skills', async ({ request }) => {
    const newSkill = {
      key: 'skill-3',
      name: 'Skill 3',
      description: 'Description for Skill 3',
      logo: 'http://example.com/logo3.png',
      tags: ['tag5', 'tag6']
    };
    const response = await request.post(BASE_URL, { data: newSkill });
    expect(response.status()).toBe(200);
    const responseBody = await response.json();
    expect(responseBody.status).toBe('success');
    expect(responseBody.data.key).toBe(newSkill.key);

    // Cleanup
    await client.query('DELETE FROM skill WHERE key = $1', [newSkill.key]);
  });

  test('GET /api/v1/skills/:key', async ({ request }) => {
    const response = await request.get(`${BASE_URL}/skill-1`);
    expect(response.status()).toBe(200);
    const responseBody = await response.json();
    expect(responseBody.status).toBe('success');
    expect(responseBody.data.key).toBe('skill-1');
  });

  test('PUT /api/v1/skills/:key', async ({ request }) => {
    const updatedSkill = {
      name: 'Updated Skill 1',
      description: 'Updated Description for Skill 1',
      logo: 'http://example.com/logo-updated.png',
      tags: ['tag3', 'tag4']
    };
    const response = await request.put(`${BASE_URL}/skill-1`, { data: updatedSkill });
    expect(response.status()).toBe(200);
    const responseBody = await response.json();
    expect(responseBody.status).toBe('success');
    expect(responseBody.data.name).toBe(updatedSkill.name);
  });

  test('PATCH /api/v1/skills/:key/actions/name', async ({ request }) => {
    const patchData = { name: 'Patched Skill Name' };
    const response = await request.patch(`${BASE_URL}/skill-1/actions/name`, { data: patchData });
    expect(response.status()).toBe(200);
    const responseBody = await response.json();
    expect(responseBody.status).toBe('success');
    expect(responseBody.data.name).toBe(patchData.name);
  });

  test('PATCH /api/v1/skills/:key/actions/description', async ({ request }) => {
    const patchData = { description : 'Patched Skill Description' };
    const response = await request.patch(`${BASE_URL}/skill-1/actions/description`, { data: patchData });
    expect(response.status()).toBe(200);
    const responseBody = await response.json();
    expect(responseBody.status).toBe('success');
    expect(responseBody.data.description).toBe(patchData.description);
  });

  test('PATCH /api/v1/skills/:key/actions/logo', async ({ request }) => {
    const patchData = { logo: 'Patched Skill Logo' };
    const response = await request.patch(`${BASE_URL}/skill-1/actions/logo`, { data: patchData });
    expect(response.status()).toBe(200);
    const responseBody = await response.json();
    expect(responseBody.status).toBe('success');
    expect(responseBody.data.logo).toBe(patchData.logo);
  });

  test('PATCH /api/v1/skills/:key/actions/tags', async ({ request }) => {
    const patchData = { tags: ['Patched Skill Tags', 'Patched Skill Tags'] };
    const response = await request.patch(`${BASE_URL}/skill-1/actions/tags`, { data: patchData });
    expect(response.status()).toBe(200);
    const responseBody = await response.json();
    expect(responseBody.status).toBe('success');
    expect(responseBody.data.tags).toStrictEqual(patchData.tags);
  });

  test('DELETE /api/v1/skills/:key', async ({ request }) => {
    const response = await request.delete(`${BASE_URL}/skill-1`);
    expect(response.status()).toBe(200);
    const responseBody = await response.json();
    expect(responseBody.status).toBe('success');

    // Re-insert the deleted skill for consistency in other tests
    const queryText = `
      INSERT INTO skill (key, name, description, logo, tags)
      VALUES ('skill-1', 'Test Skill 1', 'Description for Test Skill 1', 'http://example.com/logo1.png', ARRAY['tag1', 'tag2']);
    `;
    await client.query(queryText);
  });
});
