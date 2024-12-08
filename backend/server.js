const express = require('express');
const cors = require('cors');
const { Pool } = require('pg');
const app = express();
const port = 3000;

app.use(cors());
app.use(express.json());

const pool = new Pool({
  user: 'wira_user',
  host: 'localhost',
  database: 'wira_db',
  password: 'ilham2002',
  port: 5432,
});

app.get('/api/rankings', async (req, res) => {
  const { page = 1, limit = 10 } = req.query;

  if (isNaN(page) || page <= 0) {
    return res.status(400).send('Invalid page number');
  }
  if (isNaN(limit) || limit <= 0) {
    return res.status(400).send('Invalid limit');
  }

  const offset = (page - 1) * limit;

  const query = `
    SELECT a.username, s.reward_score 
    FROM Account a
    JOIN Scores s ON a.acc_id = s.char_id
    ORDER BY s.reward_score DESC 
    LIMIT $1 OFFSET $2
  `;
  
  const queryParams = [limit, offset];

  try {
    const { rows } = await pool.query(query, queryParams);
    console.log("Data yang dikembalikan:", rows);  
    res.json(rows);  
  } catch (err) {
    console.error("Error executing query:", err);
    res.status(500).send('Error retrieving rankings');
  }
});

app.listen(port, () => {
  console.log(`Server started at http://localhost:${port}`);
});
