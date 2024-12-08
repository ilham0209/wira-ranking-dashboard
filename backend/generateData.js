const faker = require('faker');
const { Pool } = require('pg');
const pool = new Pool({
  user: 'wira_user',           
  host: 'localhost',
  database: 'wira_db',        
  password: 'ilham2002',   
  port: 5432,
});

async function generateData() {
  const client = await pool.connect();

  try {

    for (let i = 0; i < 100000; i++) {
      let email;
      let emailExists = true;
      
      while (emailExists) {
        email = faker.internet.email();
        
        const res = await client.query('SELECT COUNT(*) FROM Account WHERE email = $1', [email]);
        if (parseInt(res.rows[0].count) === 0) {
          emailExists = false; 
        }
      }

      const username = faker.internet.userName();

      await client.query('INSERT INTO Account(username, email) VALUES($1, $2)', [username, email]);

      if (i % 1000 === 0) {
        console.log(`Inserted ${i} accounts...`);
      }
    }
    const accounts = await client.query('SELECT acc_id FROM Account');
    const charIds = [];

    for (const account of accounts.rows) {
      const classId = Math.floor(Math.random() * 8) + 1; 
      const res = await client.query('INSERT INTO Character(acc_id, class_id) VALUES($1, $2) RETURNING char_id', [account.acc_id, classId]);

      charIds.push(res.rows[0].char_id);
    }

    console.log('Character data inserted. Now inserting scores...');

    for (const char_id of charIds) {
      const score = Math.floor(Math.random() * 1000);
      await client.query('INSERT INTO Scores(char_id, reward_score) VALUES($1, $2)', [char_id, score]);
    }

    console.log('Data generation completed!');
  } catch (err) {
    console.error('Error generating data:', err);
  } finally {
    client.release();
  }
}

generateData();
